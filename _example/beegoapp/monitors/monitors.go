package monitors

import (
	"time"

	"github.com/jungju/circle_manager/_example/beegoapp/models"
	"github.com/jungju/circle_manager/_example/beegoapp/synchronization"
	"github.com/jungju/circle_manager/modules"
	"github.com/sirupsen/logrus"
)

func RunSendNotification() {
	t := time.NewTicker(60 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if err := modules.SendActiveNotifications(); err != nil {
				logrus.WithError(err).Error("RunSendNotification")
			}
		}
	}
}

func RunSync() {
	t := time.NewTicker(1 * time.Hour)
	defer t.Stop()

	for {
		select {
		case <-t.C:
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
		}
	}
}
