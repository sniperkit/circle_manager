package circle_manager

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/alecthomas/template"
)

func gen(exportDir string, templatePath string, cs *CircleSet) error {
	logrus.Debug(exportDir)
	for _, unit := range cs.Units {
		filename := getExportFilename(unit.Name)
		sourceFilepath := filepath.Join(exportDir, filename)

		f, err := os.OpenFile(sourceFilepath, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		t := template.Must(template.ParseFiles(templatePath))
		if err := t.Execute(f, unit); err != nil {
			panic(err)
		}

		exec.Command("gofmt", "-w", sourceFilepath).Output()
	}

	return nil
}

func getExportFilename(name string) string {
	return fmt.Sprintf("%s.go", makeFirstLowerCase(name))
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
