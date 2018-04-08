package circle_manager

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

// func (target *Target) GenGo() {
// 	os.MkdirAll(filepath.Join(target.BuildPath), 0777)
// 	f, err := os.OpenFile(filepath.Join(target.BuildPath, target.GoFilename), os.O_CREATE|os.O_WRONLY, 0777)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	executeTemplate(f, target.Object, target.TemplateFile)
// 	exec.Command("gofmt", "-w", filepath.Join(target.BuildPath, target.GoFilename)).Output()
// }

// func executeTemplate(f *os.File, templateObject interface{}, templateFile string) {
// 	t := template.Must(template.ParseFiles(templateFile))
// 	if err := t.Execute(f, templateObject); err != nil {
// 		panic(err)
// 	}
// }
