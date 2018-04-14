package responses

import "time"

type ResponseEmployee struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	OriginName  string
}
