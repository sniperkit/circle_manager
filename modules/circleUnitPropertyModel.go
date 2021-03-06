package modules

import (
	"strings"
	"time"
)

// gen:qs
type CircleUnitProperty struct {
	ID                             uint       `description:""`
	CreatedAt                      time.Time  `description:"등록일"`
	UpdatedAt                      time.Time  `description:"수정일"`
	Name                           string     `description:"이름"`
	Description                    string     `description:"설명" sql:"type:text"`
	CreatorID                      uint       `description:"작성자"`
	CircleUnit                     CircleUnit `description:""`
	CircleUnitID                   uint       `description:""`
	Type                           string     `description:""`
	Nullable                       bool       `description:""`
	IsEnable                       bool       `description:"사용여부"`
	IsManual                       bool       `description:""`
	IsSystem                       bool       `description:""`
	IsCreateble                    bool       `description:""`
	CreatebleUserIDs               string     `description:""`
	CreatebleUserTypeIDs           string     `description:""`
	CreatebleUserStatusIDs         string     `description:""`
	CreatebleUserExcludeIDs        string     `description:""`
	CreatebleUserExcludeTypeIDs    string     `description:""`
	CreatebleUserExcludeStatusIDs  string     `description:""`
	IsGetAllable                   bool       `description:""`
	GetAllableUserIDs              string     `description:""`
	GetAllableUserTypeIDs          string     `description:""`
	GetAllableUserStatusIDs        string     `description:""`
	GetAllableUserExcludeIDs       string     `description:""`
	GetAllableUserExcludeTypeIDs   string     `description:""`
	GetAllableUserExcludeStatusIDs string     `description:""`
	IsGetOneable                   bool       `description:""`
	GetOneableUserIDs              string     `description:""`
	GetOneableUserTypeIDs          string     `description:""`
	GetOneableUserStatusIDs        string     `description:""`
	GetOneableUserExcludeIDs       string     `description:""`
	GetOneableUserExcludeTypeIDs   string     `description:""`
	GetOneableUserExcludeStatusIDs string     `description:""`
	IsDeleteble                    bool       `description:""`
	DeletableUserIDs               string     `description:""`
	DeletableUserTypeIDs           string     `description:""`
	DeletableUserStatusIDs         string     `description:""`
	DeletableUserExcludeIDs        string     `description:""`
	DeletableUserExcludeTypeIDs    string     `description:""`
	DeletableUserExcludeStatusIDs  string     `description:""`
	IsUpdateble                    bool       `description:""`
	UpdatableUserIDs               string     `description:""`
	UpdatableUserTypeIDs           string     `description:""`
	UpdatableUserStatusIDs         string     `description:""`
	UpdatableUserExcludeIDs        string     `description:""`
	UpdatableUserExcludeTypeIDs    string     `description:""`
	UpdatableUserExcludeStatusIDs  string     `description:""`
}

func (c *CircleUnitProperty) GetTypeInModel() string {
	return strings.Replace(c.Type, "models.", "", -1)
}

func AddCircleUnitProperty(circleUnitProperty *CircleUnitProperty) (id uint, err error) {
	err = circleUnitProperty.Create(gGormDB)
	id = circleUnitProperty.ID
	return
}

func GetCircleUnitPropertyByID(id uint) (circleUnitProperty *CircleUnitProperty, err error) {
	circleUnitProperty = &CircleUnitProperty{
		ID: id,
	}
	err = NewCircleUnitPropertyQuerySet(gGormDB).One(circleUnitProperty)
	return
}

func GetAllCircleUnitProperty(queryPage *QueryPage) (circleUnitPropertys []CircleUnitProperty, err error) {
	err = NewCircleUnitPropertyQuerySet(gGormDB).All(&circleUnitPropertys)
	return
}

func UpdateCircleUnitPropertyByID(circleUnitProperty *CircleUnitProperty) (err error) {
	err = circleUnitProperty.Update(gGormDB,
		CircleUnitPropertyDBSchema.Description,
	)
	return
}

func DeleteCircleUnitProperty(id uint) (err error) {
	circleUnitProperty := &CircleUnitProperty{
		ID: id,
	}
	err = circleUnitProperty.Delete(gGormDB)
	return
}
