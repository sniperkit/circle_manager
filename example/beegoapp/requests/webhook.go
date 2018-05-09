package requests

import "time"

type TrelloActivity struct {
	ListID      string
	ListName    string
	CreatorName string
	CardID      string
	CreateCard  string
	CardName    string
	Type        string
	BoardID     string
	BoardName   string
}

type CalendarEventGoogle struct {
	Description    string
	EventEnds      time.Time
	Summary        string
	CreatorEmail   string
	EventBegins    time.Time
	GoogleID       string
	Location       string
	OrganizerEmail string
}

func (*TrelloActivity) Valid() error {
	return nil
}

func (*CalendarEventGoogle) Valid() error {
	return nil
}
