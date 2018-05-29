package modules

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/structs"
)

// gen:qs
type NotificationType struct {
	ID               uint      `description:""`
	CreatedAt        time.Time `description:"등록일"`
	UpdatedAt        time.Time `description:"수정일"`
	Name             string    `description:"이름"`
	Description      string    `description:"설명" sql:"type:text"`
	CreatorID        uint      `description:"작성자"`
	UpdaterID        uint      `description:"최종수정자"`
	Group            string    `description:""`
	IsEnable         bool      `description:""`
	IsManual         bool      `description:""`
	ActionName       string    `description:"Action 이름"`
	ActionType       string    `description:"CRUD 타입"`
	ResourceName     string    `description:"이벤트 대상"`
	ResourceID       uint      `description:"이벤트 대상 ID"`
	TargetWhere      string    `description:""`
	TitleTemplate    string    `description:""`
	MessageTemplate  string    `description:"" gorm:"size:2500"`
	ListItemTemplate string    `description:"" gorm:"size:2500"`
	WebhookURLs      string    `description:"" gorm:"size:2500"`
	ReplaceText      string    `description:""`
	DiffMode         bool      `description:""`
	DiffKey          string    `description:""`
	DiffNewValue     string    `description:""`
	DiffOldValue     string    `description:""`
}

func (c *NotificationType) GetCreatorID() uint {
	return c.CreatorID
}

func (m *NotificationType) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func (m *NotificationType) SetUpdaterID(updaterID uint) {
	m.UpdaterID = updaterID
}

func (m *NotificationType) CheckDiff(crudEvent *CrudEvent) bool {
	if m.DiffKey == "" {
		return true
	}

	type UpdateProperty struct {
		Key      string
		OldValue string
		NewValue string
	}

	mapUpdateItem := map[string]interface{}{}
	if err := json.Unmarshal([]byte(crudEvent.UpdatedData), &mapUpdateItem); err != nil {
		fmt.Println(err)
		return false
	}
	mapOldItem := map[string]interface{}{}
	if err := json.Unmarshal([]byte(crudEvent.OldData), &mapOldItem); err != nil {
		fmt.Println(err)
		return false
	}

	mapUpdateProperties := map[string]UpdateProperty{}
	for key, value := range mapUpdateItem {
		if !structs.IsStruct(value) {
			oldValue := ""
			if tempOldValue, ok := mapOldItem[key]; ok {
				oldValue = convInterface(tempOldValue)
			}

			mapUpdateProperties[key] = UpdateProperty{
				Key:      ToDBName(key),
				NewValue: convInterface(value),
				OldValue: oldValue,
			}
		}
	}

	if len(mapUpdateProperties) <= 0 {
		return false
	}

	if updateProperty, ok := mapUpdateProperties[m.DiffKey]; ok {
		if m.DiffNewValue != "" {
			if m.DiffNewValue != "" && m.DiffNewValue != updateProperty.NewValue {
				return false
			}
		}

		if m.DiffOldValue != "" {
			if m.DiffOldValue != "" && m.DiffOldValue != updateProperty.OldValue {
				return false
			}
		}

		return updateProperty.NewValue != updateProperty.OldValue
	}

	return false
}

func AddNotificationType(notificationType *NotificationType) (id uint, err error) {
	err = notificationType.Create(gGormDB)
	id = notificationType.ID
	return
}

func GetNotificationTypeByID(id uint) (notificationType *NotificationType, err error) {
	notificationType = &NotificationType{
		ID: id,
	}
	err = NewNotificationTypeQuerySet(gGormDB).One(notificationType)
	return
}

func GetAllNotificationType(queryPage *QueryPage) (notificationTypes []NotificationType, err error) {
	err = NewNotificationTypeQuerySet(gGormDB).All(&notificationTypes)
	return
}

func UpdateNotificationTypeByID(notificationType *NotificationType) (err error) {
	err = notificationType.Update(gGormDB,
		NotificationTypeDBSchema.Description,
	)
	return
}

func DeleteNotificationType(id uint) (err error) {
	notificationType := &NotificationType{
		ID: id,
	}
	err = notificationType.Delete(gGormDB)
	return
}

func GetNotificationsTypes(isManual bool) (notificationTypes []NotificationType, err error) {
	err = NewNotificationTypeQuerySet(gGormDB).
		IsManualEq(isManual).
		IsEnableEq(true).
		All(&notificationTypes)
	return
}
