package circle_manager

type ResponseFile struct {
	MapResponseItems map[string][]ResponseItem
	RawFilePath      string
	CircleSet        *CircleSet
	TemplatePath     string
}

type ResponseItem struct {
	Name string
}

//"Response/circle.go"

func NewResponseFile(circleSet *CircleSet) *ResponseFile {
	return &ResponseFile{
		MapResponseItems: map[string][]ResponseItem{
			"system": []ResponseItem{},
			"manual": []ResponseItem{},
			"auto":   []ResponseItem{},
		},
	}
}

func (fm *ResponseFile) Export(unit *CircleUnit) error {
	fm.MapResponseItems["auto"] = append(fm.MapResponseItems["auto"], ResponseItem{
		Name: unit.Name,
	})
	return nil
}

func (fm *ResponseFile) Import(status, l string) error {
	if targetLine(l, "addResourceAndMenu(") {
		fm.MapResponseItems[status] = append(fm.MapResponseItems[status], ResponseItem{
			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
		})
	}
	return nil
}
