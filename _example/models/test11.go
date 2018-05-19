package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// gen:qs
type Test11 struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	CreatorID   uint      `description:"작성자"`
}

func init() {
	registModel(&Test11{})
}

func (m *Test11) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Test11) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddTest11(test11 *Test11) (id uint, err error) {
	err = test11.Create(gGormDB)
	id = test11.ID
	return
}

func GetTest11ByID(id uint) (test11 *Test11, err error) {
	test11 = &Test11{
		ID: id,
	}
	err = NewTest11QuerySet(gGormDB).
		One(test11)
	return
}

func GetAllTest11(queryPage *modules.QueryPage) (test11s []Test11, err error) {
	err = NewTest11QuerySet(gGormDB).
		All(&test11s)
	return
}

func UpdateTest11ByID(test11 *Test11) (err error) {
	err = test11.Update(gGormDB,
		Test11DBSchema.Name,
		Test11DBSchema.Description,
	)
	return
}

func DeleteTest11(id uint) (err error) {
	test11 := &Test11{
		ID: id,
	}
	err = test11.Delete(gGormDB)
	return
}
