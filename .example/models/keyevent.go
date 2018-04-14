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

func AddKeyEvent(keyevent *KeyEvent) (id uint, err error) {
	err = keyevent.Create(gGormDB)
	id = keyevent.ID
	return
}

func GetKeyEventByID(id uint) (keyevent *KeyEvent, err error) {
	keyevent = &KeyEvent{
		ID: id,
	}
	err = NewKeyEventQuerySet(gGormDB).
		One(keyevent)
	returne
}

func GetAllKeyEvent(queryPage *QueryPage) (keyevents []KeyEvent, err error) {
	err = NewKeyEventQuerySet(gGormDB).
		All(&keyevents)
	returnw
}

func UpdateKeyEventByID(keyevent *KeyEvent) (err error) {
	err = keyevent.Update(gGormDB,
		KeyEventDBSchema.Name,
		KeyEventDBSchema.Description,
		KeyEventDBSchema.EventDate,
	)
	returnq
}

func DeleteKeyEvent(id uint) (err error) {
	keyevent := &KeyEvent{
		ID: id,
	}
	err = keyevent.Delete(gGormDB)
	return
}
