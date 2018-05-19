package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// gen:qs
type Test struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	CreatorID   uint      `description:"작성자"`
}

func init() {
	registModel(&Test{})
}

func (m *Test) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Test) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddTest(test *Test) (id uint, err error) {
	err = test.Create(gGormDB)
	id = test.ID
	return
}

func GetTestByID(id uint) (test *Test, err error) {
	test = &Test{
		ID: id,
	}
	err = NewTestQuerySet(gGormDB).
		One(test)
	return
}

func GetAllTest(queryPage *modules.QueryPage) (tests []Test, err error) {
	err = NewTestQuerySet(gGormDB).
		All(&tests)
	return
}

func UpdateTestByID(test *Test) (err error) {
	err = test.Update(gGormDB,
		TestDBSchema.Name,
		TestDBSchema.Description,
	)
	return
}

func DeleteTest(id uint) (err error) {
	test := &Test{
		ID: id,
	}
	err = test.Delete(gGormDB)
	return
}
