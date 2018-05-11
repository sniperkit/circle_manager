package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

var _ = time.Time{}

// gen:qs
type Todo struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	CreatorID   uint      `description:"작성자"`
	ListID      string
	ListName    string
	Status      string
	CardID      string
	BoardID     string
	BoardName   string
	Source      string
}

func init() {
	registModel(&Todo{})
}

func (m *Todo) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Todo) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddTodo(todo *Todo) (id uint, err error) {
	err = todo.Create(gGormDB)
	id = todo.ID
	return
}

func GetTodoByID(id uint) (todo *Todo, err error) {
	todo = &Todo{
		ID: id,
	}
	err = NewTodoQuerySet(gGormDB).
		One(todo)
	return
}

func GetAllTodo(queryPage *modules.QueryPage) (todos []Todo, err error) {
	err = NewTodoQuerySet(gGormDB).
		All(&todos)
	return
}

func UpdateTodoByID(todo *Todo) (err error) {
	err = todo.Update(gGormDB,
		TodoDBSchema.Name,
		TodoDBSchema.Description,
		TodoDBSchema.ListID,
		TodoDBSchema.ListName,
		TodoDBSchema.Status,
		TodoDBSchema.CardID,
		TodoDBSchema.BoardID,
		TodoDBSchema.BoardName,
		TodoDBSchema.Source,
	)
	return
}

func DeleteTodo(id uint) (err error) {
	todo := &Todo{
		ID: id,
	}
	err = todo.Delete(gGormDB)
	return
}
