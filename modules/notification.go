package modules

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/inflection"
)

const (
	_listCountText   = "listCount"
	_listDefaultText = "listItem"
	_nowText         = "now"
	_nowDayText      = "nowDay"
	_nowMonthText    = "nowMonth"
	_nowYearText     = "nowYear"
)

type KeyValue map[string]string

type TemplateKeyValueMaker struct {
	CrudEvent       *CrudEvent
	MapUpdateItems  map[string]interface{}
	TemplateStrings []string
	getValueParams  map[string]ParamGetValueByKeyOfTableName
	GetedValues     map[string]string
}

func NewTemplateKeyValueMaker(crudEvent *CrudEvent, notificationType *NotificationType) *TemplateKeyValueMaker {
	return &TemplateKeyValueMaker{
		CrudEvent:      crudEvent,
		MapUpdateItems: crudEvent.GetMapUpdatedItems(),
		TemplateStrings: []string{
			notificationType.TitleTemplate,
			notificationType.MessageTemplate,
		},
	}
}

func (m TemplateKeyValueMaker) GetIDOfTable(subTableName string) (uint, error) {
	tableKey := "id"
	if subTableName != m.CrudEvent.TargetObject {
		tableKey = fmt.Sprintf("%s_id", inflection.Singular(subTableName))
	}

	if subIDInterface, ok := m.MapUpdateItems[tableKey]; ok {
		if subIDFloat, ok := subIDInterface.(float64); ok {
			var subUint64 uint = uint(subIDFloat)
			return subUint64, nil
		}
	}
	return 0, errors.New("Not found column")
}

func (m *TemplateKeyValueMaker) LoadValues() {
	m.GetedValues = map[string]string{}
	for key, getValueParam := range m.getValueParams {
		if getValueParam.Value != nil {
			m.GetedValues[key] = convInterface(getValueParam.Value)
			continue
		}
		if valueInterface, err := GetValueByKeyOfTableName(getValueParam.TableName, getValueParam.Key, getValueParam.ID); err == nil {
			m.GetedValues[key] = convInterface(valueInterface)
		} else {
			fmt.Println(err)
			continue
		}
	}
}

func (m *TemplateKeyValueMaker) MakeGetValueParams() {
	matchs := []string{}
	for _, str := range m.TemplateStrings {
		re := regexp.MustCompile("\\{\\{\\.(.*?)\\}\\}")
		allMatch := re.FindAllStringSubmatch(str, -1)
		for _, match := range allMatch {
			matchs = append(matchs, match[1:len(match)]...)
		}
	}

	m.getValueParams = map[string]ParamGetValueByKeyOfTableName{}
	for _, match := range matchs {
		objectAndKey := strings.Split(match, "__")

		if len(objectAndKey) == 2 {
			if subID, err := m.GetIDOfTable(objectAndKey[0]); err == nil {
				m.getValueParams[match] = ParamGetValueByKeyOfTableName{
					TableName: objectAndKey[0],
					ID:        subID,
					Key:       objectAndKey[1],
				}
			}
		} else if len(objectAndKey) == 1 {
			m.getValueParams[match] = ParamGetValueByKeyOfTableName{
				TableName: m.CrudEvent.TargetObject,
				ID:        m.CrudEvent.TargetID,
				Key:       objectAndKey[0],
				Value:     m.MapUpdateItems[objectAndKey[0]],
			}
		}
	}
}

func MakeNotification(notificationType *NotificationType, baseGetedValues KeyValue, listValuesGroup []KeyValue) (notification *Notification) {
	builtinKeyValues := getBuiltinKeyValues()
	for key, value := range baseGetedValues {
		builtinKeyValues[key] = value
	}
	creatingTitleText := changTeamplate(notificationType.TitleTemplate, builtinKeyValues)
	creatingMessageText := changTeamplate(notificationType.MessageTemplate, builtinKeyValues)
	creatingListTemplate := changTeamplate(notificationType.ListItemTemplate, builtinKeyValues)

	if creatingListTemplate != "" &&
		len(listValuesGroup) > 0 &&
		strings.Index(creatingMessageText, _listDefaultText) >= 0 {
		list := ""
		for _, listValues := range listValuesGroup {
			list += changTeamplate(creatingListTemplate, listValues)
		}
		creatingMessageText = changTeamplate(creatingMessageText, KeyValue{_listDefaultText: list})
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

func changTeamplate(teamplate string, keyValue KeyValue) string {
	ret := teamplate

	for key, value := range keyValue {
		ret = strings.Replace(ret, fmt.Sprintf("{{.%s}}", key), convInterface(value), -1)
	}

	return ret
}
