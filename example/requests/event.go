package requests

import "time"

var _ = time.Time{}

type CreateEvent struct {
	EventCreated time.Time
	EventEnds    *time.Time
	Summary      string
	Organizer    string
	EventUser    string
	EventBegins  time.Time
	EventID      string
	Location     string
	Source       string
	Attendees    string
}

type UpdateEvent struct {
	EventCreated time.Time
	EventEnds    *time.Time
	Summary      string
	Organizer    string
	EventUser    string
	EventBegins  time.Time
	EventID      string
	Location     string
	Source       string
	Attendees    string
}

func (c *CreateEvent) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateEvent) Valid() error {
	return validate.Struct(c)
}
