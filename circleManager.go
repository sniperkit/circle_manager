package circle_manager

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/structs"

	"github.com/alecthomas/template"
	"github.com/jinzhu/gorm"
)

type CircleManager struct {
	TemplateDir     string
	ModelsDir       string
	ControllersDir  string
	TablesPath      string
	RoutersPath     string
	QORAdminPath    string
	RequestsBodyDir string
	ResponseBodyDir string
}

func (cm *CircleManager) prepare() {
	sd := func(raw, def string) string {
		if raw == "" {
			return def
		}
		return raw
	}

	cm.TemplateDir = sd(cm.TemplateDir, "templates")
	cm.ModelsDir = sd(cm.ModelsDir, "models")
	cm.TablesPath = sd(cm.TablesPath, "models/tables.go")
	cm.RoutersPath = sd(cm.RoutersPath, "routers/router.go")
	cm.ControllersDir = sd(cm.ControllersDir, "contorllers")
	cm.QORAdminPath = sd(cm.QORAdminPath, "admin/circle.go")
	cm.RequestsBodyDir = sd(cm.RequestsBodyDir, "requests")
	cm.ResponseBodyDir = sd(cm.ResponseBodyDir, "responses")
}

func (cm *CircleManager) GeneateSource(db *gorm.DB, circleIDUint uint) error {
	cs, err := getCircleSetByID(db, circleIDUint)
	if err != nil {
		return err
	}

	return cm.GeneateSourceBySet(cs)
}

func (cm *CircleManager) GeneateSourceBySet(cs *CircleSet) error {
	cm.prepare()

	makePath := func(fm string) string {
		return filepath.Join(cm.TemplateDir, fm)
	}

	//TODO: Import 하기

	//TODO: DB 싱크 하기

	//TODO: Export 하기

	ImportAndExport(NewAdminFile(cs))
	Import("models/circle.go", adminSource)
	Sync(fm)
	Export(makePath("tables.tmpl"), "models/circle.go", cs.Units, adminSource)

	TableSourceExport(cm.TablesPath, makePath("tables.tmpl"), cs)
	RouterSourceExport(cm.RoutersPath, makePath("tables.tmpl"), cs)
	AdminSourceExport(cm.QORAdminPath, makePath("tables.tmpl"), cs)

	gen(cm.ModelsDir, makePath("models.tmpl"), cs)
	gen(cm.ControllersDir, makePath("controllers.tmpl"), cs)
	gen(cm.RequestsBodyDir, makePath("requests.tmpl"), cs)
	gen(cm.ResponseBodyDir, makePath("responses.tmpl"), cs)

	return nil
}

func ImportAndExport(cs CircleSource, units []CircleUnit) {
	Import("", cs)
	mapCs := structs.Map(cs)
	Sync(cs)
	Export("", "", units, cs)
}

func getCircleSetByID(db *gorm.DB, id uint) (circleSet *CircleSet, err error) {
	circleSet = &CircleSet{
		ID: id,
	}

	err = db.Preload("Units").Preload("Units.Properties").First(circleSet, "id = ?", id).Error
	return circleSet, err
}

type CircleSource interface {
	Export(unit *CircleUnit) error
	Import(status, l string) error
}

func Import(rawFilePath string, targetParseFile CircleSource) error {
	inFile, _ := os.Open(rawFilePath)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	lineIndex := 0
	status := "meta"
	for scanner.Scan() {
		l := scanner.Text()
		l = strings.TrimSpace(l)

		status = setStatus(status, l, "// circle:system:start", "circle:system")
		status = setStatus(status, l, "// circle:manual:start", "circle:manual")
		status = setStatus(status, l, "// circle:auto:start", "circle:auto")

		targetParseFile.Import(status, l)
		lineIndex++
	}
	return nil
}

func Export(templatePath string, sourcePath string, units []CircleUnit, cs CircleSource) error {
	for _, unit := range units {
		cs.Export(&unit)
	}

	if err := ExecuteTemplate(sourcePath, templatePath, cs); err != nil {
		return err
	}

	return nil
}

func Sync(units []CircleUnit, cis []map[string]interface{}) error {
	mapUnit := map[string]CircleUnit{}
	for _, unit := range units {
		mapUnit[unit.Name] = unit
	}

	for _, item := range cis {
		if err := SaveCircleUnit(item); err != nil {
			return err
		}
	}

	return nil
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
