package modules

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMapUpdatedItems(t *testing.T) {
	tm := NewTemplateKeyValueMaker(
		&CrudEvent{
			UpdatedData: `{"test":"t1","test_id":1}`,
		},
		&NotificationType{
			TitleTemplate:   "",
			MessageTemplate: "",
		})

	assert.Equal(t, 2, len(tm.MapUpdateItems))
	assert.Equal(t, "t1", tm.MapUpdateItems["test"])

	subIDFloat := tm.MapUpdateItems["test_id"].(float64)
	var subUint64 uint = uint(subIDFloat)

	assert.Equal(t, uint(1), subUint64)

}

func TestTemplateKeyValueMaker(t *testing.T) {
	crudEvent := &CrudEvent{
		ResourceID:   1111,
		ResourceName: "Product",
		UpdatedData:  `{"ID":1111,"Name":"GoodName","ContentID":1,"IsCheck":false}`,
	}

	notificationType := &NotificationType{
		TitleTemplate:    "title {{Product.Name}} {{Content.Name}} {{Product.IsCheck}}",
		MessageTemplate:  "message {{Content.Version}}",
		ListItemTemplate: "list {{Product.Name}}",
	}

	templateKeyValueMaker := NewTemplateKeyValueMaker(crudEvent, notificationType)
	getDBValueParams := templateKeyValueMaker.MakeGetValueParams()
	assert.Equal(t, 2, len(getDBValueParams))

	assert.Equal(t, "GoodName", templateKeyValueMaker.GetedValues["Name"])
	assert.Equal(t, "false", templateKeyValueMaker.GetedValues["IsCheck"])

	assert.Equal(t, uint(1), getDBValueParams["Content.Version"].ID)
	assert.Equal(t, "version", getDBValueParams["Content.Version"].ColumnName)
	assert.Equal(t, "contents", getDBValueParams["Content.Version"].TableName)

	assert.Equal(t, uint(1), getDBValueParams["Content.Name"].ID)
	assert.Equal(t, "name", getDBValueParams["Content.Name"].ColumnName)
	assert.Equal(t, "contents", getDBValueParams["Content.Name"].TableName)
}

func TestMakeNotification(t *testing.T) {
	baseGetedValues := map[string]string{
		"Key":  "base-test1-1",
		"Key1": "base-test1-2",
		"Key2": "base-3",
		"Key3": fmt.Sprintf("%s", time.Now()),
	}
	listValues := []KeyValue{
		map[string]string{
			"List.Key":  "list-test1-1",
			"List.Key1": "list-test1-2",
			"List.Key2": "3",
			"List.Key3": fmt.Sprintf("%s", time.Now()),
		},
		map[string]string{
			"List.Key":  "list-test2-1",
			"List.Key1": "list-test2-2",
			"List.Key2": "4",
			"List.Key3": fmt.Sprintf("%s", time.Now()),
		},
		map[string]string{
			"List.Key":  "list-test3-1",
			"List.Key1": "list-test3-2",
			"List.Key2": "5",
			"List.Key3": fmt.Sprintf("%s", time.Now()),
		},
	}

	// Template
	notiType := &NotificationType{
		TitleTemplate:    "title {{Key}} {{NowMonth}}",
		MessageTemplate:  "message {{Key}} {{Key1}}\n{{ListItem}}",
		ListItemTemplate: "- {{Key}} {{List.Key}} {{List.Key1}} {{List.Key2}}\n",
		ReplaceText:      "old:new,test1:newtest1",
	}

	notification := MakeNotification(notiType, baseGetedValues, listValues)

	assert.NotNil(t, notification)
	assert.Equal(t, fmt.Sprintf("title %s %d", baseGetedValues["Key"], time.Now().Month()), notification.Title)
	assert.Equal(t, "message base-test1-1 base-test1-2\n- base-test1-1 list-test1-1 list-test1-2 3\n- base-test1-1 list-test2-1 list-test2-2 4\n- base-test1-1 list-test3-1 list-test3-2 5\n", notification.Message)
}
