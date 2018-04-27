package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fatih/structs"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	CurCircleSet *CircleSet
)

//BaseController ...
type BaseController struct {
	beego.Controller
	CurrentUserMeta *UserMeta
}

// QueryPage ...
type QueryPage struct {
	Fields     []string
	Sortby     []string
	Order      []string
	Query      map[string]string
	Limit      int
	Offset     int
	Tags       []uint
	Properties map[string][]string
}

// Success ...
func (c *BaseController) Success(statusCode int, data interface{}) {
	ext := c.Ctx.Input.Param(":ext")
	c.Ctx.Output.SetStatus(statusCode)
	if data != nil {
		if ext == "xml" {
			c.Data["xml"] = data
			c.ServeXML()
		} else {
			c.Data["json"] = data
			c.ServeJSON()
		}
	}
}

func getUserIDByUserMeta(userMeta *UserMeta) *uint {
	if userMeta == nil {
		return nil
	}
	return &userMeta.UserID
}

// SuccessCreate ...
func (c *BaseController) SuccessCreate(responseBody ResponseBody) {
	c.Success(http.StatusCreated, responseBody)
}

// SuccessUpdate ...
func (c *BaseController) SuccessUpdate(responseBody ResponseBody) {
	c.Success(http.StatusOK, responseBody)
}

// SuccessDelete ...
func (c *BaseController) SuccessDelete() {
	c.Success(http.StatusNoContent, nil)
}

func EventThenCreate(modelItem ModelItem, currentUserID *uint) {
	evnet(structs.Name(modelItem), "add", currentUserID, nil, modelItem)
}

func EventThenUpdate(modelItem ModelItem, oldModelItem ModelItem, currentUserID *uint) {
	mapUpdateProperties := makeMapUpdateProperties(modelItem, oldModelItem)
	evnet(structs.Name(modelItem), "update", currentUserID, mapUpdateProperties, modelItem)
}

func makeMapUpdateProperties(modelItem ModelItem, oldModelItem ModelItem) map[string]UpdateProperty {
	mapUpdateProperties := map[string]UpdateProperty{}

	mapModelItem := structs.Map(modelItem)
	mapOldModelItem := structs.Map(oldModelItem)

	for key, value := range mapModelItem {
		if structs.IsStruct(value) {
			continue
		}

		oldValue := ""
		if tempOldValue, ok := mapOldModelItem[key]; ok {
			oldValue = convInterface(tempOldValue)
		}

		mapUpdateProperties[key] = UpdateProperty{
			Key:      key,
			NewValue: convInterface(value),
			OldValue: oldValue,
		}
	}
	return mapUpdateProperties
}

func EventThenDelete(modelItem ModelItem, currentUserID *uint) {
	evnet(structs.Name(modelItem), "delete", currentUserID, nil, modelItem)
}

func evnet(structName string, action string, eventUserID *uint, mapUpdateProperties map[string]UpdateProperty, datas ...interface{}) {
	if err := AddActionNotification(
		fmt.Sprintf("%s,%s", structName, action),
		eventUserID,
		mapUpdateProperties,
		datas...,
	); err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}
}

// ErrorAbort ...
func (c *BaseController) ErrorAbort(code int, err error, withMsg ...interface{}) {
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
		c.CustomAbort(code, err.Error())
	}
	c.CustomAbort(code, "")
}

func (c *BaseController) GetQueryPage() *QueryPage {
	queryPage := &QueryPage{
		Query: make(map[string]string),
		Limit: 10,
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		queryPage.Fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetUint64("limit"); err == nil {
		queryPage.Limit = int(v)
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		queryPage.Offset = int(v)
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		queryPage.Sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		queryPage.Order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.ErrorAbort(400, ErrInvalidQuery)
			}
			k, v := kv[0], kv[1]
			queryPage.Query[k] = v
		}
	}
	return queryPage
}

func (c *BaseController) setRequestDataInterface(reqBody interface{}) error {
	if len(c.Ctx.Input.RequestBody) == 0 {
		logrus.Error("Body에 정보가 없습니다.")
		return ErrInvalidRequestBody
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody); err != nil {
		return err
	}
	return nil
}

func (c *BaseController) setRequestData(reqBody RequestBody) error {
	if len(c.Ctx.Input.RequestBody) == 0 {
		logrus.Error("Body에 정보가 없습니다.")
		return ErrInvalidRequestBody
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody); err != nil {
		return err
	}
	return nil
}

func (c *BaseController) SetRequestDataAndValid(reqBody RequestBody) {
	if len(c.Ctx.Input.RequestBody) == 0 {
		logrus.Error("Body에 정보가 없습니다.")
		c.ErrorAbort(400, ErrInvalidRequestBody)
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody); err != nil {
		c.ErrorAbort(400, err)
	}

	// if err := c.setRequestData(reqBody); err != nil {
	// 	c.ErrorAbort(400, err)
	// }
	if err := reqBody.Valid(); err != nil {
		c.ErrorAbort(400, err)
	}
}

func (c *BaseController) GetParamForUintTypeRequired(key string, retID *uint) error {
	idString := c.Ctx.Input.Param(key)
	if idString == "" {
		return ErrInvalidRequestParam
	}
	id, err := strconv.ParseUint(idString, 10, 64)
	if err == nil {
		if id == 0 {
			return ErrInvalidRequestParam
		}
		uintID := uint(id)
		*retID = uintID
		return nil
	}
	return ErrInvalidRequestParam
}

func (c *BaseController) GetParamID() uint {
	id := uint(0)
	if err := c.GetParamForUintTypeRequired(":id", &id); err != nil {
		c.ErrorAbort(400, err)
	}
	return id
}

func (c *BaseController) Check404And500(err error) {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.ErrorAbort(404, err)
		} else {
			c.ErrorAbort(500, err)
		}
	}
}
