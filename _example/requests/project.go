package requests

import "time"

type CreateProject struct {
	Status string
}

type UpdateProject struct {
	Status string
}

func (c *CreateProject) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateProject) Valid() error {
	return validate.Struct(c)
}
