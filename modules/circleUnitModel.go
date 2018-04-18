package modules

import (
	"fmt"
	"time"

	stringcase "github.com/reiver/go-stringcase"
)

// gen:qs
type CircleUnit struct {
	ID          uint                 `description:""`
	CreatedAt   time.Time            `description:"등록일"`
	UpdatedAt   time.Time            `description:"수정일"`
	Name        string               `description:"이름"`
	Description string               `description:"설명" sql:"type:text"`
	CircleSet   CircleSet            ``
	CircleSetID uint                 ``
	Properties  []CircleUnitProperty `description:""`
	Import      string               `description:""`
	URL         string               `description:""`
	MenuName    string               `description:""`
	MenuGroup   string               `description:""`
	IsEnable    bool                 `description:"사용여부"`
	IsManual    bool                 `description:""`
	IsSystem    bool                 `description:""`
}

func (c CircleUnit) GetVariableName() string {
	return stringcase.ToCamelCase(c.Name)
}

func (c CircleUnit) GetControllerName() string {
	return fmt.Sprintf("%sController", c.Name)
}

func AddCircleUnit(circleUnit *CircleUnit) (id uint, err error) {
	err = circleUnit.Create(gGormDB)
	id = circleUnit.ID
	return
}

func GetCircleUnitByID(id uint) (circleUnit *CircleUnit, err error) {
	circleUnit = &CircleUnit{
		ID: id,
	}
	err = NewCircleUnitQuerySet(gGormDB).One(circleUnit)
	return
}

func GetAllCircleUnit(queryPage *QueryPage) (circleUnits []CircleUnit, err error) {
	err = NewCircleUnitQuerySet(gGormDB).All(&circleUnits)
	return
}

func UpdateCircleUnitByID(circleUnit *CircleUnit) (err error) {
	err = circleUnit.Update(gGormDB,
		CircleUnitDBSchema.Description,
	)
	return
}

func DeleteCircleUnit(id uint) (err error) {
	circleUnit := &CircleUnit{
		ID: id,
	}
	err = circleUnit.Delete(gGormDB)
	return
}

func SaveCircleUnit(nameAndValue map[string]interface{}) error {
	//TODO: ....
	return nil
}
