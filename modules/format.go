package modules

import (
	"fmt"
	"strings"
	"time"
)

const (
	_listCountText   = "list_count"
	_listDefaultText = "{{list}}"
	_nowText         = "now"
	_nowDayText      = "now_day"
	_nowMonthText    = "now_month"
	_nowYearText     = "now_year"
)

type _OldAndNewText struct {
	OldText string
	NewText string
}

type _KeyValueSet struct {
	KeyValues []_KeyValue
}

type _KeyValue struct {
	Key       string
	Value     string
	ValueType string
}

type _OldAndNewTextKeyValueSet struct {
	ListOldAndNewTexts []_OldAndNewText
}

type _TextFormator struct {
	OldAndNewTexts            []_OldAndNewText
	OldAndNewTextKeyValueSets []_OldAndNewTextKeyValueSet
	AfterOldAndNewTexts       []_OldAndNewText
}

func newDefaultFomator(newKeyValueSets []_KeyValueSet, replaceTexts ...string) *_TextFormator {
	newOldAndNewTexts := []_OldAndNewText{
		_OldAndNewText{_nowText, fmt.Sprintf("%s", time.Now())},
		_OldAndNewText{_nowDayText, fmt.Sprintf("%d", time.Now().Day())},
		_OldAndNewText{_nowMonthText, fmt.Sprintf("%d", time.Now().Month())},
		_OldAndNewText{_nowYearText, fmt.Sprintf("%d", time.Now().Year())},
	}

	afterOldAndNewTexts := []_OldAndNewText{}
	if replaceTexts != nil {
		for _, item := range replaceTexts {
			itemSplit := strings.Split(item, ":")
			if len(itemSplit) != 2 {
				continue
			}
			afterOldAndNewTexts = append(afterOldAndNewTexts, _OldAndNewText{
				OldText: itemSplit[0],
				NewText: itemSplit[1],
			})
		}
	}

	newOldAndNewTexts = append(newOldAndNewTexts, _OldAndNewText{
		OldText: _listCountText,
		NewText: fmt.Sprintf("%d", len(newKeyValueSets)),
	})

	newOldAndNewTextKeyValueSets := []_OldAndNewTextKeyValueSet{}
	for _, newKeyValueSet := range newKeyValueSets {
		newOldAndNewTextKeyValueSet := _OldAndNewTextKeyValueSet{
			ListOldAndNewTexts: newOldAndNewTexts,
		}

		for _, newKeyValue := range newKeyValueSet.KeyValues {
			newOldAndNewTextKeyValueSet.ListOldAndNewTexts = append(newOldAndNewTextKeyValueSet.ListOldAndNewTexts, _OldAndNewText{
				OldText: newKeyValue.Key,
				NewText: newKeyValue.Value,
			})
		}

		newOldAndNewTextKeyValueSets = append(newOldAndNewTextKeyValueSets, newOldAndNewTextKeyValueSet)
	}

	return &_TextFormator{
		OldAndNewTexts:            newOldAndNewTexts,
		OldAndNewTextKeyValueSets: newOldAndNewTextKeyValueSets,
		AfterOldAndNewTexts:       afterOldAndNewTexts,
	}
}

func (t *_TextFormator) ConvText(raw string) string {
	return replaceAll(raw, append(t.OldAndNewTexts, t.AfterOldAndNewTexts...))
}

func (t *_TextFormator) ConvTextAndList(raw string, rawKeyValueSet string, listOldTexts ...string) string {
	raw = replaceAll(raw, t.OldAndNewTexts)
	newKeyValueSets := ""

	for _, newOldAndNewTextKeyValueSet := range t.OldAndNewTextKeyValueSets {
		newKeyValueSet := replaceAll(rawKeyValueSet, newOldAndNewTextKeyValueSet.ListOldAndNewTexts)
		newKeyValueSet = replaceAll(newKeyValueSet, t.OldAndNewTexts)
		newKeyValueSets = newKeyValueSets + newKeyValueSet
	}

	listOldText := _listDefaultText
	if len(listOldTexts) >= 1 {
		listOldText = listOldTexts[0]
	}
	output := strings.Replace(raw, listOldText, newKeyValueSets, -1)

	return replaceAllOlnyText(output, t.AfterOldAndNewTexts)
}

func replaceAll(str string, newOldAndNewTexts []_OldAndNewText) string {
	for _, newOldAndNewText := range newOldAndNewTexts {
		str = strings.Replace(str, fmt.Sprintf("{{%s}}", newOldAndNewText.OldText), newOldAndNewText.NewText, -1)
	}
	return str
}

func replaceAllOlnyText(str string, newOldAndNewTexts []_OldAndNewText) string {
	for _, newOldAndNewText := range newOldAndNewTexts {
		str = strings.Replace(str, fmt.Sprintf("%s", newOldAndNewText.OldText), newOldAndNewText.NewText, -1)
	}
	return str
}
