package modules

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
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

// PostMenualMessage ...
// @Title PostMenualMessage
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

	notificationTypes, err := GetNotificationsTypes(true)
	if err != nil {
		c.ErrorAbort(500, err)
	}

	for _, notificationType := range notificationTypes {
		if !notificationType.IsMatchTags(tags) {
			continue
		}

		listValueGroup := []KeyValue{}
		if notificationType.TargetObject != "" {
			var err error
			rows, err := GetRows(notificationType.TargetObject, notificationType.TargetWhere)
			if err != nil {
				logrus.WithError(err).Error()
				continue
			}
			for _, row := range rows {
				listValues := KeyValue{}
				for key, value := range row {
					listValues[fmt.Sprintf("list_%s", key)] = convInterface(value)
				}
				listValueGroup = append(listValueGroup, listValues)
			}
		}

		notification := MakeNotification(&notificationType, nil, listValueGroup)

		if err := addNotificationAndSendNotification(notification); err != nil {
			logrus.WithError(err).Error()
			continue
		}
	}

	c.Success(http.StatusNoContent, nil)
}

func SendActiveNotifications() error {
	crudEvents, err := GetCrudEventaByCheckedNotification(false)
	if err != nil {
		return err
	}

	notificationTypes, err := GetNotificationsTypes(false)
	if err != nil {
		return err
	}

	mapCheckedCrudEvent := []uint{}
	for _, crudEvent := range crudEvents {
		if err := sendActiveNotificationsEachCrudEvent(&crudEvent, notificationTypes); err == nil {
			mapCheckedCrudEvent = append(mapCheckedCrudEvent, crudEvent.ID)
		}
	}

	if err := UpdateCheckedNotificationByCrudEventIDs(mapCheckedCrudEvent, true); err != nil {
		fmt.Println(err)
	}

	return nil
}

func sendActiveNotificationsEachCrudEvent(crudEvent *CrudEvent, notificationTypes []NotificationType) error {
	for _, notificationType := range notificationTypes {
		if !notificationType.IsMatchTags(crudEvent.GetTags()) {
			continue
		} else if notificationType.DiffMode {
			if !notificationType.CheckDiff(crudEvent) {
				continue
			}
		}

		templateKeyValueMaker := NewTemplateKeyValueMaker(crudEvent, &notificationType)
		getDBValueParams := templateKeyValueMaker.MakeGetValueParams()
		templateKeyValueMaker.LoadValues(getDBValueParams)

		notification := MakeNotification(&notificationType, templateKeyValueMaker.GetedValues, nil)
		notification.EventUserID = crudEvent.CreatorID

		if err := addNotificationAndSendNotification(notification); err != nil {
			logrus.Error(err)
		}
	}
	return nil
}

func addNotificationAndSendNotification(notification *Notification) error {
	//TODO: Warning! 전송 후 DB에러 시 다시 노티 전송 될 수 있음

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

	if _, err := AddNotification(notification); err != nil {
		return err
	}

	return nil
}
