package controllers

import (
	"net/http"

	"github.com/google/go-github/github"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/jungju/circle_manager/example/beegoapp/models"
	"github.com/jungju/circle_manager/example/beegoapp/requests"
	"github.com/jungju/circle_manager/example/beegoapp/synchronization"
	"github.com/jungju/circle_manager/modules"
	"github.com/sirupsen/logrus"
)

//  WebhookController operations for Webhook
type WebhookController struct {
	modules.BaseUserController
}

// Sync ...
// @Title Sync
// @Description
// @Success 204 {int}
// @Failure 403 body is empty
// @router /sync [post]
func (c *WebhookController) Sync() {
	if icss, err := models.GetAllIcs(nil); err == nil {
		for _, ics := range icss {
			if err := synchronization.SyncICS(ics.ICSURL); err != nil {
				logrus.WithError(err).Error("sync.SyncICS")
			}
		}
	} else {
		logrus.WithError(err).Error("GetAllIcs")
	}

	if trellos, err := models.GetAllTrello(nil); err == nil {
		for _, trello := range trellos {
			if err := synchronization.SyncTello(trello.UserName, trello.Key, trello.Token); err != nil {
				logrus.WithError(err).Error("sync.SyncTello")
			}
		}
	} else {
		logrus.WithError(err).Error("GetAllTrello")
	}

	c.Success(http.StatusNoContent, nil)
}

// PostGithub ...
// @Title PostGithub
// @Description
// //@Param	body		body 	requests.CreateHowForGithub	true		"body for How content"
// @Success 204 {int}
// @Failure 403 body is empty
// @router /github [post]
func (c *WebhookController) PostGithub() {
	event, err := github.ParseWebHook(github.WebHookType(c.Ctx.Request), c.Ctx.Input.RequestBody)
	if err != nil {
		c.ErrorAbort(400, err)
	}

	userName := ""
	repo := ""

	switch event := event.(type) {
	case *github.PushEvent:
		name := event.GetPusher().Name
		userName = *name
		if event.Repo != nil {
			repo = *event.Repo.Name
		}

		for _, commit := range event.Commits {
			message := commit.GetMessage()

			if err := modules.CreateItem(&models.GithubCommit{
				Description: "",
				Comments:    message,
				UserName:    userName,
				RepoName:    repo,
			}); err != nil {
				c.ErrorAbort(500, err)
			}
		}
	case *github.ReleaseEvent:
		name := event.GetSender().Login
		if name != nil {
			userName = *name
		}
		if event.Repo != nil {
			repo = *event.Repo.Name
		}
		tagName := ""
		release := event.GetRelease()
		if release != nil {
			if release.TagName != nil {
				tagName = *release.TagName
			}
		}
		if err := modules.CreateItem(&models.GithubRelease{
			Description: "",
			UserName:    userName,
			RepoName:    repo,
			TagName:     tagName,
			PreRelease:  release.GetPrerelease(),
			Message:     release.GetBody(),
		}); err != nil {
			c.ErrorAbort(500, err)
		}
	default:
		c.ErrorAbort(500, err)
		return
	}
	c.Success(http.StatusNoContent, nil)
}

// CalendarEvent ...
// @Title CalendarEvent
// @Description
// //@Param	body		body 	requests.CalendarEvent	true		"body for How content"
// @Success 204 {int}
// @Failure 403 body is empty
// @router /calendar/events/google [post]
func (c *WebhookController) CalendarEvent() {
	reqBody := &requests.CalendarEventGoogle{}
	c.SetRequestDataAndValid(reqBody)

	getItem := &models.Event{}
	if err := modules.GetItemWithFilter("google_id", reqBody.GoogleID, getItem); err != nil {
		if err != gorm.ErrRecordNotFound {
			c.ErrorAbort(500, err)
		}
		getItem = nil
	}

	updateItem := &models.Event{}
	copier.Copy(updateItem, reqBody)
	if getItem == nil {
		if err := modules.CreateItem(updateItem); err != nil {
			c.ErrorAbort(500, err)
		}
	} else {
		updateItem.ID = getItem.ID
		if err := modules.SaveItem(updateItem); err != nil {
			c.ErrorAbort(500, err)
		}
	}
	c.Success(http.StatusNoContent, nil)
}
