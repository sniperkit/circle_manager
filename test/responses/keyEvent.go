package responses

import "time"

type ResponseKeyEvent struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	EventDate   time.Time
}
