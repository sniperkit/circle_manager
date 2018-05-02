package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/alecthomas/template"
	"github.com/jungju/circle_manager/modules"
)

const (
	adminTemplate = `addResourceAndMenu(&models.{{.Name}}{}, "{{.MenuName}}", "{{.MenuGroup}}", anyoneAllow, -1)
`
	routerTemplate = `beego.NSNamespace("/{{.GetURL}}",
	beego.NSInclude(
		&controllers.{{.Name}}Controller{},
	),
),
`
)

type changeTemplateFunc func(string, string) (string, error)

//TODO: 삭제하기
var circleSet *modules.CircleSet

type CircleManager struct {
	ByType          string
	MapTemplateSets map[string]*CircleTemplateSet
}

type CircleTemplateSet struct {
	SourceType   string
	SourcePath   string
	TemplatePath string
	IsMulti      bool
}

func (c CircleTemplateSet) IsExistSourceFile() bool {
	return existsFile(c.SourcePath)
}

func (c CircleTemplateSet) IsExistTemplateFile() bool {
	return existsFile(c.TemplatePath)
}

func (cm *CircleManager) prepare() {
	setTemplateSet := func(sourceType, sourcePath, templatePath string, isMulti bool) {
		if cm.MapTemplateSets == nil {
			cm.MapTemplateSets = map[string]*CircleTemplateSet{}
		}

		cm.MapTemplateSets[sourceType] = &CircleTemplateSet{
			SourceType:   sourceType,
			SourcePath:   cm.GetSourcePath(sourceType, filepath.Join(envs.RootPath, sourcePath)),
			TemplatePath: cm.GetTemplatePath(sourceType, templatePath),
			IsMulti:      isMulti,
		}
	}

	setTemplateSet("router", "routers/router.go", "templates/router.tmpl", false)
	setTemplateSet("admin", "admin/circle.go", "templates/admin.tmpl", false)
	setTemplateSet("models", "models", "templates/models.tmpl", true)
	setTemplateSet("controllers", "controllers", "templates/controllers.tmpl", true)
	setTemplateSet("requests", "requests", "templates/requests.tmpl", true)
	setTemplateSet("responses", "responses", "templates/responses.tmpl", true)
}

func (cm *CircleManager) GetSourcePath(sourceType string, def string) string {
	if raw, ok := cm.MapTemplateSets[sourceType]; ok {
		if raw.SourcePath != "" {
			return raw.SourcePath
		}
	}
	return def
}

func (cm *CircleManager) GetTemplatePath(sourceType string, def string) string {
	if raw, ok := cm.MapTemplateSets[sourceType]; ok {
		if raw.TemplatePath != "" {
			return raw.TemplatePath
		}
	}
	return def
}

func (cm *CircleManager) GenerateSource(cs *modules.CircleSet) error {
	circleSet = cs
	cm.prepare()

	for i, _ := range circleSet.Units {
		circleSet.Units[i].Import = circleSet.Import
	}

	for _, circleTemplateSet := range cm.MapTemplateSets {
		check := func(onlyItemStr string, olnyItem bool) bool {
			return circleTemplateSet.SourceType == onlyItemStr && olnyItem
		}

		if check("models", envs.OnlyModels) ||
			check("controllers", envs.OnlyControllers) ||
			check("requests", envs.OnlyRequests) ||
			check("responses", envs.OnlyResponses) {
			continue
		}

		if !circleTemplateSet.IsExistSourceFile() {
			fmt.Printf("Not found %s\n", filepath.Join(circleTemplateSet.SourcePath))
			continue
		}

		if !circleTemplateSet.IsExistTemplateFile() {
			fmt.Printf("Not found %s\n", filepath.Join(circleTemplateSet.TemplatePath))
			continue
		}

		if circleTemplateSet.IsMulti {
			for _, unit := range cs.Units {
				unitSourceFile := filepath.Join(circleTemplateSet.SourcePath, fmt.Sprintf("%s.go", unit.GetVariableName()))
				fmt.Printf("Start ExecuteTemplate %s\n", unitSourceFile)

				if err := executeTemplate(
					unitSourceFile,
					circleTemplateSet.TemplatePath,
					unit,
				); err != nil {
					fmt.Printf("Error : %s\n", err.Error())
					return err
				}

				if circleTemplateSet.SourceType == "models" {
					executeQueryset(circleTemplateSet.SourcePath, unit.GetVariableName())
				}
			}
		} else {
			fmt.Printf("%s", circleTemplateSet.SourcePath)

			if err := executeTemplate(circleTemplateSet.SourcePath, circleTemplateSet.TemplatePath, cs); err != nil {
				fmt.Printf("Error : %s\n", err.Error())
				return err
			}
		}
	}

	return nil
}

