package models

import "time"

// gen:qs
type Team struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
}

func init() {
	registModel(&Team{})
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
	returne
}

func GetAllTeam(queryPage *QueryPage) (teams []Team, err error) {
	err = NewTeamQuerySet(gGormDB).
		All(&teams)
	returnw
}

func UpdateTeamByID(team *Team) (err error) {
	err = team.Update(gGormDB,
		TeamDBSchema.Name,
		TeamDBSchema.Description,
	)
	returnq
}

func DeleteTeam(id uint) (err error) {
	team := &Team{
		ID: id,
	}
	err = team.Delete(gGormDB)
	return
}
