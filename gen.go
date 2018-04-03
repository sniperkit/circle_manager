package circle_manager

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	funk "github.com/thoas/go-funk"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	GoExt       = "go"
	TemplateExt = "tmpl"
)

type ObjectType struct {
	Target     string
	ExportPath string
	ExportName string
}

var (
	ignoreUnitNames = []string{"User", "Notification", "NotificationType", "AuthIdentity", "CircleSet", "CircleUnit", "CircleUnitProperty"}
	objectTypes     = []ObjectType{
		ObjectType{"controllers", "controllers", ""},
		ObjectType{"models", "models", ""},
		ObjectType{"requests", "requests", ""},
		ObjectType{"responses", "responses", ""},
	}
	staticObjectTypes = []ObjectType{
		ObjectType{"admin", "admin", "circle"},
		ObjectType{"router", "routers", "router"},
		ObjectType{"tables", "models", "tables"},
	}
)

func (o ObjectType) GetTemplateFilename() string {
	return fmt.Sprintf("%s.%s", makeFirstLowerCase(o.Target), TemplateExt)
}

func (o ObjectType) GetExportFilename() string {
	return fmt.Sprintf("%s.%s", makeFirstLowerCase(o.ExportName), GoExt)
}

func getCircleSetByID(db *gorm.DB, id uint) (circleSet *CircleSet, err error) {
	circleSet = &CircleSet{
		ID: id,
	}

	err = db.Preload("Units").Preload("Units.Properties").First(circleSet, "id = ?", id).Error
	return circleSet, err
}

func (cm *CircleManager) runGen(cs *CircleSet) error {
	targetUnits := []CircleUnit{}
	for _, unit := range cs.Units {
		if funk.Contains(ignoreUnitNames, unit.Name) {
			continue
		}
		targetUnits = append(targetUnits, unit)
	}

	ts := &TargetSet{
		Targets:       []*Target{},
		BuildRootPath: cm.BuildPath,
	}

	os.RemoveAll(ts.BuildRootPath)

	for _, target := range staticObjectTypes {
		cs.Units = targetUnits

		ts.Targets = append(ts.Targets, &Target{
			TemplateFile: filepath.Join(cm.TemplatePath, target.GetTemplateFilename()),
			Path:         target.ExportPath,
			BuildPath:    filepath.Join(ts.BuildRootPath, target.ExportPath),
			GoFilename:   target.GetExportFilename(),
			Object:       cs,
		})
	}

	for _, unit := range targetUnits {
		for _, target := range objectTypes {
			target.ExportName = unit.Name
			cp := unit

			ts.Targets = append(ts.Targets, &Target{
				TemplateFile: filepath.Join(cm.TemplatePath, target.GetTemplateFilename()),
				Path:         target.ExportPath,
				BuildPath:    filepath.Join(ts.BuildRootPath, target.ExportPath),
				GoFilename:   target.GetExportFilename(),
				Object:       &cp,
			})
		}
	}

	for _, target := range ts.Targets {
		target.GenGo()

		srcFile, err := os.Open(filepath.Join(target.BuildPath, target.GoFilename))
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(filepath.Join("..", target.Path, target.GoFilename))
		if err != nil {
			return err
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, srcFile); err != nil {
			return err
		}
	}

	cmd := exec.Command("./gen_queryset.sh")
	cmd.Dir = cm.ModelsPath
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func makeFirstLowerCase(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
