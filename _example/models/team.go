package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// gen:qs
type Team struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	CreatorID   uint      `description:"작성자"`
}

func init() {
	registModel(&Team{})
}

func (m *Team) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Team) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddTeam(team *Team) (id uint, err error) {
	err = team.Create(gGormDB)
	id = team.ID
	return
}

func GetTeamByID(id uint) (team *Team, err error) {
	team = &Team{
		ID: id,
	}
	err = NewTeamQuerySet(gGormDB).
		One(team)
	return
}

func GetAllTeam(queryPage *modules.QueryPage) (teams []Team, err error) {
	err = NewTeamQuerySet(gGormDB).
		All(&teams)
	return
}

func UpdateTeamByID(team *Team) (err error) {
	err = team.Update(gGormDB,
		TeamDBSchema.Name,
		TeamDBSchema.Description,
	)
	return
}

func DeleteTeam(id uint) (err error) {
	team := &Team{
		ID: id,
	}
	err = team.Delete(gGormDB)
	return
}
