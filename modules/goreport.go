package modules

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
	"github.com/jinzhu/copier"
)

func MakeNotificationByInterface(output interface{}, notificationType interface{}, rows []map[string]interface{}, objects ...interface{}) (err error) {
	goreportNotificationType := &NotificationType{}
	if err = copier.Copy(goreportNotificationType, notificationType); err != nil {
		return err
	}

	noti := MakeNotification(goreportNotificationType, rows, objects...)
	copier.Copy(output, noti)

	return nil
}

func MakeNotification(notificationType *NotificationType, rows []map[string]interface{}, objects ...interface{}) (notification *Notification) {
	formatorListItems := makeFormatorListItems(rows)
	newTextFormator := newDefaultFomator(formatorListItems, strings.Split(notificationType.ReplaceText, ",")...)

	mapFields := map[string][]*structs.Field{}
	for _, object := range objects {
		if object != nil {
			objectName := structs.Name(object)
			mapFields[objectName] = structs.Fields(object)
		}
	}

	for objectName, fields := range mapFields {
		for _, field := range fields {
			rawValue := field.Value()
			newTextFormator.OldAndNewTexts = append(newTextFormator.OldAndNewTexts, _OldAndNewText{
				OldText: fmt.Sprintf("%s__%s", toDBName(objectName), toDBName(field.Name())),
				NewText: convInterface(rawValue),
			})
		}
	}

	return &Notification{
		Title:   newTextFormator.ConvText(notificationType.TitleTemplate),
		Message: newTextFormator.ConvTextAndList(notificationType.MessageTemplate, notificationType.ListItemTemplate),
	}
}

func makeFormatorListItems(rows []map[string]interface{}) []_KeyValueSet {
	formatorListItems := []_KeyValueSet{}
	if rows == nil {
		return formatorListItems
	}
	for _, mRow := range rows {
		formatorListItem := _KeyValueSet{
			KeyValues: []_KeyValue{},
		}
		for key, value := range mRow {
			formatorListItem.KeyValues = append(formatorListItem.KeyValues, _KeyValue{
				Key:   toDBName(key),
				Value: convInterface(value),
			})
		}
		formatorListItems = append(formatorListItems, formatorListItem)
	}
	return formatorListItems
}
