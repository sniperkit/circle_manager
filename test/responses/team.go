package responses

import "time"

type ResponseTeam struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}
