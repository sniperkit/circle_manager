package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set EventQuerySet

// EventQuerySet is an queryset type for Event
type EventQuerySet struct {
	db *gorm.DB
}

// NewEventQuerySet constructs new EventQuerySet
func NewEventQuerySet(db *gorm.DB) EventQuerySet {
	return EventQuerySet{
		db: db.Model(&Event{}),
	}
}

func (qs EventQuerySet) w(db *gorm.DB) EventQuerySet {
	return NewEventQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) All(ret *[]Event) error {
	return qs.db.Find(ret).Error
}

// AttendeesEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) AttendeesEq(attendees string) EventQuerySet {
	return qs.w(qs.db.Where("attendees = ?", attendees))
}

// AttendeesIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) AttendeesIn(attendees string, attendeesRest ...string) EventQuerySet {
	iArgs := []interface{}{attendees}
	for _, arg := range attendeesRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("attendees IN (?)", iArgs))
}

// AttendeesNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) AttendeesNe(attendees string) EventQuerySet {
	return qs.w(qs.db.Where("attendees != ?", attendees))
}

// AttendeesNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) AttendeesNotIn(attendees string, attendeesRest ...string) EventQuerySet {
	iArgs := []interface{}{attendees}
	for _, arg := range attendeesRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("attendees NOT IN (?)", iArgs))
}

// Count is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Event) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) CreatedAtEq(createdAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) CreatedAtGt(createdAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) CreatedAtGte(createdAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) CreatedAtLt(createdAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) CreatedAtLte(createdAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) CreatedAtNe(createdAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Event) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) Delete() error {
	return qs.db.Delete(Event{}).Error
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionEq(description string) EventQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionIn(description string, descriptionRest ...string) EventQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionNe(description string) EventQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionNotIn(description string, descriptionRest ...string) EventQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// EventBeginsEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventBeginsEq(eventBegins time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_begins = ?", eventBegins))
}

// EventBeginsGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventBeginsGt(eventBegins time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_begins > ?", eventBegins))
}

// EventBeginsGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventBeginsGte(eventBegins time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_begins >= ?", eventBegins))
}

// EventBeginsLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventBeginsLt(eventBegins time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_begins < ?", eventBegins))
}

// EventBeginsLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventBeginsLte(eventBegins time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_begins <= ?", eventBegins))
}

// EventBeginsNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventBeginsNe(eventBegins time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_begins != ?", eventBegins))
}

// EventCreatedEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventCreatedEq(eventCreated time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_created = ?", eventCreated))
}

// EventCreatedGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventCreatedGt(eventCreated time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_created > ?", eventCreated))
}

// EventCreatedGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventCreatedGte(eventCreated time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_created >= ?", eventCreated))
}

// EventCreatedLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventCreatedLt(eventCreated time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_created < ?", eventCreated))
}

// EventCreatedLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventCreatedLte(eventCreated time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_created <= ?", eventCreated))
}

// EventCreatedNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventCreatedNe(eventCreated time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_created != ?", eventCreated))
}

// EventEndsEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsEq(eventEnds time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_ends = ?", eventEnds))
}

// EventEndsGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsGt(eventEnds time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_ends > ?", eventEnds))
}

// EventEndsGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsGte(eventEnds time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_ends >= ?", eventEnds))
}

// EventEndsIsNotNull is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsIsNotNull() EventQuerySet {
	return qs.w(qs.db.Where("event_ends IS NOT NULL"))
}

// EventEndsIsNull is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsIsNull() EventQuerySet {
	return qs.w(qs.db.Where("event_ends IS NULL"))
}

// EventEndsLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsLt(eventEnds time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_ends < ?", eventEnds))
}

// EventEndsLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsLte(eventEnds time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_ends <= ?", eventEnds))
}

// EventEndsNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventEndsNe(eventEnds time.Time) EventQuerySet {
	return qs.w(qs.db.Where("event_ends != ?", eventEnds))
}

// EventIDEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventIDEq(eventID string) EventQuerySet {
	return qs.w(qs.db.Where("event_id = ?", eventID))
}

// EventIDIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventIDIn(eventID string, eventIDRest ...string) EventQuerySet {
	iArgs := []interface{}{eventID}
	for _, arg := range eventIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("event_id IN (?)", iArgs))
}

// EventIDNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventIDNe(eventID string) EventQuerySet {
	return qs.w(qs.db.Where("event_id != ?", eventID))
}

// EventIDNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventIDNotIn(eventID string, eventIDRest ...string) EventQuerySet {
	iArgs := []interface{}{eventID}
	for _, arg := range eventIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("event_id NOT IN (?)", iArgs))
}

// EventUserEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventUserEq(eventUser string) EventQuerySet {
	return qs.w(qs.db.Where("event_user = ?", eventUser))
}

// EventUserIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventUserIn(eventUser string, eventUserRest ...string) EventQuerySet {
	iArgs := []interface{}{eventUser}
	for _, arg := range eventUserRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("event_user IN (?)", iArgs))
}

// EventUserNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventUserNe(eventUser string) EventQuerySet {
	return qs.w(qs.db.Where("event_user != ?", eventUser))
}

// EventUserNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) EventUserNotIn(eventUser string, eventUserRest ...string) EventQuerySet {
	iArgs := []interface{}{eventUser}
	for _, arg := range eventUserRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("event_user NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) GetUpdater() EventUpdater {
	return NewEventUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDEq(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDGt(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDGte(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDIn(ID uint, IDRest ...uint) EventQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDLt(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDLte(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDNe(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDNotIn(ID uint, IDRest ...uint) EventQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) Limit(limit int) EventQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// LocationEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) LocationEq(location string) EventQuerySet {
	return qs.w(qs.db.Where("location = ?", location))
}

// LocationIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) LocationIn(location string, locationRest ...string) EventQuerySet {
	iArgs := []interface{}{location}
	for _, arg := range locationRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("location IN (?)", iArgs))
}

// LocationNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) LocationNe(location string) EventQuerySet {
	return qs.w(qs.db.Where("location != ?", location))
}

// LocationNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) LocationNotIn(location string, locationRest ...string) EventQuerySet {
	iArgs := []interface{}{location}
	for _, arg := range locationRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("location NOT IN (?)", iArgs))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) NameEq(name string) EventQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) NameIn(name string, nameRest ...string) EventQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) NameNe(name string) EventQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) NameNotIn(name string, nameRest ...string) EventQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs EventQuerySet) One(ret *Event) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByCreatedAt() EventQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByEventBegins is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByEventBegins() EventQuerySet {
	return qs.w(qs.db.Order("event_begins ASC"))
}

// OrderAscByEventCreated is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByEventCreated() EventQuerySet {
	return qs.w(qs.db.Order("event_created ASC"))
}

// OrderAscByEventEnds is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByEventEnds() EventQuerySet {
	return qs.w(qs.db.Order("event_ends ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByID() EventQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByUpdatedAt() EventQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByCreatedAt() EventQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByEventBegins is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByEventBegins() EventQuerySet {
	return qs.w(qs.db.Order("event_begins DESC"))
}

// OrderDescByEventCreated is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByEventCreated() EventQuerySet {
	return qs.w(qs.db.Order("event_created DESC"))
}

// OrderDescByEventEnds is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByEventEnds() EventQuerySet {
	return qs.w(qs.db.Order("event_ends DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByID() EventQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByUpdatedAt() EventQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OrganizerEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrganizerEq(organizer string) EventQuerySet {
	return qs.w(qs.db.Where("organizer = ?", organizer))
}

// OrganizerIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrganizerIn(organizer string, organizerRest ...string) EventQuerySet {
	iArgs := []interface{}{organizer}
	for _, arg := range organizerRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("organizer IN (?)", iArgs))
}

// OrganizerNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrganizerNe(organizer string) EventQuerySet {
	return qs.w(qs.db.Where("organizer != ?", organizer))
}

// OrganizerNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrganizerNotIn(organizer string, organizerRest ...string) EventQuerySet {
	iArgs := []interface{}{organizer}
	for _, arg := range organizerRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("organizer NOT IN (?)", iArgs))
}

// SetAttendees is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetAttendees(attendees string) EventUpdater {
	u.fields[string(EventDBSchema.Attendees)] = attendees
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetCreatedAt(createdAt time.Time) EventUpdater {
	u.fields[string(EventDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetDescription(description string) EventUpdater {
	u.fields[string(EventDBSchema.Description)] = description
	return u
}

// SetEventBegins is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetEventBegins(eventBegins time.Time) EventUpdater {
	u.fields[string(EventDBSchema.EventBegins)] = eventBegins
	return u
}

// SetEventCreated is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetEventCreated(eventCreated time.Time) EventUpdater {
	u.fields[string(EventDBSchema.EventCreated)] = eventCreated
	return u
}

// SetEventEnds is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetEventEnds(eventEnds *time.Time) EventUpdater {
	u.fields[string(EventDBSchema.EventEnds)] = eventEnds
	return u
}

// SetEventID is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetEventID(eventID string) EventUpdater {
	u.fields[string(EventDBSchema.EventID)] = eventID
	return u
}

// SetEventUser is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetEventUser(eventUser string) EventUpdater {
	u.fields[string(EventDBSchema.EventUser)] = eventUser
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetID(ID uint) EventUpdater {
	u.fields[string(EventDBSchema.ID)] = ID
	return u
}

