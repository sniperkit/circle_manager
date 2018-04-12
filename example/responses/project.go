package responses

import "time"

type ResponseProject struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Status      string
}
