package responses

import "time"

type KeyEvent struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	EventDate   time.Time
}
