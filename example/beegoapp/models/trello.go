package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// gen:qs
type Trello struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	CreatorID   uint      `description:"작성자"`
	UserName    string
	Token       string
	Key         string
}

func init() {
	registModel(&Trello{})
}

func (m *Trello) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Trello) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddTrello(trello *Trello) (id uint, err error) {
	err = trello.Create(gGormDB)
	id = trello.ID
	return
}

func GetTrelloByID(id uint) (trello *Trello, err error) {
	trello = &Trello{
		ID: id,
	}
	err = NewTrelloQuerySet(gGormDB).
		One(trello)
	return
}

func GetAllTrello(queryPage *modules.QueryPage) (trellos []Trello, err error) {
	err = NewTrelloQuerySet(gGormDB).
		All(&trellos)
	return
}

func UpdateTrelloByID(trello *Trello) (err error) {
	err = trello.Update(gGormDB,
		TrelloDBSchema.Name,
		TrelloDBSchema.Description,
	)
	return
}

func DeleteTrello(id uint) (err error) {
	trello := &Trello{
		ID: id,
	}
	err = trello.Delete(gGormDB)
	return
}
