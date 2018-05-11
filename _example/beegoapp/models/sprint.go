package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// gen:qs
type Sprint struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	CreatorID   uint      `description:"작성자"`
	Current     bool      ``
}

func init() {
	registModel(&Sprint{})
}

func (m *Sprint) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Sprint) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddSprint(sprint *Sprint) (id uint, err error) {
	err = sprint.Create(gGormDB)
	id = sprint.ID
	return
}

func GetSprintByID(id uint) (sprint *Sprint, err error) {
	sprint = &Sprint{
		ID: id,
	}
	err = NewSprintQuerySet(gGormDB).One(sprint)
	return
}

func GetAllSprint(queryPage *modules.QueryPage) (sprints []Sprint, err error) {
	err = NewSprintQuerySet(gGormDB).All(&sprints)
	return
}

func UpdateSprintByID(sprint *Sprint) (err error) {
	err = sprint.Update(gGormDB,
		SprintDBSchema.Description,
	)
	return
}

func DeleteSprint(id uint) (err error) {
	sprint := &Sprint{
		ID: id,
	}
	err = sprint.Delete(gGormDB)
	return
}
