package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/structs"
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

		if err := addNotificationAndSendNotification(notification); err != nil {
			logrus.WithError(err).Error()
			continue
		}
	}

	c.Success(http.StatusNoContent, nil)
}

func SendActiveNotifications() error {
	crudEvents, err := GetAllCrudEventOlnyChekedNotification()
	if err != nil {
		return err
	}

	notificationTypes, err := GetNotificationsTypes(false)
	if err != nil {
		return err
	}

	for _, crudEvent := range crudEvents {
		tags := fmt.Sprintf("%s,%s", crudEvent.TargetObject, crudEvent.Action)
		mapUpdateProperties := makeMapUpdateProperties(crudEvent.UpdatedData, crudEvent.OldData)

		for _, notificationType := range notificationTypes {
			if !isExistsTag(tags, notificationType.Tags) {
				continue
			}

			if !checkDiff(mapUpdateProperties, notificationType) {
				continue
			}

			notification := MakeNotification(&notificationType, nil)
			notification.EventUserID = crudEvent.CreatorID
			notification.NotificationType = notificationType
			notification.NotificationTypeID = notificationType.ID

			if err := addNotificationAndSendNotification(notification); err != nil {
				logrus.Error(err)
			}
		}
		if err := UpdateCrudEventByNotification(crudEvent.ID); err != nil {
			logrus.Error(err)
		}
	}

	return nil
}

type UpdateProperty struct {
	Key      string
	OldValue string
	NewValue string
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

func checkDiff(mapUpdateProperties map[string]UpdateProperty, notificationType NotificationType) bool {
	if notificationType.DiffMode {
		if len(mapUpdateProperties) <= 0 {
			return false
		}

		updateProperty, ok := mapUpdateProperties[notificationType.DiffKey]
		if !ok {
			return false
		}

		if notificationType.DiffNewValue != "" {
			if notificationType.DiffNewValue != "" && notificationType.DiffNewValue != updateProperty.NewValue {
				return false
			}
		}

		if notificationType.DiffOldValue != "" {
			if notificationType.DiffOldValue != "" && notificationType.DiffOldValue != updateProperty.OldValue {
				return false
			}
		}

		return updateProperty.NewValue != updateProperty.OldValue
	}
	return true
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

	now := time.Now()
	notification.SentAt = &now

	if _, err := AddNotification(notification); err != nil {
		return err
	}

	return nil
}

func makeMapUpdateProperties(updatedData, oldData string) map[string]UpdateProperty {
	mapUpdateProperties := map[string]UpdateProperty{}

	if oldData == "" || updatedData == "" {
		return mapUpdateProperties
	}

	mapUpdateItem := map[string]interface{}{}
	if err := json.Unmarshal([]byte(updatedData), &mapUpdateItem); err != nil {
		return mapUpdateProperties
	}
	mapOldItem := map[string]interface{}{}
	if err := json.Unmarshal([]byte(oldData), &mapOldItem); err != nil {
		return mapUpdateProperties
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
			Key:      key,
			NewValue: convInterface(value),
			OldValue: oldValue,
		}
	}
	return mapUpdateProperties
}
