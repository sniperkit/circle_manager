package responses

import (
	"time"
)

type Employee struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	OriginName  string
}
