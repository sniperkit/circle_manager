package circle_manager

type RequestFile struct {
	MapRequestItems map[string][]RequestItem
	RawFilePath     string
	CircleSet       *CircleSet
	TemplatePath    string
}

type RequestItem struct {
	Name string
}

//"Request/circle.go"

func NewRequestFile(circleSet *CircleSet) *RequestFile {
	return &RequestFile{
		MapRequestItems: map[string][]RequestItem{
			"system": []RequestItem{},
			"manual": []RequestItem{},
			"auto":   []RequestItem{},
		},
	}
}

func (fm *RequestFile) Export(unit *CircleUnit) error {
	fm.MapRequestItems["auto"] = append(fm.MapRequestItems["auto"], RequestItem{
		Name: unit.Name,
	})
	return nil
}

func (fm *RequestFile) Import(status, l string) error {
	if targetLine(l, "addResourceAndMenu(") {
		fm.MapRequestItems[status] = append(fm.MapRequestItems[status], RequestItem{
			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
		})
	}
	return nil
}
