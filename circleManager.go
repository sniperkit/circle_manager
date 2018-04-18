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
	"github.com/jinzhu/gorm"
	"github.com/jungju/circle_manager/modules"
	"github.com/sirupsen/logrus"
)

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

func (cm *CircleManager) GetSourcePath(sourceType string) string {
	if _, ok := cm.MapTemplateSets[sourceType]; !ok {
		return ""
	}
	return cm.MapTemplateSets[sourceType].SourcePath
}

func (cm *CircleManager) GetTemplatePath(sourceType string) string {
	if _, ok := cm.MapTemplateSets[sourceType]; !ok {
		return ""
	}
	return cm.MapTemplateSets[sourceType].TemplatePath
}

func (cm *CircleManager) prepare() {
	sd := func(raw, def string) string {
		if raw == "" {
			return def
		}
		return raw
	}

	setTemplateSet := func(sourceType, sourcePath, templatePath string, isMulti bool) {
		if cm.MapTemplateSets == nil {
			cm.MapTemplateSets = map[string]*CircleTemplateSet{}
		}
		cm.MapTemplateSets[sourceType] = &CircleTemplateSet{
			SourceType:   sourceType,
			SourcePath:   sd(cm.GetSourcePath(sourceType), filepath.Join(envs.RootPath, sourcePath)),
			TemplatePath: sd(cm.GetTemplatePath(sourceType), templatePath),
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

var circleSet *modules.CircleSet

func (cm *CircleManager) GeneateSource(db *gorm.DB, circleIDUint uint) error {
	cs, err := getCircleSetByID(db, circleIDUint)
	if err != nil {
		return err
	}

	return cm.GeneateSourceBySet(cs)
}

func (cm *CircleManager) GeneateSourceBySet(cs *modules.CircleSet) error {
	circleSet = cs
	cm.prepare()

	for i, _ := range circleSet.Units {
		circleSet.Units[i].Import = circleSet.Import
	}

	for _, circleTemplateSet := range cm.MapTemplateSets {
		if circleTemplateSet.SourceType != "models" && envs.OnlyModels {
			continue
		}
		if circleTemplateSet.SourceType != "controllers" && envs.OnlyControllers {
			continue
		}
		if circleTemplateSet.SourceType != "requests" && envs.OnlyRequests {
			continue
		}
		if circleTemplateSet.SourceType != "responses" && envs.OnlyResponses {
			continue
		}

		if _, err := os.Stat(filepath.Join(circleTemplateSet.SourcePath)); os.IsNotExist(err) {
			return errors.New("Not found" + filepath.Join(circleTemplateSet.SourcePath))
		}
		if _, err := os.Stat(filepath.Join(circleTemplateSet.TemplatePath)); os.IsNotExist(err) {
			return errors.New("Not found" + filepath.Join(circleTemplateSet.SourcePath))
		}

		if circleTemplateSet.IsMulti {
			for _, unit := range cs.Units {
				unitSourceFile := filepath.Join(circleTemplateSet.SourcePath, fmt.Sprintf("%s.go", unit.GetVariableName()))
				if err := ExecuteTemplate(
					unitSourceFile,
					circleTemplateSet.TemplatePath,
					unit,
				); err != nil {
					return err
				}
				if circleTemplateSet.SourceType == "models" {
					cmd := exec.Command("goqueryset", "-in", fmt.Sprintf("%s.go", unit.GetVariableName()))
					cmd.Dir = circleTemplateSet.SourcePath
					if out, err := cmd.Output(); err != nil {
						logrus.WithError(err).
							WithField("source", unitSourceFile).
							Error("goqueryset 에러:", string(out))
					}
				}
			}
		} else {
			if err := ExecuteTemplate(circleTemplateSet.SourcePath, circleTemplateSet.TemplatePath, cs); err != nil {
				logrus.WithField("SourcePath", circleTemplateSet.SourcePath).
					WithField("TemplatePath", circleTemplateSet.TemplatePath).
					Error()
				return err
			}
		}
	}

	return nil
}

func (cm *CircleManager) AppendManual(unit *modules.CircleUnit) error {
	cm.prepare()

	if envs.OnlyModels || envs.OnlyControllers || envs.OnlyRequests || envs.OnlyResponses {
		return nil
	}

	routerTemplateSet := cm.MapTemplateSets["router"]
	routerTemplate := `beego.NSNamespace("/{{.URL}}",
			beego.NSInclude(
				&controllers.{{.Name}}Controller{},
			),
		),
		`
	if err := appendManual(routerTemplateSet.TemplatePath, routerTemplate, unit); err != nil {
		return err
	}

	adminTemplateSet := cm.MapTemplateSets["admin"]
	adminTemplate := `addResourceAndMenu(&models.{{.Name}}{}, "{{.MenuName}}", "{{.MenuGroup}}", anyoneAllow, -1)
	`
	if err := appendManual(adminTemplateSet.TemplatePath, adminTemplate, unit); err != nil {
		return err
	}

	return nil
}

func (cm *CircleManager) DeleteManual(unit *modules.CircleUnit) error {
	cm.prepare()

	routerTemplateSet := cm.MapTemplateSets["router"]
	routerTemplate := `beego.NSNamespace("/{{.URL}}",
			beego.NSInclude(
				&controllers.{{.Name}}Controller{},
			),
		),
		`
	if err := deleteManual(routerTemplateSet.TemplatePath, routerTemplate, unit); err != nil {
		return err
	}

	adminTemplateSet := cm.MapTemplateSets["admin"]
	adminTemplate := `addResourceAndMenu(&models.{{.Name}}{}, "{{.MenuName}}", "{{.MenuGroup}}", anyoneAllow, -1)
	`
	if err := deleteManual(adminTemplateSet.TemplatePath, adminTemplate, unit); err != nil {
		return err
	}

	os.Remove(filepath.Join(envs.RootPath, "controllers", fmt.Sprintf("%s.go", unit.GetVariableName())))
	os.Remove(filepath.Join(envs.RootPath, "models", fmt.Sprintf("%s.go", unit.GetVariableName())))
	os.Remove(filepath.Join(envs.RootPath, "models", fmt.Sprintf("autogenerated_%s.go", unit.GetVariableName())))
	os.Remove(filepath.Join(envs.RootPath, "requests", fmt.Sprintf("%s.go", unit.GetVariableName())))
	os.Remove(filepath.Join(envs.RootPath, "responses", fmt.Sprintf("%s.go", unit.GetVariableName())))

	return nil
}

func deleteManual(templatefile string, appendText string, unit *modules.CircleUnit) error {
	t := template.Must(template.New("t1").Parse(appendText))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, unit); err != nil {
		return err
	}

	read, err := ioutil.ReadFile(templatefile)
	if err != nil {
		return err
	}

	newContents := strings.Replace(string(read), tpl.String(), "", -1)

	err = ioutil.WriteFile(templatefile, []byte(newContents), 0)
	if err != nil {
		return err
	}
	return nil
}

func appendManual(templatefile string, appendText string, unit *modules.CircleUnit) error {
	t := template.Must(template.New("t1").Parse(appendText))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, unit); err != nil {
		return err
	}

	read, err := ioutil.ReadFile(templatefile)
	if err != nil {
		return err
	}

	append := fmt.Sprintf("%s// circle:manual:end", tpl.String())

	if strings.Index(string(read), tpl.String()) >= 0 {
		return errors.New("이미 추가된 수동 소스 입니다.")
	}

	newContents := strings.Replace(string(read), "// circle:manual:end", append, -1)

	err = ioutil.WriteFile(templatefile, []byte(newContents), 0)
	if err != nil {
		return err
	}
	return nil
}

func setStatus(rawWhere, l, targetLine, newWhere string) string {
	if strings.Index(l, targetLine) == 0 {
		return strings.Replace(newWhere, "circle:", "", -1)
	}
	return rawWhere
}

func ExecuteTemplate(dest string, templatePath string, templateObject interface{}) error {
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

	logrus.Info("Created source :" + dest)
	if _, err = exec.Command("gofmt", "-w", dest).Output(); err != nil {
		logrus.WithError(err).
			WithField("dest", dest).
			WithField("TemplatePath", templatePath).
			Error("formating중 에러")
	}

	return nil
}

func getCircleSetByID(db *gorm.DB, id uint) (circleSet *modules.CircleSet, err error) {
	circleSet = &modules.CircleSet{
		ID: id,
	}

	err = db.Preload("Units").Preload("Units.Properties").First(circleSet, "id = ?", id).Error
	return circleSet, err
}
