package modules

// func TestMakeNotification(t *testing.T) {
// 	notiType := &NotificationType{
// 		ID:                      1,
// 		TargetObject:            "product",
// 		TargetAction:            "add",
// 		TitleTemplateForWebhook: "{{user.name}}/{{link}}",
// 	}

// 	notiTemplate := &NotificationTemplate{
// 		TargetObject: "content",
// 		TargetAction: "",
// 		TargetID:     1,
// 		UserID:       10,
// 		Username:     "테스트맨",
// 	}

// 	noti := makeNotification(notiType, notiTemplate)

// 	assert.Equal(t,
// 		fmt.Sprintf("%s/%s/%s/http://mypage/admin/%s/%d",
// 			notiTemplate.Username,
// 			notiTemplate.ProductName,
// 			notiTemplate.ContentVersion,
// 			inflection.Plural(notiTemplate.TargetObject),
// 			notiTemplate.TargetID), noti.TitleForWebhook)

// 	assert.Equal(t,
// 		fmt.Sprintf("http://mypage/admin/%s/%d",
// 			inflection.Plural(notiTemplate.TargetObject),
// 			notiTemplate.TargetID),
// 		noti.TargetLink)
// }

// func TestSendNotification(t *testing.T) {
// 	cnt := sendNotification(&Notification{
// 		//User: models.User{},
// 		NotificationType: NotificationType{
// 			PosibleSendWebhook:       true,
// 			SlackChannelIDForWebhook: "XXXXXX",
// 		},
// 		TitleForWebhook:   "나는 테스트입니다. 타이틀",
// 		MessageForWebhook: "나는 테스트입니다. 메세지~~~~~",
// 		TargetLink:        "http://test",
// 		TargetID:          1,
// 	})

// 	assert.Equal(t, 1, cnt)
// }
