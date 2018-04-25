package modules

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	ErrDataNotFound        = errors.New("ErrDataNotFound")
	ErrDulicatedObject     = errors.New("ErrDulicatedObject")
	ErrInvalidQuery        = errors.New("ErrInvalidQuery")
	ErrUnauthorized        = errors.New("ErrUnauthorized")
	ErrInvalidRequestBody  = errors.New("ErrInvalidRequestBody")
	ErrInvalidRequestParam = errors.New("ErrInvalidRequestParam")
	ErrInvalidUsername     = errors.New("ErrInvalidUsername")
	ErrInvalidToken        = errors.New("ErrInvalidToken")
	ErrUnknown             = errors.New("Unknown")
)

func Initzation(db *gorm.DB, systemToken string, secretKeys string, userTokenHeaderName string) error {
	_SystemToken = systemToken
	_SecretKeys = secretKeys
	_UserTokenHeaderName = userTokenHeaderName

	gGormDB = db
	if err := gGormDB.AutoMigrate(
		&Notification{},
		&NotificationType{},

		&CircleSet{},
		&CircleUnit{},
		&CircleUnitProperty{},
	).Error; err != nil {
		return err
	}
	return nil
}
