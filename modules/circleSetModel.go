package modules

import (
	"time"
)

// gen:qs
type CircleSet struct {
	ID                    uint         `description:""`
	CreatedAt             time.Time    `description:"등록일"`
	UpdatedAt             time.Time    `description:"수정일"`
	Name                  string       `description:"이름"`
	Description           string       `description:"설명" sql:"type:text"`
	Import                string       `description:""`
	Units                 []CircleUnit `description:""`
	IsEnable              bool         `description:"사용여부"`
	AppVersion            string       `description:""`
	AppTitle              string       `description:""`
	AppDescription        string       `description:""`
	AppContact            string       `description:""`
	AppTermsOfServiceUrl  string       `description:""`
	AppLicense            string       `description:""`
	AppSecurityDefinition string       `description:""`
	RunAppEnvs            string       `description:""`
}

func AddCircleSet(circleSet *CircleSet) (id uint, err error) {
	err = circleSet.Create(gGormDB)
	id = circleSet.ID
	return
}

func GetCircleSetByID(id uint) (circleSet *CircleSet, err error) {
	circleSet = &CircleSet{
		ID: id,
	}
	err = NewCircleSetQuerySet(gGormDB).One(circleSet)
	return
}

func GetAllCircleSet(queryPage *QueryPage) (circleSets []CircleSet, err error) {
	err = NewCircleSetQuerySet(gGormDB).All(&circleSets)
	return
}

func UpdateCircleSetByID(circleSet *CircleSet) (err error) {
	err = circleSet.Update(gGormDB,
		CircleSetDBSchema.Description,
	)
	return
}

func DeleteCircleSet(id uint) (err error) {
	circleSet := &CircleSet{
		ID: id,
	}
	err = circleSet.Delete(gGormDB)
	return
}
