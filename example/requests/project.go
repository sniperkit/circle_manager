package requests

import "time"

var _ = time.Time{}

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
