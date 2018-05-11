package responses

import "time"

type Event struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Description  string
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
