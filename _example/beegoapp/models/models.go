package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jungju/circle_manager/_example/beegoapp/envs"
	"github.com/jungju/circle_manager/_example/beegoapp/utils"
	"github.com/jungju/circle_manager/modules"
)

var (
	gGormDB *gorm.DB
)

var registedModels []interface{}

func registModel(model interface{}) {
	if registedModels == nil {
		registedModels = []interface{}{}
	}
	registedModels = append(registedModels, model)
}

func Initzation(db *gorm.DB) error {
	gGormDB = db

	if err := gGormDB.AutoMigrate(registedModels...).Error; err != nil {
		return err
	}

	if err := InitDefaultRow(); err != nil {
		return err
	}

	return nil
}

// InitDefaultRow ...
func InitDefaultRow() error {
	password, _ := utils.Digest(envs.SystemPassword)
	for _, obj := range []interface{}{
		&User{
			ID:                1,
			Email:             envs.SystemAdmin,
			EncryptedPassword: password,
		},
		&AuthIDentity{
			Provider:          "password",
			EncryptedPassword: password,
			UserID:            1,
			UID:               envs.SystemAdmin,
		},
	} {
		_er(modules.CreateItem(obj))
	}

	return nil
}

func _er(err error) {
	if err != nil {
		panic("테스트 파일 넣는중 에러발생 : " + err.Error())
	}
}
