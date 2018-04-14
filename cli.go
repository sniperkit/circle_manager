package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jungju/gorm_manager"
	"github.com/urfave/cli"
)

var (
	envs *Envs
)

type Envs struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassWord string
	CircleID   uint
	RootPath   string
}

func (envs *Envs) Valid() error {
	if envs.DBHost == "" {
		return errors.New("Require DBHost")
	}
	if envs.DBPort <= 0 {
		return errors.New("Require DBPort")
	}
	if envs.DBName == "" {
		return errors.New("Require DBName")
	}
	if envs.DBUser == "" {
		return errors.New("Require DBUser")
	}
	if envs.DBPassWord == "" {
		return errors.New("Require DBPassWord")
	}
	if envs.CircleID <= 0 {
		return errors.New("Require CircleID")
	}
	if envs.RootPath == "" {
		envs.RootPath = "./"
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Circle-Manager"
	app.Usage = "for NO-CODE Platform"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "dbHost", Value: "localhost", Usage: "DB Host"},
		cli.StringFlag{Name: "dbPort", Value: "3306", Usage: "DB Port"},
		cli.StringFlag{Name: "dbName", Value: "circle", Usage: "DB Name"},
		cli.StringFlag{Name: "dbUser", Value: "root", Usage: "DB User"},
		cli.StringFlag{Name: "dbPassword", Value: "password", Usage: "DB Password"},
		cli.StringFlag{Name: "circleID", Value: "1", Usage: "CircleID"},
		cli.StringFlag{Name: "rootPath", Value: "./", Usage: "RootPath"},
	}
	app.Action = func(c *cli.Context) error {
		envs = &Envs{
			DBHost:     c.String("dbHost"),
			DBPort:     c.Int("dbPort"),
			DBName:     c.String("dbName"),
			DBUser:     c.String("dbUser"),
			DBPassWord: c.String("dbPassword"),
			CircleID:   c.Uint("circleID"),
			RootPath:   c.String("rootPath"),
		}

		err := envs.Valid()
		if err != nil {
			return err
		}

		return runGen()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runGen() error {
	var err error
	db, err := initDB()
	if err != nil {
		return err
	}

	cm := &CircleManager{}

	if err := db.AutoMigrate(
		&CircleSet{},
		&CircleUnit{},
		&CircleUnitProperty{},
	).Error; err != nil {
		return err
	}

	if err := cm.GeneateSource(db, envs.CircleID); err != nil {
		return err
	}
	return nil
}

func initDB() (*gorm.DB, error) {
	var err error
	dbm, err := gorm_manager.New(&gorm_manager.DBConfig{
		DBType:             "mysql",
		DBHost:             envs.DBHost,
		DBPort:             fmt.Sprintf("%d", envs.DBPort),
		DBName:             envs.DBName,
		DBUser:             envs.DBUser,
		DBPass:             envs.DBPassWord,
		AutoCreateDatabase: true,
		RecreateDatabase:   false,
		OnLog:              true,
	})
	if err != nil {
		return nil, err
	}
	return dbm.GetDB(), nil
}
