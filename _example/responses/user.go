package responses

import "time"

type ResponseUser struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Owner       string
	CarID       uint64
}
