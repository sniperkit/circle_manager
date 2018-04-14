package responses

import "time"

type ResponseTodo struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	ListID      string
	ListName    string
	Status      string
	CardID      string
	BoardID     string
	BoardName   string
	Source      string
}
