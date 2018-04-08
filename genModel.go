package circle_manager

type ModelFile struct {
	MapModelItems map[string][]ModelItem
	RawFilePath   string
	CircleSet     *CircleSet
	TemplatePath  string
}

type ModelItem struct {
	Name string
}

//"Model/circle.go"

func NewModelFile(circleSet *CircleSet) *ModelFile {
	return &ModelFile{
		MapModelItems: map[string][]ModelItem{
			"system": []ModelItem{},
			"manual": []ModelItem{},
			"auto":   []ModelItem{},
		},
	}
}

func (fm *ModelFile) Export(unit *CircleUnit) error {
	fm.MapModelItems["auto"] = append(fm.MapModelItems["auto"], ModelItem{
		Name: unit.Name,
	})
	return nil
}

func (fm *ModelFile) Import(status, l string) error {
	if targetLine(l, "addResourceAndMenu(") {
		fm.MapModelItems[status] = append(fm.MapModelItems[status], ModelItem{
			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
		})
	}
	return nil
}
