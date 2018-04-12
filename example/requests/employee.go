package requests

import "time"

type CreateEmployee struct {
	OriginName string
}

type UpdateEmployee struct {
	OriginName string
}

func (c *CreateEmployee) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateEmployee) Valid() error {
	return validate.Struct(c)
}
