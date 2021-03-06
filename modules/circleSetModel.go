package modules

import (
	"time"

	"github.com/jinzhu/gorm"
)

// gen:qs
type CircleSet struct {
	ID                    uint          `description:""`
	CreatedAt             time.Time     `description:"등록일"`
	UpdatedAt             time.Time     `description:"수정일"`
	Name                  string        `description:"이름"`
	Description           string        `description:"설명" sql:"type:text"`
	CreatorID             uint          `description:"작성자"`
	UpdaterID             uint          `description:"최종수정자"`
	Import                string        `description:""`
	Units                 []*CircleUnit `description:""`
	IsEnable              bool          `description:"사용여부"`
	AppVersion            string        `description:""`
	AppTitle              string        `description:""`
	AppDescription        string        `description:""`
	AppContact            string        `description:""`
	AppTermsOfServiceUrl  string        `description:""`
	AppLicense            string        `description:""`
	AppSecurityDefinition string        `description:""`
	RunAppEnvs            string        `description:""`
}

func (c CircleSet) GetUnit(unitName string) *CircleUnit {
	for _, unit := range c.Units {
		if !unit.IsSystem {
			if unit.Name == unitName {
				copy := unit
				return copy
			}
		}
	}
	return nil
}

func (c CircleSet) GetAutoGenUnits() []*CircleUnit {
	units := []*CircleUnit{}
	for _, unit := range c.Units {
		if !unit.IsManual && !unit.IsSystem {
			units = append(units, unit)
		}
	}
	return units
}

func (c *CircleSet) SetAppMeta(metaKey, value string) {
	switch metaKey {
	case "Title":
		c.AppTitle = value
	case "APIVersion":
		c.AppVersion = value
	case "Description":
		c.AppDescription = value
	case "Contact":
		c.AppContact = value
	case "TermsOfServiceUrl":
		c.AppTermsOfServiceUrl = value
	case "License":
		c.AppLicense = value
	case "SecurityDefinition":
		c.AppSecurityDefinition = value
	}
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

	preloadDB := gGormDB.Preload("Units").Preload("Units.Properties")

	err = NewCircleSetQuerySet(preloadDB).
		One(circleSet)
	return
}

func GetCircleSetByIDOnlyManual(id uint) (circleSet *CircleSet, err error) {
	circleSet = &CircleSet{
		ID: id,
	}

	preloadDB := gGormDB.Preload("Units", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_manual = 1")
	})

	preloadDB = preloadDB.Preload("Units.Properties")

	err = NewCircleSetQuerySet(preloadDB).
		One(circleSet)
	return
}

func GetCircleSetByIDForGen(id uint) (circleSet *CircleSet, err error) {
	circleSet = &CircleSet{
		ID: id,
	}

	preloadDB := gGormDB.Preload("Units", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_enable = ? && is_manual = 0", true)
	})

	preloadDB = preloadDB.Preload("Units.Properties")

	err = NewCircleSetQuerySet(preloadDB).
		One(circleSet)
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
