package modules

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMapUpdatedItems(t *testing.T) {
	crudEvent := CrudEvent{
		UpdatedData: `{"test":"t1","test_id":1}`,
	}
	ret := crudEvent.GetMapUpdatedItems()
	assert.Equal(t, 2, len(ret))
	assert.Equal(t, "t1", ret["test"])

	subIDFloat := ret["test_id"].(float64)
	var subUint64 uint = uint(subIDFloat)

	assert.Equal(t, uint(1), subUint64)

}

func TestTemplateKeyValueMaker(t *testing.T) {
	crudEvent := &CrudEvent{
		TargetID:     1111,
		TargetObject: "products",
		UpdatedData:  `{"ID":1111,"Name":"GoodName","ContentID":1}`,
	}

	notificationType := &NotificationType{
		TitleTemplate:    "title {{.name}} {{.contents__name}}",
		MessageTemplate:  "message {{.contents__version}}",
		ListItemTemplate: "list {{.name}}",
	}

	templateKeyValueMaker := NewTemplateKeyValueMaker(crudEvent, notificationType)
	templateKeyValueMaker.MakeGetValueParams()

	for _, dd := range templateKeyValueMaker.getValueParams {
		fmt.Println(dd)
	}
	assert.Equal(t, 3, len(templateKeyValueMaker.getValueParams))
	assert.Equal(t, "GoodName", templateKeyValueMaker.getValueParams["name"].Value)
	assert.Equal(t, crudEvent.TargetID, templateKeyValueMaker.getValueParams["name"].ID)
	assert.Equal(t, "name", templateKeyValueMaker.getValueParams["name"].Key)
	assert.Equal(t, crudEvent.TargetObject, templateKeyValueMaker.getValueParams["name"].TableName)

	assert.Nil(t, templateKeyValueMaker.getValueParams["contents__version"].Value)
	assert.Equal(t, uint(1), templateKeyValueMaker.getValueParams["contents__version"].ID)
	assert.Equal(t, "version", templateKeyValueMaker.getValueParams["contents__version"].Key)
	assert.Equal(t, "contents", templateKeyValueMaker.getValueParams["contents__version"].TableName)

	assert.Nil(t, templateKeyValueMaker.getValueParams["contents__name"].Value)
	assert.Equal(t, uint(1), templateKeyValueMaker.getValueParams["contents__name"].ID)
	assert.Equal(t, "name", templateKeyValueMaker.getValueParams["contents__name"].Key)
	assert.Equal(t, "contents", templateKeyValueMaker.getValueParams["contents__name"].TableName)
}

func TestMakeNotification(t *testing.T) {
	baseGetedValues := map[string]string{
		"key":  "base-test1-1",
		"key1": "base-test1-2",
		"key2": "base-3",
		"key3": fmt.Sprintf("%s", time.Now()),
	}
	listValues := []KeyValue{
		map[string]string{
			"list__key":  "list-test1-1",
			"list__key1": "list-test1-2",
			"list__key2": "3",
			"list__key3": fmt.Sprintf("%s", time.Now()),
		},
		map[string]string{
			"list__key":  "list-test2-1",
			"list__key1": "list-test2-2",
			"list__key2": "4",
			"list__key3": fmt.Sprintf("%s", time.Now()),
		},
		map[string]string{
			"list__key":  "list-test3-1",
			"list__key1": "list-test3-2",
			"list__key2": "5",
			"list__key3": fmt.Sprintf("%s", time.Now()),
		},
	}

	// Template
	notiType := &NotificationType{
		TitleTemplate:    "title {{.key}} {{.nowMonth}}",
		MessageTemplate:  "message {{.key}} {{.key1}}\n{{.listItem}}",
		ListItemTemplate: "- {{.key}} {{.list__key}} {{.list__key1}} {{.list__key2}}\n",
		ReplaceText:      "old:new,test1:newtest1",
	}

	notification := MakeNotification(notiType, baseGetedValues, listValues)

	assert.NotNil(t, notification)
	assert.Equal(t, fmt.Sprintf("title %s %d", baseGetedValues["key"], time.Now().Month()), notification.Title)
	assert.Equal(t, "message base-test1-1 base-test1-2\n- base-test1-1 list-test1-1 list-test1-2 3\n- base-test1-1 list-test2-1 list-test2-2 4\n- base-test1-1 list-test3-1 list-test3-2 5\n", notification.Message)
}