// SetLocation is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetLocation(location string) EventUpdater {
	u.fields[string(EventDBSchema.Location)] = location
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetName(name string) EventUpdater {
	u.fields[string(EventDBSchema.Name)] = name
	return u
}

// SetOrganizer is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetOrganizer(organizer string) EventUpdater {
	u.fields[string(EventDBSchema.Organizer)] = organizer
	return u
}

// SetSource is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetSource(source string) EventUpdater {
	u.fields[string(EventDBSchema.Source)] = source
	return u
}

// SetSummary is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetSummary(summary string) EventUpdater {
	u.fields[string(EventDBSchema.Summary)] = summary
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetUpdatedAt(updatedAt time.Time) EventUpdater {
	u.fields[string(EventDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SourceEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SourceEq(source string) EventQuerySet {
	return qs.w(qs.db.Where("source = ?", source))
}

// SourceIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SourceIn(source string, sourceRest ...string) EventQuerySet {
	iArgs := []interface{}{source}
	for _, arg := range sourceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("source IN (?)", iArgs))
}

// SourceNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SourceNe(source string) EventQuerySet {
	return qs.w(qs.db.Where("source != ?", source))
}

// SourceNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SourceNotIn(source string, sourceRest ...string) EventQuerySet {
	iArgs := []interface{}{source}
	for _, arg := range sourceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("source NOT IN (?)", iArgs))
}

// SummaryEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SummaryEq(summary string) EventQuerySet {
	return qs.w(qs.db.Where("summary = ?", summary))
}

// SummaryIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SummaryIn(summary string, summaryRest ...string) EventQuerySet {
	iArgs := []interface{}{summary}
	for _, arg := range summaryRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("summary IN (?)", iArgs))
}

// SummaryNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SummaryNe(summary string) EventQuerySet {
	return qs.w(qs.db.Where("summary != ?", summary))
}

// SummaryNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) SummaryNotIn(summary string, summaryRest ...string) EventQuerySet {
	iArgs := []interface{}{summary}
	for _, arg := range summaryRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("summary NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u EventUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u EventUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) UpdatedAtEq(updatedAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) UpdatedAtGt(updatedAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) UpdatedAtGte(updatedAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) UpdatedAtLt(updatedAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) UpdatedAtLte(updatedAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) UpdatedAtNe(updatedAt time.Time) EventQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set EventQuerySet

// ===== BEGIN of Event modifiers

type eventDBSchemaField string

func (f eventDBSchemaField) String() string {
	return string(f)
}

// EventDBSchema stores db field names of Event
var EventDBSchema = struct {
	ID           eventDBSchemaField
	CreatedAt    eventDBSchemaField
	UpdatedAt    eventDBSchemaField
	Name         eventDBSchemaField
	Description  eventDBSchemaField
	EventCreated eventDBSchemaField
	EventEnds    eventDBSchemaField
	Summary      eventDBSchemaField
	Organizer    eventDBSchemaField
	EventUser    eventDBSchemaField
	EventBegins  eventDBSchemaField
	EventID      eventDBSchemaField
	Location     eventDBSchemaField
	Source       eventDBSchemaField
	Attendees    eventDBSchemaField
}{

	ID:           eventDBSchemaField("id"),
	CreatedAt:    eventDBSchemaField("created_at"),
	UpdatedAt:    eventDBSchemaField("updated_at"),
	Name:         eventDBSchemaField("name"),
	Description:  eventDBSchemaField("description"),
	EventCreated: eventDBSchemaField("event_created"),
	EventEnds:    eventDBSchemaField("event_ends"),
	Summary:      eventDBSchemaField("summary"),
	Organizer:    eventDBSchemaField("organizer"),
	EventUser:    eventDBSchemaField("event_user"),
	EventBegins:  eventDBSchemaField("event_begins"),
	EventID:      eventDBSchemaField("event_id"),
	Location:     eventDBSchemaField("location"),
	Source:       eventDBSchemaField("source"),
	Attendees:    eventDBSchemaField("attendees"),
}

// Update updates Event fields by primary key
func (o *Event) Update(db *gorm.DB, fields ...eventDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":            o.ID,
		"created_at":    o.CreatedAt,
		"updated_at":    o.UpdatedAt,
		"name":          o.Name,
		"description":   o.Description,
		"event_created": o.EventCreated,
		"event_ends":    o.EventEnds,
		"summary":       o.Summary,
		"organizer":     o.Organizer,
		"event_user":    o.EventUser,
		"event_begins":  o.EventBegins,
		"event_id":      o.EventID,
		"location":      o.Location,
		"source":        o.Source,
		"attendees":     o.Attendees,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Event %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// EventUpdater is an Event updates manager
type EventUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewEventUpdater creates new Event updater
func NewEventUpdater(db *gorm.DB) EventUpdater {
	return EventUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Event{}),
	}
}

// ===== END of Event modifiers

// ===== END of all query sets
