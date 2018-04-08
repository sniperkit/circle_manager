package circle_manager

type ControllerFile struct {
	MapControllerItems map[string][]ControllerItem
	RawFilePath        string
	CircleSet          *CircleSet
	TemplatePath       string
}

type ControllerItem struct {
	Name string
}

//"Controller/circle.go"

func NewControllerFile(circleSet *CircleSet) *ControllerFile {
	return &ControllerFile{
		MapControllerItems: map[string][]ControllerItem{
			"system": []ControllerItem{},
			"manual": []ControllerItem{},
			"auto":   []ControllerItem{},
		},
	}
}

func (fm *ControllerFile) Export(unit *CircleUnit) error {
	fm.MapControllerItems["auto"] = append(fm.MapControllerItems["auto"], ControllerItem{
		Name: unit.Name,
	})
	return nil
}

func (fm *ControllerFile) Import(status, l string) error {
	if targetLine(l, "addResourceAndMenu(") {
		fm.MapControllerItems[status] = append(fm.MapControllerItems[status], ControllerItem{
			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
		})
	}
	return nil
}
