package modules

import (
	"time"
)

// gen:qs
type CrudEvent struct {
	ID                  uint      `description:""`
	CreatedAt           time.Time `description:"등록일"`
	UpdatedAt           time.Time `description:"수정일"`
	Name                string    `description:"이름"`
	Description         string    `description:"설명" sql:"type:text"`
	CreatorID           uint      `description:"작성자"`
	ActionName          string    `description:"Action 이름"`
	ActionType          string    `description:"CRUD 타입"`
	ResourceName        string    `description:"이벤트 대상"`
	ResourceID          uint      `description:"이벤트 대상 ID"`
	Where               string    `description:"이벤트 발생 위치"`
	UpdatedData         string    `description:"업데이트 된 대상 Data" sql:"type:text"`
	OldData             string    `description:"업데이트 되기 전 대상 Data" sql:"type:text"`
	RequestData         string    `description:"요청 Data" sql:"type:text"`
	ResponseData        string    `description:"출력 Data" sql:"type:text"`
	CheckedNotification bool      `description:"알림 체크"`
}

func AddCrudEvent(crudEvent *CrudEvent) (id uint, err error) {
	err = crudEvent.Create(gGormDB)
	id = crudEvent.ID
	return
}

func GetCrudEventaByCheckedNotification(checkedNotification bool) (crudEvents []CrudEvent, err error) {
	err = NewCrudEventQuerySet(gGormDB).
		CheckedNotificationEq(checkedNotification).
		All(&crudEvents)
	return
}

func UpdateCheckedNotificationByCrudEventIDs(ids []uint, checkedNotification bool) (err error) {
	err = NewCrudEventQuerySet(gGormDB).IDIn(0, ids...).
		GetUpdater().SetCheckedNotification(true).Update()
	return
}
