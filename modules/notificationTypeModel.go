package modules

import (
	"time"
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

func GetNotificationsTypesByManualSend(isManual bool) (notificationTypes []NotificationType, err error) {
	err = NewNotificationTypeQuerySet(gGormDB).
		IsManualEq(isManual).
		IsEnableEq(true).
		All(&notificationTypes)
	return
}
