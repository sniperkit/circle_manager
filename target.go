package circle_manager

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/alecthomas/template"
)

type TargetSet struct {
	Targets       []*Target
	BuildRootPath string
}

type Target struct {
	TemplateFile string
	Path         string
	BuildPath    string
	GoFilename   string
	Object       interface{}
}

func (ts *TargetSet) AppendTarget(target ObjectType, object interface{}) {
	ts.Targets = append(ts.Targets, &Target{
		TemplateFile: target.GetTemplateFilename(),
		Path:         target.ExportPath,
		BuildPath:    filepath.Join(ts.BuildRootPath, target.ExportPath),
		GoFilename:   target.GetExportFilename(),
		Object:       object,
	})
}

func (target *Target) GenGo() {
	os.MkdirAll(filepath.Join(target.BuildPath), 0777)
	f, err := os.OpenFile(filepath.Join(target.BuildPath, target.GoFilename), os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	executeTemplate(f, target.Object, target.TemplateFile)
	exec.Command("gofmt", "-w", filepath.Join(target.BuildPath, target.GoFilename)).Output()
}

func executeTemplate(f *os.File, templateObject interface{}, templateFile string) {
	t := template.Must(template.ParseFiles(templateFile))
	if err := t.Execute(f, templateObject); err != nil {
		panic(err)
	}
}
