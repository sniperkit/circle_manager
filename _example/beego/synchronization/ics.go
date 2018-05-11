package synchronization

import (
	"io/ioutil"
	"net/http"
	"time"

	ics "github.com/PuloV/ics-golang"
	"github.com/jungju/circle_manager/_example/beegoapp/models"
	"github.com/jungju/circle_manager/modules"
	"github.com/sirupsen/logrus"
)

func SyncICS(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	parser := ics.New()
	parser.Load(string(body))

	cals, err := parser.GetCalendars()
	if err != nil {
		return err
	}

	events := []ics.Event{}
	for _, cal := range cals {
		events = append(events, cal.GetEvents()...)
	}

	gotEvents := []models.Event{}
	if err := modules.GetItems(&gotEvents, nil); err != nil {
		return err
	}
	mapGotEvents := map[string]models.Event{}
	mapExistsEvents := map[string]bool{}
	for _, gotEvent := range gotEvents {
		mapGotEvents[gotEvent.EventID] = gotEvent
		mapExistsEvents[gotEvent.EventID] = false
	}

	for _, event := range events {
		var end *time.Time = nil
		if 2000 < event.GetEnd().Year() {
			getEnd := event.GetEnd()
			end = &getEnd
		}

		org := event.GetOrganizer()
		creatorUser := ""
		if org != nil {
			creatorUser = org.GetEmail()
		}

		attendees := event.GetAttendees()
		attendeesStr := ""
		for _, attendee := range attendees {
			if attendeesStr == "" {
				attendeesStr = attendee.GetEmail()
			} else {
				attendeesStr = attendeesStr + "," + attendee.GetEmail()
			}
		}

		eventID := event.GetID()
		if _, ok := mapExistsEvents[eventID]; ok {
			mapExistsEvents[eventID] = true
		}

		updateItem := &models.Event{
			Name:         event.GetSummary(),
			EventEnds:    end,
			Summary:      event.GetSummary(),
			Organizer:    creatorUser,
			EventCreated: event.GetCreated(),
			EventBegins:  event.GetStart(),
			EventID:      eventID,
			Location:     event.GetLocation(),
			Source:       "ics",
			Attendees:    attendeesStr,
		}

		if getItem, ok := mapGotEvents[eventID]; !ok {
			if err := modules.CreateItem(updateItem); err != nil {
				return err
			}
		} else {
			updateItem.ID = getItem.ID
			updateItem.CreatedAt = getItem.CreatedAt

			if !IsEqual(updateItem, getItem) {
				if err := modules.SaveItem(updateItem); err != nil {
					return err
				}
			}
		}
	}

	for uuidKey, exists := range mapExistsEvents {
		if !exists {
			if err := modules.DeleteItemByColName("events", "event_id", uuidKey); err != nil {
				logrus.WithError(err).Error()
			}
		}
	}

	return nil
}
