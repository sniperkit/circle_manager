package modules

import (
	"encoding/json"
	"fmt"
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
	Action              string    `description:"CRUD 타입"`
	TargetObject        string    `description:"이벤트 대상"`
	TargetID            uint      `description:"이벤트 대상 ID"`
	Where               string    `description:"이벤트 발생 위치"`
	UpdatedData         string    `description:"업데이트 된 대상 Data" sql:"type:text"`
	OldData             string    `description:"업데이트 되기 전 대상 Data" sql:"type:text"`
	RequestData         string    `description:"요청 Data" sql:"type:text"`
	ResponseData        string    `description:"출력 Data" sql:"type:text"`
	CheckedNotification bool      `description:"알림 체크"`
}

func (c *CrudEvent) GetCreatorID() uint {
	return c.CreatorID
}

func (m *CrudEvent) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func (m *CrudEvent) GetTags() string {
	return fmt.Sprintf("%s,%s", m.TargetObject, m.Action)
}

func (m *CrudEvent) GetMapUpdatedItems() map[string]interface{} {
	mapUpdateItems := map[string]interface{}{}
	if m.UpdatedData != "" {
		if err := json.Unmarshal([]byte(m.UpdatedData), &mapUpdateItems); err != nil {
			fmt.Println(err)
		}
	}

	retMapUpdateItems := map[string]interface{}{}
	for key, value := range mapUpdateItems {
		retMapUpdateItems[toDBName(key)] = value
	}
	return retMapUpdateItems
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
