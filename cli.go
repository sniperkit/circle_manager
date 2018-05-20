package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jungju/circle_manager/modules"
	"github.com/jungju/gorm_manager"
	"github.com/urfave/cli"
)

var (
	envs *Envs
)

type Envs struct {
	Mode           string
	Name           string
	DBHost         string
	DBPort         int
	DBName         string
	DBUser         string
	DBPassWord     string
	CircleID       uint
	RootPath       string
	DockerImageURL string
}

func envsValid() error {
	if envs.Mode == "generate" || envs.Mode == "envs" || envs.Mode == "import" {
		if envs.CircleID <= 0 && envs.Mode != "import" {
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
	} else if envs.Mode == "build" {
		if envs.DockerImageURL == "" {
			return errors.New("Require DockerImageURL")
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
		cli.StringFlag{Name: "mode", Value: "", Usage: "generate, add, delete, scan, safe, envs, "},
		cli.StringFlag{Name: "name", Value: "", Usage: "target name, ex)Car, Group, User"},
		cli.StringFlag{Name: "dbHost", Value: "localhost", Usage: "DB Host", EnvVar: "DB_HOST"},
		cli.IntFlag{Name: "dbPort", Value: 3306, Usage: "DB Port", EnvVar: "DB_PORT"},
		cli.StringFlag{Name: "dbName", Value: "circle", Usage: "DB Name", EnvVar: "DB_NAME"},
		cli.StringFlag{Name: "dbUser", Value: "root", Usage: "DB User", EnvVar: "DB_USER"},
		cli.StringFlag{Name: "dbPassword", Value: "password", Usage: "DB Password", EnvVar: "DB_PASSWORD"},
		cli.UintFlag{Name: "circleID", Value: 0, Usage: "CircleID", EnvVar: "CIRCLE_ID"},
		cli.StringFlag{Name: "rootPath", Value: "./", Usage: "RootPath", EnvVar: "ROOT_PATH"},
		cli.BoolFlag{Name: "onlyControllers", Usage: "RootPaths"},
		cli.BoolFlag{Name: "onlyModels", Usage: "onlyModels"},
		cli.BoolFlag{Name: "onlyRequests", Usage: "onlyRequests"},
		cli.BoolFlag{Name: "onlyResponses", Usage: "onlyResponses"},
		cli.StringFlag{Name: "dockerImageURL", Value: "", Usage: "jungju/circle", EnvVar: "DOCKER_IMAGE_URL"},
	}
	app.Action = func(c *cli.Context) error {
		logrus.Info("Set flag for prepare action")

		envs = &Envs{
			Mode:           c.String("mode"),
			Name:           c.String("name"),
			DBHost:         c.String("dbHost"),
			DBPort:         c.Int("dbPort"),
			DBName:         c.String("dbName"),
			DBUser:         c.String("dbUser"),
			DBPassWord:     c.String("dbPassword"),
			CircleID:       c.Uint("circleID"),
			RootPath:       c.String("rootPath"),
			DockerImageURL: c.String("dockerImageURL"),
		}

		logrus.Info("Start circle cmd.")

		if err := envsValid(); err != nil {
			logrus.WithError(err).Error()
		}

		logrus.WithField("mode", envs.Mode).Info("Start action for request mode.")
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
		} else if envs.Mode == "build" {
			return runBuild()
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("End app.")
}

func runBuild() error {
	logrus.Info("Start build")

	if err := genBeegoAppResource(); err != nil {
		return err
	}
	if err := dockerBuild(envs.DockerImageURL); err != nil {
		return err
	}
	return dockerPush(envs.DockerImageURL)
}

func runImport() error {
	logrus.Info("Start import")

	if err := initDB(); err != nil {
		return err
	}

	cm := &CircleManager{}

	cs, err := cm.ImportCircle()
	if err != nil {
		logrus.WithError(err).Error()
		return err
	}
	cs.ID = envs.CircleID

	return cm.SaveManualCircleSetToDB(cs)
}

func runSafemode() error {
	logrus.Info("Start safe mode")

	cm := &CircleManager{}
	if err := cm.GenerateSource(&modules.CircleSet{}); err != nil {
		return err
	}
	return nil
}

func runDelete() error {
	logrus.WithField("name", envs.Name).Info("Start delete sources and code")
	routerPath := filepath.Join(envs.RootPath, ROUTER_PATH)

	read, err := ioutil.ReadFile(routerPath)
	if err != nil {
		return err
	}

	source := removeRouterSource(string(read), envs.Name)
	if err := ioutil.WriteFile(routerPath, []byte(source), 0); err != nil {
		return err
	}

	executeGofmtW(routerPath)

	logrus.WithField("name", envs.Name).Infof("Deleting all sources")
	for _, sourceTypes := range []string{"models", "controllers", "requests", "responses"} {
		removeFile := filepath.Join(envs.RootPath, sourceTypes, fmt.Sprintf("%s.go", modules.MakeFirstLowerCase(envs.Name)))
		logrus.WithField("removeFile", removeFile).Info("Delet source file")
		if err := os.Remove(removeFile); err != nil {
			logrus.WithError(err).Error()
		}
	}
	logrus.WithField("name", envs.Name).Infof("Deleted all sources")
	return nil
}

func runAdd() error {
	logrus.WithField("name", envs.Name).Info("Start add sources and code")

	cm := &CircleManager{}
	return cm.GenerateSource(&modules.CircleSet{
		Units: []*modules.CircleUnit{
			&modules.CircleUnit{
				Name: envs.Name,
				EnableControllerSource: true,
				EnableModelSource:      true,
				EnableRequestSource:    true,
				EnableResponseSource:   true,
				IsManual:               true,
			},
		},
	})
}

func runGen() error {
	if err := initDB(); err != nil {
		return err
	}

	cm := &CircleManager{}

	cs, err := modules.GetCircleSetByIDForGen(envs.CircleID)
	if err != nil {
		return err
	}

	if err := cm.GenerateSource(cs); err != nil {
		return err
	}

	return nil
}

func runSetEnv() error {
	if err := initDB(); err != nil {
		return err
	}

	cs, err := modules.GetCircleSetByIDForGen(envs.CircleID)
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
