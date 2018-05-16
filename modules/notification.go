package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	_listCountText   = "ConstListCount"
	_listDefaultText = "ConstListItem"
	_nowText         = "ConstNow"
	_nowDayText      = "ConstNowDay"
	_nowMonthText    = "ConstNowMonth"
	_nowYearText     = "ConstNowYear"
)

type KeyValue map[string]string

type TemplateKeyValueMaker struct {
	CrudEvent       *CrudEvent
	MapUpdateItems  map[string]interface{}
	TemplateStrings []string
	GetedValues     map[string]string
}

func NewTemplateKeyValueMaker(crudEvent *CrudEvent, notificationType *NotificationType) *TemplateKeyValueMaker {
	mapUpdateItems := map[string]interface{}{}
	if crudEvent.UpdatedData != "" {
		if err := json.Unmarshal([]byte(crudEvent.UpdatedData), &mapUpdateItems); err != nil {
			fmt.Println(err)
		}
	}

	return &TemplateKeyValueMaker{
		CrudEvent:      crudEvent,
		MapUpdateItems: mapUpdateItems,
		TemplateStrings: []string{
			notificationType.TitleTemplate,
			notificationType.MessageTemplate,
		},
		GetedValues: map[string]string{},
	}
}

func (m TemplateKeyValueMaker) GetForienKeyValue(objectName string) (uint, error) {
	tableKey := fmt.Sprintf("%sID", objectName)

	if subIDInterface, ok := m.MapUpdateItems[tableKey]; ok {
		if subIDFloat, ok := subIDInterface.(float64); ok {
			var subUint64 uint = uint(subIDFloat)
			return subUint64, nil
		}
	}
	return 0, errors.New("Not found column")
}

func (m *TemplateKeyValueMaker) LoadValues(getDBValueParams map[string]ParamGetValueByKeyOfTableName) {
	m.GetedValues = map[string]string{}
	for key, getValueParam := range getDBValueParams {
		if valueInterface, err := GetValueByKeyOfTableName(getValueParam.TableName, getValueParam.ColumnName, getValueParam.ID); err == nil {
			m.GetedValues[key] = convInterface(valueInterface)
		} else {
			fmt.Println(err)
			continue
		}
	}
}

func (m *TemplateKeyValueMaker) MakeGetValueParams() map[string]ParamGetValueByKeyOfTableName {
	matchs := []string{}
	for _, str := range m.TemplateStrings {
		re := regexp.MustCompile("\\{\\{\\.(.*?)\\}\\}")
		allMatch := re.FindAllStringSubmatch(str, -1)
		for _, match := range allMatch {
			matchs = append(matchs, match[1:len(match)]...)
		}
	}

	getDBValueParams := map[string]ParamGetValueByKeyOfTableName{}
	for _, match := range matchs {
		objectAndKey := strings.Split(match, ".")

		if len(objectAndKey) == 1 {
			fieldName := objectAndKey[0]
			m.GetedValues[match] = convInterface(m.MapUpdateItems[fieldName])
		} else if len(objectAndKey) == 2 {
			subID, err := m.GetForienKeyValue(objectAndKey[0])
			if err != nil {
				fmt.Println("알수없는 ObjectID : ", objectAndKey)
			}

			getDBValueParams[match] = ParamGetValueByKeyOfTableName{
				TableName:  toDBTableName(objectAndKey[0]),
				ID:         subID,
				ColumnName: toDBName(objectAndKey[1]),
			}
		} else {
			fmt.Println("알수없는 Object : ", objectAndKey)
		}
	}
	return getDBValueParams
}

func MakeNotification(notificationType *NotificationType, baseGetedValues KeyValue, listValuesGroup []KeyValue) (notification *Notification) {
	builtinKeyValues := getBuiltinKeyValues()
	for key, value := range baseGetedValues {
		builtinKeyValues[key] = value
	}
	creatingTitleText := changeTeamplate(notificationType.TitleTemplate, builtinKeyValues)
	creatingMessageText := changeTeamplate(notificationType.MessageTemplate, builtinKeyValues)
	creatingListTemplate := changeTeamplate(notificationType.ListItemTemplate, builtinKeyValues)

	if creatingListTemplate != "" &&
		len(listValuesGroup) > 0 &&
		strings.Index(creatingMessageText, _listDefaultText) >= 0 {
		list := ""
		for _, listValues := range listValuesGroup {
			list += changeTeamplate(creatingListTemplate, listValues)
		}
		creatingMessageText = changeTeamplate(creatingMessageText, KeyValue{_listDefaultText: list})
	}

	return &Notification{
		Title:              creatingTitleText,
		Message:            creatingMessageText,
		NotificationType:   *notificationType,
		NotificationTypeID: notificationType.ID,
	}
}

func getBuiltinKeyValues() KeyValue {
	return KeyValue{
		_nowText:      fmt.Sprintf("%s", time.Now()),
		_nowDayText:   fmt.Sprintf("%d", time.Now().Day()),
		_nowMonthText: fmt.Sprintf("%d", time.Now().Month()),
		_nowYearText:  fmt.Sprintf("%d", time.Now().Year()),
	}
}

func changeTeamplate(teamplate string, keyValue KeyValue) string {
	ret := teamplate

	for key, value := range keyValue {
		ret = strings.Replace(ret, fmt.Sprintf("{{.%s}}", key), convInterface(value), -1)
	}

	return ret
}
