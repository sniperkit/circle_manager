package requests

import "github.com/jungju/circle_manager/example/beegoapp/errors"

type CreateAdmin struct {
	Username    string `validate:"required,min=2,max=36"`
	Password    string `validate:"required,min=6,max=36"`
	Description string ``
	Email       string `validate:"required"`
	Name        string `validate:"required,min=3,max=36"`
}

type UpdateAdmin struct {
	Username    string `validate:"required,min=2,max=36"`
	Password    string `validate:"required,min=6,max=36"`
	Description string ``
	Email       string `validate:"required"`
	Name        string `validate:"required,min=3,max=36"`
}

func (c *CreateAdmin) Valid() error {
	if c.Username == "system" || c.Username == "admin" {
		return errors.ErrInvalidUsername
	}
	return validate.Struct(c)
}

func (c *UpdateAdmin) Valid() error {
	if c.Username == "system" || c.Username == "admin" {
		return errors.ErrInvalidUsername
	}
	return validate.Struct(c)
}
