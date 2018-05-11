package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// gen:qs
type Ics struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	CreatorID   uint      `description:"작성자"`
	ICSURL      string
}

func init() {
	registModel(&Ics{})
}

func (m *Ics) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Ics) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddIcs(ics *Ics) (id uint, err error) {
	err = ics.Create(gGormDB)
	id = ics.ID
	return
}

func GetIcsByID(id uint) (ics *Ics, err error) {
	ics = &Ics{
		ID: id,
	}
	err = NewIcsQuerySet(gGormDB).
		One(ics)
	return
}

func GetAllIcs(queryPage *modules.QueryPage) (icss []Ics, err error) {
	err = NewIcsQuerySet(gGormDB).
		All(&icss)
	return
}

func UpdateIcsByID(ics *Ics) (err error) {
	err = ics.Update(gGormDB,
		IcsDBSchema.Name,
		IcsDBSchema.Description,
	)
	return
}

func DeleteIcs(id uint) (err error) {
	ics := &Ics{
		ID: id,
	}
	err = ics.Delete(gGormDB)
	return
}
