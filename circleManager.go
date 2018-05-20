package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/alecthomas/template"
	"github.com/jinzhu/copier"
	"github.com/jungju/circle_manager/modules"
)

const (
	CIRCLE_AUTO_START_WORD = "// circle:auto:start"
	CIRCLE_AUTO_END_WORD   = "// circle:auto:end"
)

const ROUTER_PATH = "routers/router.go"

type CircleManager struct{}

func (cm *CircleManager) GenerateSource(cs *modules.CircleSet) error {
	routerPath := filepath.Join(envs.AppDir, ROUTER_PATH)
	read, err := ioutil.ReadFile(routerPath)
	if err != nil {
		return err
	}

	output, err := generateRouter(string(read), cs)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(routerPath, []byte(output), 0); err != nil {
		return err
	}

	executeGofmtW(ROUTER_PATH)

	for _, sourceTypes := range []string{"models", "controllers", "requests", "responses"} {
		if err := generateItems(sourceTypes, cs); err != nil {
			return err
		}
	}

	return nil
}

func generateRouter(rawSource string, cs *modules.CircleSet) (string, error) {
	routerCodes := ""
	logrus.Info("Generate Total : ", len(cs.Units))
	for _, unit := range cs.Units {
		logrus.Info("Start generate : ", unit.Name)
		if !unit.EnableControllerSource {
			logrus.Warn("Skip controller : ", unit.Name)
			continue
		}
		if routerCode, err := saveRouterSource(routerTemplate, unit); err == nil {
			routerCodes = routerCodes + routerCode
		} else {
			return "", err
		}
	}

	targetPlace := "manual"
	if envs.Mode == "gen" {
		targetPlace = "auto"
	}
	targetStartComment := fmt.Sprintf("// circle:%s:start", targetPlace)
	targetEndComment := fmt.Sprintf("// circle:%s:end", targetPlace)

	cleanedSource := cleanRouterSource(rawSource)
	appendCode := fmt.Sprintf("%s\n\t\t%s\n\t\t%s", targetStartComment, routerCodes, targetEndComment)
	logrus.WithField("code", appendCode).Info("Adding code")
	return strings.Replace(cleanedSource, fmt.Sprintf("%s\n\t\t%s", targetStartComment, targetEndComment), appendCode, -1), nil
}

func cleanRouterSource(sources string) string {
	start := strings.Index(sources, CIRCLE_AUTO_START_WORD)
	end := strings.Index(sources, CIRCLE_AUTO_END_WORD)
	if start <= 0 || end <= 0 {
		logrus.Error("Failed clean code.")
	}
	return sources[0:start+len(CIRCLE_AUTO_START_WORD)+1] + "\t\t" + sources[end:len(sources)]
}

func removeRouterSource(sources string, targetName string) string {
	routerCode, err := saveRouterSource(routerTemplate, &modules.CircleUnit{
		Name: targetName,
	})
	if err != nil {
		logrus.WithError(err).Warn()
		return sources
	}
	return strings.Replace(sources, routerCode, "", -1)
}

func generateItems(sourceType string, cs *modules.CircleSet) error {
	for _, unit := range cs.Units {
		if unit.IsManual && unit.IsSystem {
			continue
		}
		unitSourceFile := filepath.Join(envs.AppDir, sourceType, fmt.Sprintf("%s.go", unit.GetVariableName()))
		logrus.Info("Start ExecuteTemplate %s\n", unitSourceFile)

		newCU := &modules.CircleUnit{}
		copier.Copy(newCU, unit)
		newCU.Properties = []*modules.CircleUnitProperty{}
		for _, property := range unit.Properties {
			if (sourceType == "requests" || sourceType == "responses") &&
				strings.Index(property.Type, "models.") == 0 {
				logrus.Info("Skip Property : ", property.Name, property.Type)
				continue
			}
			newCU.Properties = append(newCU.Properties, property)
		}

		teamplateSource := ""
		switch sourceType {
		case "models":
			teamplateSource = MODEL_TEMPLATE
		case "controllers":
			teamplateSource = CONTROLLER_TEMPLATE
		case "requests":
			teamplateSource = REQUEST_TEMPLATE
		case "responses":
			teamplateSource = RESPONSE_TEMPLATE
		}

		if err := executeTemplate(
			unitSourceFile,
			teamplateSource,
			newCU,
		); err != nil {
			logrus.WithError(err).Error()
			return err
		}

		if sourceType == "models" {
			executeQueryset(unit.GetVariableName())
		}
	}

	return nil
}

func saveRouterSource(appendText string, unit *modules.CircleUnit) (string, error) {
	t, err := template.New("").Parse(appendText)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, unit); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func executeGofmtW(filepath string) {
	logrus.WithField("filepath", filepath).Info("executeGofmtW")
	cmd := exec.Command("gofmt", "-w", filepath)
	cmd.Dir = envs.AppDir
	if out, err := cmd.Output(); err != nil {
		logrus.WithError(err).WithField("output", out).Error()
	} else {
		logrus.WithField("output", out).Info("executeGofmtW")
	}
}

func executeQueryset(varName string) {
	logrus.WithField("varName", varName).Info("goqueryset")
	cmd := exec.Command("goqueryset", "-in", fmt.Sprintf("%s.go", varName))
	cmd.Dir = envs.AppDir
	if out, err := cmd.Output(); err != nil {
		logrus.WithError(err).WithField("output", out).Error()
	} else {
		logrus.WithField("output", out).Info("goqueryset")
	}
}

func executeTemplate(dest string, templateSource string, templateObject interface{}) error {
	os.Remove(dest)

	f, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	t, _ := template.New("").Parse(templateSource)
	if err := t.Execute(f, templateObject); err != nil {
		return err
	}

	logrus.WithField("dest", dest).Info("Excuted source")

	if _, err = exec.Command("gofmt", "-w", dest).Output(); err != nil {
		logrus.WithError(err).Error()
	}

	return nil
}
