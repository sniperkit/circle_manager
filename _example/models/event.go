package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// gen:qs
type Event struct {
	ID              uint          `description:""`
	CreatedAt       time.Time     `description:"등록일"`
	UpdatedAt       time.Time     `description:"수정일"`
	Name            string        `description:"이름"`
	Description     string        `description:"설명" sql:"type:text"`
	CreatorID       uint          `description:"작성자"`
	EventCreated    time.Time     `description:""`
	EventEnds       *time.Time    `description:""`
	Summary         string        `description:""`
	Organizer       string        `description:""`
	EventUser       string        `description:""`
	EventBegins     time.Time     `description:""`
	EventID         string        `description:""`
	Location        string        `description:""`
	Source          string        `description:""`
	Attendees       string        `description:""`
	GithubRelease   GithubRelease `description:""`
	GithubReleaseID uint          `description:""`
}

func init() {
	registModel(&Event{})
}

func (m *Event) GetCreatorID() uint {
	return m.CreatorID
}

func (m *Event) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddEvent(event *Event) (id uint, err error) {
	err = event.Create(gGormDB)
	id = event.ID
	return
}

func GetEventByID(id uint) (event *Event, err error) {
	event = &Event{
		ID: id,
	}
	err = NewEventQuerySet(gGormDB).
		One(event)
	return
}

func GetAllEvent(queryPage *modules.QueryPage) (events []Event, err error) {
	err = NewEventQuerySet(gGormDB).
		All(&events)
	return
}

func UpdateEventByID(event *Event) (err error) {
	err = event.Update(gGormDB,
		EventDBSchema.Name,
		EventDBSchema.Description,
		EventDBSchema.EventCreated,
		EventDBSchema.EventEnds,
		EventDBSchema.Summary,
		EventDBSchema.Organizer,
		EventDBSchema.EventUser,
		EventDBSchema.EventBegins,
		EventDBSchema.EventID,
		EventDBSchema.Location,
		EventDBSchema.Source,
		EventDBSchema.Attendees,
		EventDBSchema.GithubRelease,
		EventDBSchema.GithubReleaseID,
	)
	return
}

func DeleteEvent(id uint) (err error) {
	event := &Event{
		ID: id,
	}
	err = event.Delete(gGormDB)
	return
}
