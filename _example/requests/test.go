package requests

import "time"

type CreateTest struct {
}

type UpdateTest struct {
}

func (c *CreateTest) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateTest) Valid() error {
	return validate.Struct(c)
}
