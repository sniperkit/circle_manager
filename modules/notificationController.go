package modules

import (
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"

	"github.com/jungju/goreport"
)

var notiManager goreport.NotiManager

//  NotificationController operations for Notification
type NotificationController struct {
	BaseCircleController
}

func init() {
	notiManager = goreport.NotiManager{
		GetRowFunc:             GetRows,
		UpdateSentNotification: UpdateSentNotification,
		AddNotification: func(gorNotification *goreport.Notification, isSent bool) error {
			notification := &Notification{}
			copier.Copy(notification, &gorNotification)
			notification.NotificationType = NotificationType{}

			if isSent {
				now := time.Now()
				notification.SentAt = &now
			}

			if _, err := AddNotification(notification); err != nil {
				return err
			}

			return nil
		},
	}
}

// PostMessage ...
// @Title PostMessage
// @Description create Notification
// @Param	body		body 	Notification	true		"body for Notification content"
// @Success 201 {int} Notification
// @Failure 403 body is empty
// @router /post [post]
func (c *NotificationController) PostMessage() {
	if err := SendActiveNotifications(); err != nil {
		c.ErrorAbort(500, err)
	}

	c.Success(http.StatusNoContent, nil)
}

func SendActiveNotifications() error {
	notifications, err := GetNotificationNoSent()
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		gorNotification := &goreport.Notification{}
		copier.Copy(gorNotification, &notification)

		if err := notiManager.Send(gorNotification); err != nil {
			logrus.Error(err)
			continue
		}
	}
	return nil
}

// PostMessage ...
// @Title PostMessage
// @Description create Notification
// @Param	body		body 	Notification	true		"body for Notification content"
// @Success 201 {int} Notification
// @Failure 403 body is empty
// @router /post/:key [post]
func (c *NotificationController) PostMenualMessage() {
	key := c.Ctx.Input.Param(":key")
	if key == "" {
		c.ErrorAbort(400, nil)
	}

	//TODO: key에 대한 처리

	tags := c.Ctx.Input.Query("tags")

	notificationTypes, err := GetNotificationsTypesByManualSend(true)
	if err != nil {
		c.ErrorAbort(500, err)
	}

	for _, notificationType := range notificationTypes {
		gorNotificationType := &goreport.NotificationType{}
		copier.Copy(gorNotificationType, &notificationType)

		if err := notiManager.SendManual(tags, gorNotificationType); err != nil {
			logrus.Error(err)
			continue
		}
	}

	c.Success(http.StatusNoContent, nil)
}

func AddActionNotification(tags string, objects ...interface{}) error {
	notificationTypes, err := GetNotificationsTypesByManualSend(false)
	if err != nil {
		return err
	}

	for _, notificationType := range notificationTypes {
		gorNotificationType := &goreport.NotificationType{}
		copier.Copy(gorNotificationType, &notificationType)

		if err := notiManager.AddActionNotification(tags, gorNotificationType); err != nil {
			logrus.Error(err)
			continue
		}
	}
	return nil
}
