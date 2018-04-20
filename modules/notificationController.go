package modules

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"

	"github.com/jungju/goreport"
)

//  NotificationController operations for Notification
type NotificationController struct {
	BaseUserController
}

// PostMessage ...
// @Title PostMessage
// @Description create Notification
// @Param	body		body 	modules.Notification	true		"body for Notification content"
// @Success 201 {int} modules.Notification
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
		if err := sendNotification(&notification); err != nil {
			logrus.Error(err)
			continue
		}
	}
	return nil
}

// PostMessage ...
// @Title PostMessage
// @Description create Notification
// @Param	body		body 	modules.Notification	true		"body for Notification content"
// @Success 201 {int} modules.Notification
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
		if !isExistsTag(tags, notificationType.Tags) {
			continue
		}

		gorNotificationType := &goreport.NotificationType{}
		copier.Copy(gorNotificationType, &notificationType)

		rows := []map[string]interface{}{}
		if notificationType.TargetObject != "" {
			var err error
			rows, err = GetRows(notificationType.TargetObject, notificationType.TargetWhere)
			if err != nil {
				logrus.WithError(err).Error()
				continue
			}
		}

		notification := MakeNotification(&notificationType, rows)
		notification.NotificationType = notificationType
		notification.NotificationTypeID = notificationType.ID

		if err := sendNotification(notification); err != nil {
			logrus.WithError(err).Error()
			continue
		}

		if _, err := AddNotification(notification); err != nil {
			logrus.WithError(err).Error()
			continue
		}
	}

	c.Success(http.StatusNoContent, nil)
}

func AddActionNotification(tags string, eventUserID *uint, objects ...interface{}) error {
	notificationTypes, err := GetNotificationsTypesByManualSend(false)
	if err != nil {
		return err
	}

	for _, notificationType := range notificationTypes {
		if !isExistsTag(tags, notificationType.Tags) {
			continue
		}

		notification := MakeNotification(&notificationType, nil, objects...)
		notification.EventUserID = eventUserID
		notification.NotificationType = notificationType
		notification.NotificationTypeID = notificationType.ID

		if _, err := AddNotification(notification); err != nil {
			return err
		}

		return nil
	}
	return nil
}

func isExistsTag(reqTags string, notiTypeTags string) bool {
	mapTag := map[string]bool{}
	for _, tag := range strings.Split(reqTags, ",") {
		mapTag[tag] = true
	}

	mapNotiTypeTags := map[string]bool{}
	for _, notiTypeTag := range strings.Split(notiTypeTags, ",") {
		mapNotiTypeTags[notiTypeTag] = true
	}

	for _, tag := range strings.Split(reqTags, ",") {
		if _, ok := mapNotiTypeTags[tag]; !ok {
			return false
		}
	}
	return true
}

func sendNotification(notification *Notification) error {
	if notification.NotificationType.WebhookURLs == "" {
		return nil
	}
	for _, targetURL := range strings.Split(notification.NotificationType.WebhookURLs, "\n") {
		webhookURL, err := url.Parse(strings.TrimSpace(targetURL))
		if err != nil {
			fmt.Printf("Error : %s", err.Error())
			continue
		}

		parameters := webhookURL.Query()
		parameters.Add("title", notification.Title)
		parameters.Add("message", notification.Message)
		webhookURL.RawQuery = parameters.Encode()

		if _, err := req("GET", webhookURL.String(), nil, nil, nil); err != nil {
			fmt.Printf("Error : %s", err.Error())
		}
	}

	if err := UpdateSentNotification(notification.ID); err != nil {
		return err
	}

	return nil
}
