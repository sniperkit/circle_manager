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

const ROUTER_PATH = "_example/routers/router.go"

type CircleManager struct{}

func (cm *CircleManager) GenerateSource(cs *modules.CircleSet) error {
	read, err := ioutil.ReadFile(ROUTER_PATH)
	if err != nil {
		return err
	}

	output, err := generateRouter(string(read), cs)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(ROUTER_PATH, []byte(output), 0); err != nil {
		return err
	}

	for _, sourceTypes := range []string{"models", "controllers", "requests", "responses"} {
		if err := generateItems(sourceTypes, cs); err != nil {
			return err
		}
	}

	return nil
}

func generateRouter(rawSource string, cs *modules.CircleSet) (string, error) {
	routerCodes := ""
	logrus.Infof("Generate Total %d", len(cs.Units))
	for _, unit := range cs.Units {
		logrus.Infof("Start generate %s", unit.Name)
		if !unit.EnableControllerSource {
			fmt.Println("Skip controller : ", unit.Name)
			continue
		}
		if routerCode, err := saveRouterSource(routerTemplate, unit); err == nil {
			routerCodes = routerCodes + routerCode
		} else {
			return "", err
		}
	}

	cleanedSource := cleanRouterSource(rawSource)
	append := fmt.Sprintf("// circle:manual:end\n\t\t%s\n\t\t// circle:manual:end", routerCodes)
	return strings.Replace(cleanedSource, "// circle:manual:start\n\t\t// circle:manual:end", append, -1), nil
}

func cleanRouterSource(sources string) string {
	start := strings.Index(sources, CIRCLE_AUTO_START_WORD)
	end := strings.Index(sources, CIRCLE_AUTO_END_WORD)
	return sources[0:start+len(CIRCLE_AUTO_START_WORD)+1] + "\t\t" + sources[end:len(sources)]
}

func generateItems(sourceType string, cs *modules.CircleSet) error {
	for _, unit := range cs.Units {
		if unit.IsManual && unit.IsSystem {
			continue
		}
		unitSourceFile := filepath.Join(envs.RootPath, sourceType, fmt.Sprintf("%s.go", unit.GetVariableName()))
		fmt.Printf("Start ExecuteTemplate %s\n", unitSourceFile)

		newCU := &modules.CircleUnit{}
		copier.Copy(newCU, unit)
		newCU.Properties = []*modules.CircleUnitProperty{}
		for _, property := range unit.Properties {
			if (sourceType == "requests" || sourceType == "responses") &&
				strings.Index(property.Type, "models.") == 0 {
				fmt.Println("Skip Property : ", property.Name, property.Type)
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
			fmt.Printf("Error : %s\n", err.Error())
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

func executeQueryset(varName string) {
	fmt.Printf("goqueryset 실행 %s\n", fmt.Sprintf("%s.go", varName))
	cmd := exec.Command("goqueryset", "-in", fmt.Sprintf("%s.go", varName))
	cmd.Dir = envs.RootPath
	if out, err := cmd.Output(); err != nil {
		fmt.Printf("Error : %s. %s\n", err.Error(), out)
	} else {
		fmt.Printf("goqueryset : %s\n", out)
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

	fmt.Printf("Excuted source : %s\n", dest)

	if _, err = exec.Command("gofmt", "-w", dest).Output(); err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}

	return nil
}
