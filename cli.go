package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/inflection"
	"github.com/jungju/circle_manager/modules"
	"github.com/jungju/gorm_manager"
	"github.com/urfave/cli"
)

var (
	envs *Envs
)

type Envs struct {
	Mode       string
	Name       string
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassWord string
	CircleID   uint
	RootPath   string
}

func envsValid() error {
	if envs.Mode == "generate" {
		if envs.CircleID <= 0 {
			return errors.New("circleID가 없으므로 종료")
		}

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
		if envs.RootPath == "" {
			envs.RootPath = "./"
		}
	} else if envs.Mode == "add" {
		if envs.Name == "" {
			return errors.New("Require Name")
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Circle-Manager"
	app.Usage = "for NO-CODE Platform"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "mode", Value: "", Usage: "generate, add, scan"},
		cli.StringFlag{Name: "name", Value: "", Usage: "add mame"},
		cli.StringFlag{Name: "dbHost", Value: "localhost", Usage: "DB Host", EnvVar: "DB_HOST"},
		cli.IntFlag{Name: "dbPort", Value: 3306, Usage: "DB Port", EnvVar: "DB_PORT"},
		cli.StringFlag{Name: "dbName", Value: "circle", Usage: "DB Name", EnvVar: "DB_NAME"},
		cli.StringFlag{Name: "dbUser", Value: "root", Usage: "DB User", EnvVar: "DB_USER"},
		cli.StringFlag{Name: "dbPassword", Value: "password", Usage: "DB Password", EnvVar: "DB_PASSWORD"},
		cli.UintFlag{Name: "circleID", Value: 1, Usage: "CircleID", EnvVar: "CIRCLE_ID"},
		cli.StringFlag{Name: "rootPath", Value: "./", Usage: "RootPath", EnvVar: "ROOT_PATH"},
	}
	app.Action = func(c *cli.Context) error {
		envs = &Envs{
			Mode:       c.String("mode"),
			Name:       c.String("name"),
			DBHost:     c.String("dbHost"),
			DBPort:     c.Int("dbPort"),
			DBName:     c.String("dbName"),
			DBUser:     c.String("dbUser"),
			DBPassWord: c.String("dbPassword"),
			CircleID:   c.Uint("circleID"),
			RootPath:   c.String("rootPath"),
		}

		err := envsValid()
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if envs.Mode == "generate" {
			return runGen()
		} else if envs.Mode == "add" {
			return runAdd()
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runAdd() error {
	cm := &CircleManager{}

	manualUnit := modules.CircleUnit{
		Name:      envs.Name,
		URL:       inflection.Plural(envs.Name),
		MenuName:  envs.Name,
		MenuGroup: "etc.",
		IsManual:  true,
		IsEnable:  true,
	}

	if err := cm.AppendManual(&manualUnit); err != nil {
		return err
	}

	if err := cm.GeneateSourceBySet(&modules.CircleSet{
		Units: []modules.CircleUnit{manualUnit},
	}); err != nil {
		return err
	}

	return nil
}

func runGen() error {
	var err error
	db, err := initDB()
	if err != nil {
		return err
	}

	cm := &CircleManager{}

	if err := db.AutoMigrate(
		&modules.CircleSet{},
		&modules.CircleUnit{},
		&modules.CircleUnitProperty{},
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
		DBPassword:         envs.DBPassWord,
		AutoCreateDatabase: true,
		RecreateDatabase:   false,
		OnLog:              true,
	})
	if err != nil {
		return nil, err
	}
	return dbm.GetDB(), nil
}
