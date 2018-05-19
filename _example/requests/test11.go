package requests

import "time"

type CreateTest11 struct {
}

type UpdateTest11 struct {
}

func (c *CreateTest11) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateTest11) Valid() error {
	return validate.Struct(c)
}
