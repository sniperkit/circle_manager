package circle_manager

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/alecthomas/template"
	"github.com/jinzhu/gorm"
)

var (
	curAppInfo *AppInfo
	base       string
)

type CircleManager struct {
	ByType              string
	ModelsDir           string
	ControllersDir      string
	RequestsBodyDir     string
	ResponseBodyDir     string
	TablesTemplateSet   CircleTemplateSet
	RouterTemplateSet   CircleTemplateSet
	QORAdminTemplateSet CircleTemplateSet
}

type CircleTemplateSet struct {
	SourceType   string
	SourcePath   string
	TemplatePath string
}

type AppInfo struct {
	APIVersion         string
	Title              string
	Description        string
	Contact            string
	TermsOfServiceUrl  string
	License            string
	SecurityDefinition string
}

func (cm *CircleManager) prepare() {
	sd := func(raw, def string) string {
		if raw == "" {
			return filepath.Join("/", base, def)
		}
		return filepath.Join("/", base, raw)
	}

	cm.ModelsDir = sd(cm.ModelsDir, "/models")
	cm.ControllersDir = sd(cm.ControllersDir, "/controllers")
	cm.RequestsBodyDir = sd(cm.RequestsBodyDir, "/requests")
	cm.ResponseBodyDir = sd(cm.ResponseBodyDir, "/responses")
	cm.RouterTemplateSet.SourceType = "router"
	cm.RouterTemplateSet.SourcePath = sd(cm.RouterTemplateSet.SourcePath, "/routers/router.go")
	cm.RouterTemplateSet.TemplatePath = sd(cm.RouterTemplateSet.TemplatePath, "/templates/router.tmpl")
	cm.QORAdminTemplateSet.SourceType = "admin"
	cm.QORAdminTemplateSet.SourcePath = sd(cm.QORAdminTemplateSet.SourcePath, "/admin/circle.go")
	cm.QORAdminTemplateSet.TemplatePath = sd(cm.QORAdminTemplateSet.TemplatePath, "/templates/circle.tmpl")
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

	mapDBUnit := map[string]CircleUnit{}
	for _, dbUnit := range cs.Units {
		mapDBUnit[dbUnit.Name] = dbUnit
	}

	newCS := &CircleSet{
		Units: []CircleUnit{},
	}
	if cm.ByType == "db" {

	} else if cm.ByType == "source" {

	}

	mapUpdateUnit := map[string]CircleUnit{}
	for _, circleTemplateSet := range []CircleTemplateSet{
		cm.QORAdminTemplateSet,
		cm.RouterTemplateSet,
	} {
		for _, sourceUnit := range circleTemplateSet.Extract() {
			var newUnit CircleUnit
			if dbUnit, ok := mapDBUnit[sourceUnit.Name]; ok {
				if cm.ByType == "db" {
					newUnit = merge(dbUnit, sourceUnit)
				} else if cm.ByType == "source" {
					newUnit = merge(sourceUnit, dbUnit)
				}
			} else {
				newUnit = sourceUnit
			}
			mapUpdateUnit[newUnit.Name] = newUnit
		}

	}

	// for _, dbUnit := range cs.Units {
	// 	if sourceUnit, ok := mapSourceUnit[dbUnit.Name]; ok {
	// 		if cm.ByType == "db" {
	// 			mapSourceUnit[dbUnit.Name] = merge(dbUnit, sourceUnit)
	// 		} else if cm.ByType == "source" {
	// 			mapSourceUnit[dbUnit.Name] = merge(sourceUnit, dbUnit)
	// 		}
	// 	}
	// }

	// newCS := &CircleSet{
	// 	Units: []CircleUnit{},
	// }
	// if cm.ByType == "db" {

	// } else if cm.ByType == "source" {

	// }

	for _, unit := range mapSourceUnit {
		newCS.Units = append(newCS.Units, unit)
	}

	ExecuteTemplate()

	// gen(cm.ModelsDir, makePath("models.tmpl"), cs)
	// gen(cm.ControllersDir, makePath("controllers.tmpl"), cs)
	// gen(cm.RequestsBodyDir, makePath("requests.tmpl"), cs)
	// gen(cm.ResponseBodyDir, makePath("responses.tmpl"), cs)

	return nil
}

func merge(baseCU CircleUnit, changeCU CircleUnit) CircleUnit {
	setOnlyExistValue := func(baseString, newValue string) string {
		if newValue == "" {
			return baseString
		}
		return newValue
	}
	setOnlyExistValueForBool := func(baseString, newValue bool) bool {
		if newValue == false {
			return baseString
		}
		return newValue
	}
	changeCU.Name = setOnlyExistValue(changeCU.Name, baseCU.Name)
	changeCU.Description = setOnlyExistValue(changeCU.Description, baseCU.Description)
	changeCU.ControllerName = setOnlyExistValue(changeCU.ControllerName, baseCU.ControllerName)
	changeCU.VariableName = setOnlyExistValue(changeCU.VariableName, baseCU.VariableName)
	changeCU.Import = setOnlyExistValue(changeCU.Import, baseCU.Import)
	changeCU.URL = setOnlyExistValue(changeCU.URL, baseCU.URL)
	changeCU.MenuName = setOnlyExistValue(changeCU.MenuName, baseCU.MenuName)
	changeCU.MenuGroup = setOnlyExistValue(changeCU.MenuGroup, baseCU.MenuGroup)
	changeCU.IsEnable = setOnlyExistValueForBool(changeCU.IsEnable, baseCU.IsEnable)
	changeCU.IsManual = setOnlyExistValueForBool(changeCU.IsManual, baseCU.IsManual)
	changeCU.IsSystem = setOnlyExistValueForBool(changeCU.IsSystem, baseCU.IsSystem)
	return changeCU
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
