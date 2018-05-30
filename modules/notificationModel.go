package modules

import (
	"time"
)

// gen:qs
type Notification struct {
	ID                 uint             `description:""`
	CreatedAt          time.Time        `description:"등록일"`
	UpdatedAt          time.Time        `description:"수정일"`
	Name               string           `description:"이름"`
	Description        string           `description:"설명" sql:"type:text"`
	CreatorID          uint             `description:"작성자"`
	EventUserID        uint             `description:""`
	NotificationType   NotificationType `description:""`
	NotificationTypeID uint             `description:""`
	NotiType           string           `description:""`
	Title              string           `description:""`
	Message            string           `description:"" gorm:"type:text"`
}

func AddNotification(notification *Notification) (id uint, err error) {
	err = notification.Create(gGormDB)
	id = notification.ID
	return
}

func GetNotificationByID(id uint) (notification *Notification, err error) {
	notification = &Notification{
		ID: id,
	}
	err = NewNotificationQuerySet(gGormDB).One(notification)
	return
}

func GetAllNotification(queryPage *QueryPage) (notifications []Notification, err error) {
	err = NewNotificationQuerySet(gGormDB).
		PreloadNotificationType().
		All(&notifications)
	return
}

func DeleteNotification(id uint) (err error) {
	notification := &Notification{
		ID: id,
	}
	err = notification.Delete(gGormDB)
	return
}
