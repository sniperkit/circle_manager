package circle_manager

type TableFile struct {
	MapTableItems map[string][]TableItem
	RawFilePath   string
	CircleSet     *CircleSet
	TemplatePath  string
}

type TableItem struct {
	Name string
}

func TableFileParseFile(status, l string, targetFile interface{}) {
	rf := targetFile.(*TableFile)

	if targetLine(l, "addResourceAndMenu(") {
		rf.MapTableItems["auto"] = append(rf.MapTableItems["auto"], TableItem{
			Name: getWord(l, "&", "{},"),
		})
	}
}

func (fm *TableFile) Export(unit *CircleUnit) error {
	fm.MapTableItems["auto"] = append(fm.MapTableItems["auto"], TableItem{
		Name: unit.Name,
	})
	return nil
}

func (fm *TableFile) Import(status, l string) error {
	if targetLine(l, "addResourceAndMenu(") {
		fm.MapTableItems[status] = append(fm.MapTableItems[status], TableItem{
			Name: getWord(l, "&", "{},"),
		})
	}
	return nil
}
