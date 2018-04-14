package requests

import "time"

type CreateKeyEvent struct {
	EventDate time.Time
}

type UpdateKeyEvent struct {
	EventDate time.Time
}

func (c *CreateKeyEvent) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateKeyEvent) Valid() error {
	return validate.Struct(c)
}
