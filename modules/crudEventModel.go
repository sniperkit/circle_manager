package modules

import (
	"time"
)

// gen:qs
type CrudEvent struct {
	ID           uint      `description:""`
	CreatedAt    time.Time `description:"등록일"`
	UpdatedAt    time.Time `description:"수정일"`
	Name         string    `description:"이름"`
	Description  string    `description:"설명" sql:"type:text"`
	CreatorID    uint      `description:"작성자"`
	Action       string    `description:"CRUD 타입"`
	TargetObject string    `description:"이벤트 대상"`
	TargetID     uint      `description:"이벤트 대상 ID"`
	Where        string    `description:"이벤트 발생 위치"`
	UpdatedData  string    `description:"업데이트 된 대상 Data"`
	OldData      string    `description:"업데이트 되기 전 대상 Data"`
	RequestData  string    `description:"요청 Data"`
	ResponseData string    `description:"출력 Data"`
}

func (c *CrudEvent) GetCreatorID() uint {
	return c.CreatorID
}

func (m *CrudEvent) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func AddCrudEvent(crudEvent *CrudEvent) (id uint, err error) {
	err = crudEvent.Create(gGormDB)
	id = crudEvent.ID
	return
}