func executeQueryset(dir string, varName string) {
	fmt.Printf("goqueryset 실행 %s\n", fmt.Sprintf("%s.go", varName))
	cmd := exec.Command("goqueryset", "-in", fmt.Sprintf("%s.go", varName))
	cmd.Dir = dir
	if out, err := cmd.Output(); err != nil {
		fmt.Printf("Error : %s. %s\n", err.Error(), out)
	} else {
		fmt.Printf("goqueryset : %s\n", out)
	}
}

func (cm *CircleManager) AppendManual() error {
	manualUnit := &modules.CircleUnit{
		Name:      envs.Name,
		MenuName:  envs.Name,
		MenuGroup: "etc.",
		IsManual:  true,
		IsEnable:  true,
	}

	if envs.OnlyModels || envs.OnlyControllers || envs.OnlyRequests || envs.OnlyResponses {
		return nil
	}

	routerTemplateSet := cm.MapTemplateSets["router"]
	if err := appendManual(routerTemplateSet.TemplatePath, routerTemplate, manualUnit); err != nil {
		return err
	}

	adminTemplateSet := cm.MapTemplateSets["admin"]
	if err := appendManual(adminTemplateSet.TemplatePath, adminTemplate, manualUnit); err != nil {
		return err
	}

	return cm.GenerateSource(&modules.CircleSet{
		Units: []*modules.CircleUnit{manualUnit},
	})
}

func (cm *CircleManager) DeleteManual() error {
	manualUnit := &modules.CircleUnit{
		Name: envs.Name,
	}

	routerTemplateSet := cm.MapTemplateSets["router"]
	if err := deleteManual(routerTemplateSet.TemplatePath, routerTemplate, manualUnit); err != nil {
		return err
	}

	adminTemplateSet := cm.MapTemplateSets["admin"]
	if err := deleteManual(adminTemplateSet.TemplatePath, adminTemplate, manualUnit); err != nil {
		return err
	}

	removeFunc := func(dirName string) {
		os.Remove(filepath.Join(envs.RootPath, dirName, fmt.Sprintf("%s.go", manualUnit.GetVariableName())))
		fmt.Printf("Deleted %s\n", filepath.Join(envs.RootPath, dirName, fmt.Sprintf("%s.go", manualUnit.GetVariableName())))
	}

	removeFunc("controllers")
	removeFunc("models")
	removeFunc("requests")
	removeFunc("responses")

	fmt.Printf("Deleted %s\n", filepath.Join(envs.RootPath, "models", fmt.Sprintf("autogenerated_%s.go", manualUnit.GetVariableName())))
	os.Remove(filepath.Join(envs.RootPath, "models", fmt.Sprintf("autogenerated_%s.go", manualUnit.GetVariableName())))

	return cm.GenerateSource(&modules.CircleSet{})
}

func deleteManual(templatefile string, appendText string, unit *modules.CircleUnit) error {
	return saveTemplate(templatefile, appendText, unit, func(read string, tpl string) (string, error) {
		return strings.Replace(string(read), tpl, "", -1), nil
	})
}

func appendManual(templatefile string, appendText string, unit *modules.CircleUnit) error {
	return saveTemplate(templatefile, appendText, unit, func(read string, tpl string) (string, error) {
		append := fmt.Sprintf("%s// circle:manual:end", tpl)

		if strings.Index(string(read), tpl) >= 0 {
			return "", errors.New("이미 추가된 수동 소스 입니다.")
		}

		return strings.Replace(string(read), "// circle:manual:end", append, -1), nil
	})
}

func saveTemplate(templatefile string, appendText string, unit *modules.CircleUnit, ctFunc changeTemplateFunc) error {
	t := template.Must(template.ParseFiles(appendText))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, unit); err != nil {
		return err
	}

	read, err := ioutil.ReadFile(templatefile)
	if err != nil {
		return err
	}

	output, err := ctFunc(string(read), tpl.String())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(templatefile, []byte(output), 0)
}

func executeTemplate(dest string, templatePath string, templateObject interface{}) error {
	os.Remove(dest)

	f, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	t := template.Must(template.ParseFiles(templatePath))
	if err := t.Execute(f, templateObject); err != nil {
		return err
	}

	fmt.Printf("Excuted source : %s\n", dest)

	if _, err = exec.Command("gofmt", "-w", dest).Output(); err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}

	return nil
}

func existsFile(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}
