package requests

import "time"

type CreateUser struct {
	Owner string
	CarID uint64
}

type UpdateUser struct {
	Owner string
	CarID uint64
}

func (c *CreateUser) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateUser) Valid() error {
	return validate.Struct(c)
}
