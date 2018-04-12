package circle_manager

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/alecthomas/template"
	"github.com/jinzhu/gorm"
)

var (
	basePath string
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
			SourcePath:   sd(cm.GetSourcePath(sourceType), filepath.Join(basePath, sourcePath)),
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

var circleSet *CircleSet

func (cm *CircleManager) GeneateSource(db *gorm.DB, circleIDUint uint) error {
	cs, err := getCircleSetByID(db, circleIDUint)
	if err != nil {
		return err
	}

	return cm.GeneateSourceBySet(cs)
}

func (cm *CircleManager) GeneateSourceBySet(cs *CircleSet) error {
	circleSet = cs
	cm.prepare()

	for _, circleTemplateSet := range cm.MapTemplateSets {
		if circleTemplateSet.IsMulti {
			for _, unit := range cs.Units {
				if err := ExecuteTemplate(
					filepath.Join(circleTemplateSet.SourcePath, fmt.Sprintf("%s.go", unit.GetVariableName())),
					circleTemplateSet.TemplatePath,
					unit,
				); err != nil {
					return err
				}
			}
		} else {
			if err := ExecuteTemplate(circleTemplateSet.SourcePath, circleTemplateSet.TemplatePath, cs); err != nil {
				return err
			}
		}
	}

	return nil
}

func getCircleSetByID(db *gorm.DB, id uint) (circleSet *CircleSet, err error) {
	circleSet = &CircleSet{
		ID: id,
	}

	err = db.Preload("Units").Preload("Units.Properties").First(circleSet, "id = ?", id).Error
	return circleSet, err
}

func ExecuteTemplate(dest string, templatePath string, templateObject interface{}) error {
	f, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	t := template.Must(template.ParseFiles(templatePath))
	if err := t.Execute(f, templateObject); err != nil {
		return err
	}

	_, err = exec.Command("gofmt", "-w", dest).Output()
	return err
}
