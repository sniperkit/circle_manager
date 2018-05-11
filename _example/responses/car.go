package responses

import "time"

type ResponseCar struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Key1        string
}
