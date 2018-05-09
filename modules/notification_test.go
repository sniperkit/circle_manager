package modules

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// type NotificationTypeTest struct {
// 	TargetObject                    string `description:""`
// 	TargetWhere                     string `description:""`
// 	PosibleSendSMS                  bool   `description:""`
// 	PosibleSendEmail                bool   `description:""`
// 	PosibleSendWeb                  bool   `description:""`
// 	PosibleSendSlack                bool   `description:""`
// 	PosibleSendWebhook              bool   `description:""`
// 	UseLink                         bool   `description:""`
// 	TitleTemplate                   string `description:""`
// 	MessageTemplate                 string `description:""`
// 	ListItemTemplate                string `description:""`
// 	WebhookURLs                     string `description:""`
// 	SlackChannelIDForWebhook        string `description:""`
// 	SlackPrivateChannelIDForWebhook string `description:""`
// 	ReplaceText                     string `description:""`
// }

type _object1 struct {
	Name    string
	Desc    string
	TimeNow time.Time
	Number  int
}
type _object2 struct {
	Key  string
	Val  string
	Val2 time.Time
	Val3 int
}

func TestMakeNotificationByInterface(t *testing.T) {
	// DB rows
	rows := []map[string]interface{}{
		map[string]interface{}{
			"Key":  "test1-1",
			"Key1": "test1-2",
			"Key2": 3,
			"Key3": time.Now(),
		},
		map[string]interface{}{
			"Key":  "test2-1",
			"Key1": "test2-2",
			"Key2": 113,
			"Key3": time.Now(),
		},
	}

	// Template
	notiType := &NotificationType{
		TitleTemplate:    "title {{now_day}} {{now_month}}",
		MessageTemplate:  "message {{_object1__name}} {{_object2__val}}\n{{list}}",
		ListItemTemplate: "- {{key}} {{key1}} {{key2}}\n",
		ReplaceText:      "old:new,test1:newtest1",
	}

	//Objects
	object1 := &_object1{
		Name:    "object1Name", //Replace text : {{_object1__name}}
		Desc:    "object1Desc", //Replace text : {{_object1__desc}}
		TimeNow: time.Now(),    //Replace text : {{_object1__time_now}}
		Number:  10,            //Replace text : {{_object1__number}}
	}
	object2 := &_object2{
		Key:  "object2Name", //Replace text : {{_object2__key}}
		Val:  "object2Desc", //Replace text : {{_object2__val}}
		Val2: time.Now(),    //Replace text : {{_object2__val2}}
		Val3: 30,            //Replace text : {{_object2__val3}}
	}

	notification := MakeNotification(notiType, rows, object1, object2)

	assert.NotNil(t, notification)
	assert.Equal(t, fmt.Sprintf("title %d %d", time.Now().Day(), time.Now().Month()), notification.Title)
	assert.Equal(t, "message object1Name object2Desc\n- newtest1-1 newtest1-2 3\n- test2-1 test2-2 113\n", notification.Message)
}

// Before
// Title : title {{now_day}} {{now_month}}
// Message : message {{_object1__name}} {{_object2__val}}\n{{list}}
// ListItem : - {{key}} {{key1}} {{key2}} {{key3}}\n
// ReplaceText : "old:new,test1:newtest1"

// After(today 10.21)
// Title : title 10 21
// Message : message object1Name object2Desc\n- newtest1-1 newtest1-2 3 2018.3.22 18시\n- test2-1 test2-2 113 2018.3.22 18시\n
