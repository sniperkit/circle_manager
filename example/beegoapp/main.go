package main

import (
	"os"

	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jungju/circle_manager/example/beegoapp/admin"
	"github.com/jungju/circle_manager/example/beegoapp/envs"
	"github.com/jungju/circle_manager/example/beegoapp/models"
	"github.com/jungju/circle_manager/example/beegoapp/monitors"
	_ "github.com/jungju/circle_manager/example/beegoapp/routers"
	"github.com/jungju/circle_manager/modules"
	"github.com/jungju/gorm_manager"
	"github.com/jungju/qor_admin_auth"
	"github.com/sirupsen/logrus"
)

var (
	dbm *gorm_manager.DBManager
)

func main() {
	if os.Getenv("CIRCLE_MODE") == "GEN" {
		logrus.Info("Gen 모드입니다. 종료합니다.")
		os.Exit(0)
	}
	envs.InitEnvs()

	initDB()

	initServer()

	initRouter()

	logrus.Info("Start App...")
	beego.Run()
}

func initRouter() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
}

func initDB() {
	var err error

	dbm, err = gorm_manager.New(&gorm_manager.DBConfig{
		DBType:             "mysql",
		DBHost:             envs.DBHost,
		DBPort:             envs.DBPort,
		DBName:             envs.DBName,
		DBUser:             envs.DBUser,
		DBPassword:         envs.DBPassword,
		AutoCreateDatabase: true,
		RecreateDatabase:   false,
		OnLog:              true,
	})
	if err != nil {
		panic(err)
	}

	if err := modules.Initzation(dbm.GetDB(), envs.SystemToken, envs.SecretKey, envs.UserTokenHeaderName); err != nil {
		panic(err)
	}

	if err := models.Initzation(dbm.GetDB()); err != nil {
		panic(err)
	}
}

func initServer() {
	mQorAdmin, err := qor_admin_auth.New(&qor_admin_auth.QorAdminConfig{
		AdminURL:               "admin",
		DB:                     dbm.GetDB(),
		OnAuth:                 true,
		AuthLoginURL:           "/auth/login",
		AuthLogoutURL:          "/auth/logout",
		AuthURL:                "auth",
		AdminUserModel:         &models.User{},
		AuthEnableGoogle:       false,
		AuthEnablePassword:     true,
		AuthGoogleClientID:     envs.GoogleClientID,
		AuthGoogleClientSecret: envs.GoogleClientSecret,
	})
	if err != nil {
		panic(err)
	}
	if err := admin.SetAdmin(mQorAdmin, dbm.GetDB()); err != nil {
		panic(err)
	}

	//hostname := utils.GetHostname()
	//if hostname == "circle-1" {
	go monitors.RunSendNotification()
	go monitors.RunSync()
	//}
}
