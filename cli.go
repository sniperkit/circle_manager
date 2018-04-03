package circle_manager

import (
	"errors"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/jungju/gorm_manager"
)

var (
	envs *Envs
)

type Envs struct {
	DBHost   string
	DBPort   string
	DBName   string
	DBUser   string
	DBPass   string
	CircleID uint
}

func runGen() error {
	var err error
	if envs, err = initEnvs(); err != nil {
		return err
	}

	db, err := initDB()
	if err != nil {
		return err
	}

	cm, err := New(db)
	if err != nil {
		return err
	}

	if err := cm.GeneateSource(envs.CircleID); err != nil {
		return err
	}
	return nil
}

func initDB() (*gorm.DB, error) {
	var err error
	dbm, err := gorm_manager.New(&gorm_manager.DBConfig{
		DBType:             "mysql",
		DBHost:             envs.DBHost,
		DBPort:             envs.DBPort,
		DBName:             envs.DBName,
		DBUser:             envs.DBUser,
		DBPass:             envs.DBPass,
		AutoCreateDatabase: true,
		RecreateDatabase:   false,
		OnLog:              true,
	})
	if err != nil {
		return nil, err
	}
	return dbm.GetDB(), nil
}

func initEnvs() (*Envs, error) {
	circleIDstr := os.Getenv("CIRCLE_ID")
	circleIDUint, err := strconv.ParseUint(circleIDstr, 10, 64)
	if err != nil {
		return nil, err
	}

	envs := &Envs{
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:   os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASS"),
		CircleID: uint(circleIDUint),
	}

	if envs.DBHost == "" ||
		envs.DBPort == "" ||
		envs.DBName == "" ||
		envs.DBUser == "" ||
		envs.DBPass == "" {
		return nil, errors.New("need db info")
	}

	return envs, nil
}
