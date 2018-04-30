package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jungju/circle_manager/modules"
	"github.com/jungju/gorm_manager"
	"github.com/urfave/cli"
)

var (
	envs *Envs
)

type Envs struct {
	Mode            string
	Name            string
	DBHost          string
	DBPort          int
	DBName          string
	DBUser          string
	DBPassWord      string
	CircleID        uint
	RootPath        string
	OnlyControllers bool
	OnlyModels      bool
	OnlyRequests    bool
	OnlyResponses   bool
}

func envsValid() error {
	if envs.Mode == "generate" || envs.Mode == "envs" {
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
	} else if envs.Mode == "add" || envs.Mode == "delete" {
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
		cli.StringFlag{Name: "mode", Value: "", Usage: "generate, add, delete, scan, safe, envs"},
		cli.StringFlag{Name: "name", Value: "", Usage: "target name, ex)Car, Group, User"},
		cli.StringFlag{Name: "dbHost", Value: "localhost", Usage: "DB Host", EnvVar: "DB_HOST"},
		cli.IntFlag{Name: "dbPort", Value: 3306, Usage: "DB Port", EnvVar: "DB_PORT"},
		cli.StringFlag{Name: "dbName", Value: "circle", Usage: "DB Name", EnvVar: "DB_NAME"},
		cli.StringFlag{Name: "dbUser", Value: "root", Usage: "DB User", EnvVar: "DB_USER"},
		cli.StringFlag{Name: "dbPassword", Value: "password", Usage: "DB Password", EnvVar: "DB_PASSWORD"},
		cli.UintFlag{Name: "circleID", Value: 1, Usage: "CircleID", EnvVar: "CIRCLE_ID"},
		cli.StringFlag{Name: "rootPath", Value: "./", Usage: "RootPath", EnvVar: "ROOT_PATH"},
		cli.BoolFlag{Name: "onlyControllers", Usage: "RootPaths"},
		cli.BoolFlag{Name: "onlyModels", Usage: "onlyModels"},
		cli.BoolFlag{Name: "onlyRequests", Usage: "onlyRequests"},
		cli.BoolFlag{Name: "onlyResponses", Usage: "onlyResponses"},
	}
	app.Action = func(c *cli.Context) error {
		envs = &Envs{
			Mode:            c.String("mode"),
			Name:            c.String("name"),
			DBHost:          c.String("dbHost"),
			DBPort:          c.Int("dbPort"),
			DBName:          c.String("dbName"),
			DBUser:          c.String("dbUser"),
			DBPassWord:      c.String("dbPassword"),
			CircleID:        c.Uint("circleID"),
			RootPath:        c.String("rootPath"),
			OnlyControllers: c.Bool("onlyControllers"),
			OnlyModels:      c.Bool("onlyModels"),
			OnlyRequests:    c.Bool("onlyRequests"),
			OnlyResponses:   c.Bool("onlyResponses"),
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
		} else if envs.Mode == "delete" {
			return runDelete()
		} else if envs.Mode == "safe" {
			return runSafemode()
		} else if envs.Mode == "envs" {
			return runSetEnv()
		} else if envs.Mode == "import" {
			return runImport()
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runImport() error {
	cm := &CircleManager{}
	cm.prepare()

	cm.ImportCircle()

	return nil
}

func runSafemode() error {
	cm := &CircleManager{}
	if err := cm.GenerateSource(&modules.CircleSet{}); err != nil {
		return err
	}
	return nil
}

func runDelete() error {
	cm := &CircleManager{}
	cm.prepare()

	return cm.DeleteManual()
}

func runAdd() error {
	cm := &CircleManager{}
	cm.prepare()

	return cm.AppendManual()
}

func runGen() error {
	if err := initDB(); err != nil {
		return err
	}

	cm := &CircleManager{}
	cm.prepare()

	cs, err := modules.GetCircleSetByID(envs.CircleID)
	if err != nil {
		return err
	}

	return cm.GenerateSource(cs)
}

func runSetEnv() error {
	if err := initDB(); err != nil {
		return err
	}

	cs, err := modules.GetCircleSetByID(envs.CircleID)
	if err != nil {
		return err
	}

	return setRunAppEnv(cs.RunAppEnvs)
}

func initDB() error {
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
		return err
	}

	return modules.Initzation(dbm.GetDB(), "", "", "")
}

func setRunAppEnv(runAppEnvs string) error {
	for _, envStr := range strings.Split(runAppEnvs, " ") {
		envKeyValue := strings.Split(envStr, "=")
		if len(envKeyValue) == 2 {
			if envKeyValue[0] != "" && envKeyValue[1] != "" {
				os.Setenv(envKeyValue[0], envKeyValue[1])
			}
		}
	}
	return nil
}
