package modules

import (
	"time"
)

// gen:qs
type CircleUnitProperty struct {
	ID           uint       `description:""`
	CreatedAt    time.Time  `description:"등록일"`
	UpdatedAt    time.Time  `description:"수정일"`
	Name         string     `description:"이름"`
	Description  string     `description:"설명" sql:"type:text"`
	CreatorID    uint       `description:"작성자"`
	CircleUnit   CircleUnit `description:""`
	CircleUnitID uint       `description:""`
	Type         string     `description:""`
	Nullable     bool       `description:""`
	IsEnable     bool       `description:"사용여부"`
	IsManual     bool       `description:""`
	IsSystem     bool       `description:""`
	UseRequest   bool       `description:""`
	UseResponses bool       `description:""`
}

func (c *CircleUnitProperty) GetCreatorID() uint {
	return c.CreatorID
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
