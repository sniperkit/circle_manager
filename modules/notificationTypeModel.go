package modules

import (
	"encoding/json"
	"fmt"
	"strings"
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
	IsEnable         bool      `description:""`
	IsManual         bool      `description:""`
	TargetObject     string    `description:""`
	TargetWhere      string    `description:""`
	Tags             string    `description:""`
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

func (c *NotificationType) IsMatchTags(tags string) bool {
	mapTag := map[string]bool{}
	for _, tag := range strings.Split(tags, ",") {
		mapTag[tag] = true
	}

	mapNotiTypeTags := map[string]bool{}
	for _, notiTypeTag := range strings.Split(c.Tags, ",") {
		mapNotiTypeTags[notiTypeTag] = true
	}

	for _, tag := range strings.Split(tags, ",") {
		if _, ok := mapNotiTypeTags[tag]; !ok {
			return false
		}
	}
	return true
}

func (m *NotificationType) CheckDiff(crudEvent *CrudEvent) bool {
	type UpdateProperty struct {
		Key      string
		OldValue string
		NewValue string
	}

	mapUpdateProperties := map[string]UpdateProperty{}

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

	for key, value := range mapUpdateItem {
		if structs.IsStruct(value) {
			continue
		}

		oldValue := ""
		if tempOldValue, ok := mapOldItem[key]; ok {
			oldValue = convInterface(tempOldValue)
		}

		mapUpdateProperties[key] = UpdateProperty{
			Key:      toDBName(key),
			NewValue: convInterface(value),
			OldValue: oldValue,
		}
	}

	if len(mapUpdateProperties) <= 0 {
		return false
	}

	updateProperty, ok := mapUpdateProperties[m.DiffKey]
	if !ok {
		return false
	}

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
