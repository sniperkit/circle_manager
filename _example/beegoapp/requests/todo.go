package requests

import "time"

var _ = time.Time{}

type CreateTodo struct {
	ListID    string
	ListName  string
	Status    string
	CardID    string
	BoardID   string
	BoardName string
	Source    string
}

type UpdateTodo struct {
	ListID    string
	ListName  string
	Status    string
	CardID    string
	BoardID   string
	BoardName string
	Source    string
}

func (c *CreateTodo) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateTodo) Valid() error {
	return validate.Struct(c)
}
