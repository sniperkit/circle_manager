package modules

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	gGormDB          *gorm.DB
	userGormDB       map[uint]*gorm.DB
	validate         *validator.Validate
	querygt          = "gt"
	querygte         = "gte"
	queryin          = "in"
	queryexact       = "exact"
	queryiexact      = "iexact"
	querycontains    = "contains"
	queryicontains   = "icontains"
	querylt          = "lt"
	querylte         = "lte"
	querystartswith  = "startswith"
	queryistartswith = "istartswith"
	queryendswith    = "endswith"
	queryiendswith   = "iendswith"
	queryisnull      = "isnull"
	queryisnotnull   = "isnotnull"
	queryFilters     = map[string]string{
		querygt:          "> ?",
		querygte:         ">= ?",
		queryin:          "in ?",
		queryexact:       "= ?",
		queryiexact:      "LIKE ?",
		querycontains:    "LIKE BINARY ?",
		queryicontains:   "LIKE ?",
		querylt:          "< ?",
		querylte:         "<= ?",
		querystartswith:  "LIKE BINARY ?",
		queryistartswith: "LIKE ?",
		queryendswith:    "LIKE BINARY ?",
		queryiendswith:   "LIKE ?",
		queryisnull:      "IS NULL ?",
		queryisnotnull:   "IS NOT NULL ?",
	}
)

type ParamGetValueByKeyOfTableName struct {
	TableName string
	ID        uint
	Key       string
	Value     interface{}
}

func CreateItem(item interface{}) error {
	err := gGormDB.Create(item).Error
	if err != nil {
		return SaveItem(item)
	}
	return nil
}

func SaveItem(item interface{}) error {
	if err := gGormDB.Save(item).Error; err != nil {
		return err
	}
	return nil
}

func GetItemByID(id uint, item interface{}) error {
	return gGormDB.Where("id = ?", id).First(item).Error
}

func GetItemWithFilter(filterName string, filterValue interface{}, item interface{}) error {
	return gGormDB.Where(fmt.Sprintf("%s = ?", filterName), filterValue).First(item).Error
}

func GetItems(items interface{}, queryPage *QueryPage) error {
	return gGormDB.Find(items).Error
}

func GetValueByKeyOfTableName(param ParamGetValueByKeyOfTableName) (interface{}, error) {
	var value interface{}
	if err := gGormDB.Table(param.TableName).Select(param.Key).Where("id = ?", param.ID).Row().Scan(&value); err != nil {
		return nil, err
	}
	if value != nil {
		return string(value.([]byte)), nil
	}
	return nil, errors.New("GetValueByKeyOfTableName 알수없는 에러")
}

func GetItemsOnlyUserData(items interface{}, queryPage *QueryPage, userID uint) error {
	return gGormDB.Where("creator_id == ?", userID).Find(items).Error
}

func UpdateItem(id uint, item interface{}) error {
	return gGormDB.Where("id = ?", id).Save(item).Error
}

func DeleteItem(id uint, item interface{}) error {
	return gGormDB.Delete(item, "id = ?", id).Error
}

func (q QueryPage) ParsedQueryPage(db *gorm.DB) *gorm.DB {
	db = db.Offset(q.Offset)
	db = db.Limit(q.Limit)

	//https://beego.me/docs/mvc/model/query.md 사용
	for key, v := range q.Query {
		keySplit := strings.Split(key, "__")
		queryKey := ""
		queryFilter := queryexact
		if len(keySplit) == 1 {
			queryKey = keySplit[0]
		} else if len(keySplit) == 2 {
			if _, ok := queryFilters[keySplit[1]]; ok {
				queryFilter = keySplit[1]
				queryKey = keySplit[0]
			} else {
				queryKey = fmt.Sprintf("%s.%s", keySplit[0], keySplit[1])
			}
		} else if len(keySplit) == 3 {
			queryKey = fmt.Sprintf("%s.%s", keySplit[0], keySplit[1])
			queryFilter = keySplit[2]
		}

		nv := v
		if queryFilter == querycontains || queryFilter == queryicontains {
			nv = fmt.Sprintf("%%%s%%", v)
		} else if queryFilter == querystartswith || queryFilter == queryistartswith {
			nv = fmt.Sprintf("%s%%", v)
		} else if queryFilter == queryendswith || queryFilter == queryiendswith {
			nv = fmt.Sprintf("%%%s", v)
		}

		queryFilterParsed := queryFilters[queryexact]
		if _, ok := queryFilters[queryFilter]; ok {
			queryFilterParsed = queryFilters[queryFilter]
		}

		db = db.Where(fmt.Sprintf("%s %s", queryKey, queryFilterParsed), nv)
	}

	if len(q.Sortby) != 0 {
		if len(q.Sortby) == len(q.Order) {
			for i, v := range q.Sortby {
				asc := true

				if len(q.Order) <= i+1 {
					if q.Order[i] == "desc" {
						asc = false
					}
				}

				db = db.Order(v, asc)
			}
		}
	}

	return db
}

func GetRows(tableName string, where string) ([]map[string]interface{}, error) {
	workingDB := gGormDB.Table(tableName)
	if where != "" {
		workingDB = workingDB.Where(where)
	}
	rows, err := workingDB.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mapColumnTypeIndex := map[int]*sql.ColumnType{}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	for index, columnType := range columnTypes {
		mapColumnTypeIndex[index] = columnType
	}

	retRows := []map[string]interface{}{}
	for rows.Next() {
		//https://kylewbanks.com/blog/query-result-to-map-in-golang
		columns := make([]interface{}, len(columnTypes))
		columnPointers := make([]interface{}, len(columnTypes))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		retCols := map[string]interface{}{}
		for index, column := range columnPointers {
			colType := mapColumnTypeIndex[index]
			colName := colType.Name()
			val := column.(*interface{})
			retCols[colName] = *val
		}
		retRows = append(retRows, retCols)
	}
	return retRows, nil
}

func ResetTable(tableName string) error {
	if err := gGormDB.Exec(fmt.Sprintf("DELETE FROM %s", tableName)).Error; err != nil {
		return err
	}
	return gGormDB.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", tableName)).Error
}

func DeleteItemByColName(tableName string, colName string, value string) error {
	return gGormDB.Table(tableName).Where(fmt.Sprintf("%s = ?", colName), value).Delete(nil).Error
}
