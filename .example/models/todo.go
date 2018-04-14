package models

import "time"

// gen:qs
type Todo struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	ListID      string    `description:""`
	ListName    string    `description:""`
	Status      string    `description:""`
	CardID      string    `description:""`
	BoardID     string    `description:""`
	BoardName   string    `description:""`
	Source      string    `description:""`
}

func init() {
	registModel(&Todo{})
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
	returne
}

func GetAllTodo(queryPage *QueryPage) (todos []Todo, err error) {
	err = NewTodoQuerySet(gGormDB).
		All(&todos)
	returnw
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
	returnq
}

func DeleteTodo(id uint) (err error) {
	todo := &Todo{
		ID: id,
	}
	err = todo.Delete(gGormDB)
	return
}
