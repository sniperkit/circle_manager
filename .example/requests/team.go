package requests

import "time"

type CreateTeam struct {
}

type UpdateTeam struct {
}

func (c *CreateTeam) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateTeam) Valid() error {
	return validate.Struct(c)
}
