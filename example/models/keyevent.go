package models

import "time"

// gen:qs
type KeyEvent struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	EventDate   time.Time `description:""`
}

func init() {
	registModel(&KeyEvent{})
}

func AddKeyEvent(keyEvent *KeyEvent) (id uint, err error) {
	err = keyEvent.Create(gGormDB)
	id = keyEvent.ID
	return
}

func GetKeyEventByID(id uint) (keyEvent *KeyEvent, err error) {
	keyEvent = &KeyEvent{
		ID: id,
	}
	err = NewKeyEventQuerySet(gGormDB).
		One(keyEvent)
	return
}

func GetAllKeyEvent(queryPage *QueryPage) (keyEvents []KeyEvent, err error) {
	err = NewKeyEventQuerySet(gGormDB).
		All(&keyEvents)
	return
}

func UpdateKeyEventByID(keyEvent *KeyEvent) (err error) {
	err = keyEvent.Update(gGormDB,
		KeyEventDBSchema.Name,
		KeyEventDBSchema.Description,
		KeyEventDBSchema.EventDate,
	)
	return
}

func DeleteKeyEvent(id uint) (err error) {
	keyEvent := &KeyEvent{
		ID: id,
	}
	err = keyEvent.Delete(gGormDB)
	return
}
