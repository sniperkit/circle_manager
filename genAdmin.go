package circle_manager

type AdminFile struct {
	MapAdminItems map[string][]AdminItem
	RawFilePath   string
	CircleSet     *CircleSet
	TemplatePath  string
}

//"admin/circle.go"

func NewAdminFile(circleSet *CircleSet) *AdminFile {
	return &AdminFile{
		MapAdminItems: map[string][]AdminItem{
			"system": []AdminItem{},
			"manual": []AdminItem{},
			"auto":   []AdminItem{},
		},
	}
}

func (fm *AdminFile) Export(unit *CircleUnit) error {
	fm.MapAdminItems["auto"] = append(fm.MapAdminItems["auto"], AdminItem{
		Name:      unit.Name,
		MenuName:  unit.MenuName,
		MenuGroup: unit.MenuGroup,
	})
	return nil
}

func (fm *AdminFile) Import(status, l string) error {
	if targetLine(l, "addResourceAndMenu(") {
		fm.MapAdminItems[status] = append(fm.MapAdminItems[status], AdminItem{
			Name:      getWord(l, "addResourceAndMenu(&models.", "{}, \""),
			MenuGroup: getWord(l, "\", \"", "\", anyoneAllow, -1)"),
			MenuName:  getWord(l, "{}, \"", "\", \""),
		})
	}
	return nil
}
