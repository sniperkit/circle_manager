package circle_manager

// func (source *CircleTemplateSet) ExtractWithProperty() []CircleUnit {
// 	err := filepath.Walk(source.SourcePath, func(path string, info os.FileInfo, err error) error {
// 		fset := token.NewFileSet()
// 		d, err := parser.ParseDir(fset, "./", nil, parser.ParseComments)
// 		if err != nil {
// 			return err
// 		}

// 		for k, f := range d {
// 			fmt.Println("package", k)
// 			p := doc.New(f, "./", 0)

// 			for _, t := range p.Types {
// 				if strings.Index(t.Doc, "qs") >= 0 {
// 					// t.Name
// 					// t.
// 				}
// 				// fmt.Println("  type", t.Name)
// 				// fmt.Println("    docs:", t.Doc)
// 			}
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil
// 	}

// 	return nil
// }

// func (source *CircleTemplateSet) Extract() []CircleUnit {
// 	inFile, _ := os.Open(source.SourcePath)
// 	defer inFile.Close()
// 	scanner := bufio.NewScanner(inFile)
// 	scanner.Split(bufio.ScanLines)

// 	lineIndex := 0
// 	status := "meta"
// 	extractUnits := []CircleUnit{}
// 	for scanner.Scan() {
// 		l := scanner.Text()
// 		l = strings.TrimSpace(l)

// 		status = setStatus(status, l, "// circle:system:start", "circle:system")
// 		status = setStatus(status, l, "// circle:manual:start", "circle:manual")
// 		status = setStatus(status, l, "// circle:auto:start", "circle:auto")

// 		//if source.SourceType == "router" {
// 		curAppInfo.APIVersion = replaceAndTrim(curAppInfo.APIVersion, l, "// @APIVersion")
// 		curAppInfo.Title = replaceAndTrim(curAppInfo.Title, l, "// @Title")
// 		curAppInfo.Description = replaceAndTrim(curAppInfo.Description, l, "// @Description")
// 		curAppInfo.Contact = replaceAndTrim(curAppInfo.Contact, l, "// @Contact")
// 		curAppInfo.TermsOfServiceUrl = replaceAndTrim(curAppInfo.TermsOfServiceUrl, l, "// @TermsOfServiceUrl")
// 		curAppInfo.License = replaceAndTrim(curAppInfo.License, l, "// @License")
// 		curAppInfo.SecurityDefinition = replaceAndTrim(curAppInfo.SecurityDefinition, l, "// @SecurityDefinition")
// 		//}

// 		unit := GenUnit(status, l)
// 		extractUnits = append(extractUnits, *unit)

// 		lineIndex++
// 	}
// 	return extractUnits
// }

// func GenUnit(l, itemType string) *CircleUnit {
// 	if targetLine(l, "addResourceAndMenu(") {
// 		return &CircleUnit{
// 			Name:      getWord(l, "addResourceAndMenu(&models.", "{}, \""),
// 			MenuGroup: getWord(l, "\", \"", "\", anyoneAllow, -1)"),
// 			MenuName:  getWord(l, "{}, \"", "\", \""),
// 			IsSystem:  itemType == "system",
// 			IsManual:  itemType == "manual",
// 		}
// 	} else if targetLine(l, "addResourceAndMenu(") {
// 		return &CircleUnit{
// 			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
// 		}
// 	} else if targetLine(l, "addResourceAndMenu(") {
// 		return &CircleUnit{
// 			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
// 		}
// 	} else if targetLine(l, "addResourceAndMenu(") {
// 		return &CircleUnit{
// 			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
// 		}
// 	} else if targetLine(l, "addResourceAndMenu(") {
// 		return &CircleUnit{
// 			Name: getWord(l, "addResourceAndMenu(&models.", "{}, \""),
// 		}
// 	} else if targetLine(l, "addResourceAndMenu(") {
// 		return &CircleUnit{
// 			ControllerName: getWord(l, "beego.NSInclude(&controllers.", "Controller{}))"),
// 			URL:            getWord(l, "beego.NSNamespace(\"", "\", beego.NSInclude(&controllers."),
// 		}
// 	} else if targetLine(l, "addResourceAndMenu(") {
// 		return &CircleUnit{
// 			Name: getWord(l, "&", "{},"),
// 		}
// 	}
// 	return nil
// }

// func replaceAndTrim(target string, s string, old string) string {
// 	if target != "" || strings.Index(s, old) != 0 {
// 		return target
// 	}

// 	ret := strings.Replace(s, old, "", -1)
// 	return strings.TrimSpace(ret)
// }

// func getWord(s string, start string, end string) string {
// 	startIndex := strings.Index(s, start) + len(start)
// 	endIndex := strings.Index(s, end)

// 	return s[startIndex:endIndex]
// }

// func setStatus(rawSatatus, l, targetLine, newStatus string) string {
// 	if strings.Index(l, targetLine) == 0 {
// 		return strings.Replace(newStatus, "circle:", "", -1)
// 	}
// 	return rawSatatus
// }

// func targetLine(l, target string) bool {
// 	return strings.Index(l, target) == 0
// }
