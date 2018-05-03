package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/copier"

	"github.com/davecgh/go-spew/spew"

	"github.com/fatih/structtag"

	"github.com/jinzhu/gorm"
	"github.com/jungju/circle_manager/modules"
)

type FlagRead struct {
	RouterReadedAppTitle              bool
	RouterReadedAppVersion            bool
	RouterReadedAppDescription        bool
	RouterReadedAppContact            bool
	RouterReadedAppTermsOfServiceUrl  bool
	RouterReadedAppLicense            bool
	RouterReadedAppSecurityDefinition bool
}

func (cm *CircleManager) ImportCircle() (*modules.CircleSet, error) {
	adminCircleSet, err := cm.ImportCircleAdmin()
	if err != nil {
		return nil, err
	}

	routerCircleSet, err := cm.ImportCircleRouter()
	if err != nil {
		return nil, err
	}

	if err := mergeFromAdmin(routerCircleSet, adminCircleSet); err != nil {
		return nil, err
	}

	controllersCircleSet, err := scanSource("controllers", cm.MapTemplateSets["controllers"].SourcePath)
	if err != nil {
		return nil, err
	}

	modelsCircleSet, err := scanSource("models", cm.MapTemplateSets["models"].SourcePath)
	if err != nil {
		return nil, err
	}

	requestsCircleSet, err := scanSource("requests", cm.MapTemplateSets["requests"].SourcePath)
	if err != nil {
		return nil, err
	}

	responsesCircleSet, err := scanSource("responses", cm.MapTemplateSets["responses"].SourcePath)
	if err != nil {
		return nil, err
	}

	if err := mergeFromModelsAndRequestsAndResponses(routerCircleSet, controllersCircleSet, modelsCircleSet, requestsCircleSet, responsesCircleSet); err != nil {
		return nil, err
	}

	return routerCircleSet, nil
}

func (cm *CircleManager) SaveManualCircleSetToDB(manualCS *modules.CircleSet) error {
	fmt.Println("SaveManualCircleSetToDB")
	var dbCircleSet *modules.CircleSet
	createDB := false
	if manualCS.ID > 0 {
		var err error
		dbCircleSet, err = modules.GetCircleSetByID(manualCS.ID)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return err
			}
			createDB = true
			dbCircleSet = nil
		}
	}

	if dbCircleSet == nil {
		dbCircleSet = &modules.CircleSet{}
		if err := copier.Copy(dbCircleSet, manualCS); err != nil {
			return err
		}
	} else {
		mapDBCircleSet := map[string]*modules.CircleUnit{}
		for _, unit := range dbCircleSet.Units {
			mapDBCircleSet[unit.Name] = unit
		}

		for _, unit := range manualCS.Units {
			if dbUnit, ok := mapDBCircleSet[unit.Name]; ok {
				dbUnit.IsManual = !unit.IsAutogen
			} else {
				dbCircleSet.Units = append(dbCircleSet.Units, unit)
			}
		}
	}

	//TODO: Save로 처리 할것
	if createDB {
		fmt.Println("CS 추가")
		_, err := modules.AddCircleSet(dbCircleSet)
		return err
	}
	fmt.Println("CS 수정")
	return modules.UpdateCircleSetByID(dbCircleSet)
}

func mergeFromAdmin(routerCircleSet *modules.CircleSet, adminCircleSet *modules.CircleSet) error {
	checkUnits := map[string][]bool{}

	mapRouterCircleSet := map[string]*modules.CircleUnit{}
	for _, unit := range routerCircleSet.Units {
		checkUnits[unit.Name] = []bool{true, false}
		mapRouterCircleSet[unit.Name] = unit
	}

	mapAdminCircleSet := map[string]*modules.CircleUnit{}
	for _, unit := range adminCircleSet.Units {
		if _, ok := checkUnits[unit.Name]; !ok {
			checkUnits[unit.Name] = []bool{false, true}
		} else {
			checkUnits[unit.Name][1] = true
		}
		mapAdminCircleSet[unit.Name] = unit
	}

	for checkUnitName, checks := range checkUnits {
		if checks[0] && !checks[1] {
			//라우터는 있고 Admin은 없다
		} else if !checks[0] && checks[1] {
			//라우터는 없고 Admin은 있다.
			routerCircleSet.Units = append(routerCircleSet.Units, mapAdminCircleSet[checkUnitName])
		} else if checks[0] && checks[1] {
			mapRouterCircleSet[checkUnitName].EnableAdminSource = true
		}
	}
	return nil
}

func mergeFromModelsAndRequestsAndResponses(
	routerCircleSet *modules.CircleSet,
	controllersCircleSet *modules.CircleSet,
	modelsCircleSet *modules.CircleSet,
	requestsCircleSet *modules.CircleSet,
	responsesCircleSet *modules.CircleSet,
) error {
	mapRouterCircleSet := map[string]*modules.CircleUnit{}
	for _, unit := range routerCircleSet.Units {
		mapRouterCircleSet[unit.Name] = unit
	}

	for _, unit := range controllersCircleSet.Units {
		if cu, ok := mapRouterCircleSet[unit.Name]; ok {
			cu.EnableControllerSource = true
		} else {
			// TODO:
			cu := &modules.CircleUnit{
				Name:         unit.Name,
				IsCreateble:  unit.IsCreateble,
				IsUpdateble:  unit.IsUpdateble,
				IsGetAllable: unit.IsGetAllable,
				IsGetOneable: unit.IsGetOneable,
				IsDeleteble:  unit.IsDeleteble,
			}
			mapRouterCircleSet[unit.Name] = cu
			routerCircleSet.Units = append(routerCircleSet.Units, cu)
		}
	}

	mapProperies := map[string]map[string]*modules.CircleUnitProperty{}
	for _, unit := range modelsCircleSet.Units {
		if cu, ok := mapRouterCircleSet[unit.Name]; ok {
			cu.EnableModelSource = true
			cu.Properties = append(mapRouterCircleSet[unit.Name].Properties, unit.Properties...)
		} else {
			// TODO:
			mapRouterCircleSet[unit.Name] = unit
			routerCircleSet.Units = append(routerCircleSet.Units, unit)
		}

		mapProperies[unit.Name] = map[string]*modules.CircleUnitProperty{}
		for _, property := range unit.Properties {
			mapProperies[unit.Name][property.Name] = property
		}
	}

	// for _, unit := range requestsCircleSet.Units {
	// 	if _, ok := mapRouterCircleSet[unit.Name]; ok {
	// 		// TODO:
	// 	} else {
	// 		// TODO:
	// 		cu := &modules.CircleUnit{
	// 			Name: unit.Name,
	// 		}
	// 		routerCircleSet.Units = append(routerCircleSet.Units, cu)
	// 	}

	// 	for _, property := range unit.Properties {
	// 		if _, ok := mapProperies[unit.Name]; !ok {
	// 			fmt.Println("속성 추가중에 ", unit.Name, " 없음 발생")
	// 			continue
	// 		}
	// 		mapProperies[unit.Name][property.Name] = property
	// 	}
	// 	//TODO: 속성 합치기
	// }

	// for _, unit := range responsesCircleSet.Units {
	// 	if _, ok := mapRouterCircleSet[unit.Name]; ok {
	// 		// TODO:
	// 	} else {
	// 		// TODO:
	// 		cu := &modules.CircleUnit{
	// 			Name: unit.Name,
	// 		}
	// 		routerCircleSet.Units = append(routerCircleSet.Units, cu)
	// 	}

	// 	//TODO: 속성 합치기
	// }

	return nil
}

func (cm *CircleManager) ImportCircleAdmin() (*modules.CircleSet, error) {
	flagRead := &FlagRead{}
	cs := &modules.CircleSet{}

	inFile, _ := os.Open(cm.MapTemplateSets["admin"].SourcePath)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	currentWhere := "meta"
	for scanner.Scan() {
		l := scanner.Text()
		l = strings.TrimSpace(l)
		curWhere(&currentWhere, l)

		scanLineForAdmin(flagRead, cs, &currentWhere, l)
	}

	return cs, nil
}

func (cm *CircleManager) ImportCircleRouter() (*modules.CircleSet, error) {
	flagRead := &FlagRead{}
	cs := &modules.CircleSet{}

	inFile, _ := os.Open(cm.MapTemplateSets["router"].SourcePath)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	currentWhere := "meta"
	for scanner.Scan() {
		l := scanner.Text()
		l = strings.TrimSpace(l)
		curWhere(&currentWhere, l)

		scanLineForRouter(flagRead, cs, &currentWhere, l)
	}

	return cs, nil
}

func scanLineForRouter(flagRead *FlagRead, cs *modules.CircleSet, currentWhere *string, l string) {
	if *currentWhere == "meta" {
		flagRead.RouterReadedAppTitle, cs.AppTitle = extract(flagRead.RouterReadedAppTitle, cs.AppTitle, l, "// @Title")
		flagRead.RouterReadedAppVersion, cs.AppVersion = extract(flagRead.RouterReadedAppVersion, cs.AppVersion, l, "// @APIVersion")
		flagRead.RouterReadedAppDescription, cs.AppDescription = extract(flagRead.RouterReadedAppDescription, cs.AppDescription, l, "// @Description")
		flagRead.RouterReadedAppContact, cs.AppContact = extract(flagRead.RouterReadedAppContact, cs.AppContact, l, "// @Contact")
		flagRead.RouterReadedAppTermsOfServiceUrl, cs.AppTermsOfServiceUrl = extract(flagRead.RouterReadedAppTermsOfServiceUrl, cs.AppTermsOfServiceUrl, l, "// @TermsOfServiceUrl")
		flagRead.RouterReadedAppLicense, cs.AppLicense = extract(flagRead.RouterReadedAppLicense, cs.AppLicense, l, "// @License")
		flagRead.RouterReadedAppSecurityDefinition, cs.AppSecurityDefinition = extract(flagRead.RouterReadedAppSecurityDefinition, cs.AppSecurityDefinition, l, "// @SecurityDefinition")
	}

	switch *currentWhere {
	case "manual", "auto":
		name := ""
		if strings.Index(l, "&modules.") >= 0 {
			name = getWord(l, "&modules.", "Controller{},")
		} else if strings.Index(l, "&controllers.") >= 0 {
			name = getWord(l, "&controllers.", "Controller{},")
		}

		if name == "" {
			return
		}
		cs.Units = append(cs.Units, &modules.CircleUnit{
			Name:      name,
			IsSystem:  *currentWhere == "system",
			IsManual:  *currentWhere == "manual",
			IsAutogen: *currentWhere == "auto",
		})
	}
}

func scanLineForAdmin(flagRead *FlagRead, cs *modules.CircleSet, currentWhere *string, l string) {
	switch *currentWhere {
	case "system", "manual", "auto":
		name := getWord(l, "addResourceAndMenu(&models.", "{}, \"")
		if name == "" {
			return
		}
		cs.Units = append(cs.Units, &modules.CircleUnit{
			Name:              name,
			MenuGroup:         getWord(l, "\", \"", "\", anyoneAllow, -1)"),
			MenuName:          getWord(l, "{}, \"", "\", \""),
			IsSystem:          *currentWhere == "system",
			IsManual:          *currentWhere == "manual",
			EnableAdminSource: true,
		})
	}
}

func scanSource(sourceType string, sourceDirPath string) (*modules.CircleSet, error) {
	cs := &modules.CircleSet{}

	if err := filepath.Walk(sourceDirPath, func(path string, info os.FileInfo, err error) error {
		fset := token.NewFileSet()
		d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
		if err != nil {
			fmt.Println(err)
			// First Error만 발생하고 진행됨
			// return err
		}

		for _, f := range d {
			p := doc.New(f, "./", 0)

			// TODO
			if sourceType == "models" ||
				sourceType == "requests" ||
				sourceType == "responses" {
				if err := scanSourceModel(sourceType, cs, p); err != nil {
					continue
				}
			} else if sourceType == "controllers" {
				if err := scanSourceForControllers(cs, p); err != nil {
					continue
				}
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return cs, nil
}

func scanSourceForControllers(cs *modules.CircleSet, p *doc.Package) error {
	for _, t := range p.Types {
		if strings.Index(t.Name, "Controller") < 0 {
			continue
		}

		isCreateble := false
		isUpdateble := false
		isDeleteble := false
		isGetAllable := false
		isGetOneable := false

		for _, method := range t.Methods {
			// requestCreateBodyName := ""
			// requestUpdateBodyName := ""
			// responseBodyName := ""
			// routerURL := ""
			routerMethod := ""
			//requestBodyName := ""

			if method.Name == "Post" {
				isCreateble = true
			} else if method.Name != "Put" {
				isUpdateble = true
			} else if method.Name != "GetOne" {
				isGetOneable = true
			} else if method.Name != "GetAll" {
				isGetAllable = true
			} else if method.Name != "Delete" {
				isDeleteble = true
			} else {
				//TODO: 다른 메소드 지원은 나중에
				continue
			}

			for _, docLine := range strings.Split(method.Doc, "\n") {
				if strings.Index(docLine, "@Param") == 0 {
					tempDocLine := strings.Replace(docLine, "\t", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "  ", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "  ", " ", -1)
					tempDocLineArray := strings.Split(tempDocLine, " ")
					if tempDocLineArray[1] == "body" {
						//equestBodyName = tempDocLineArray[3]
					}
				} else if strings.Index(docLine, "@Success") == 0 {
					tempDocLine := strings.Replace(docLine, "\t", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "  ", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "  ", " ", -1)
					tempDocLineArray := strings.Split(tempDocLine, " ")
					if tempDocLineArray[1] == "204" {
						continue
					}
					//responseBodyName = tempDocLineArray[2]
				} else if strings.Index(docLine, "@router") == 0 {
					tempDocLine := strings.Replace(docLine, "\t", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "  ", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "  ", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "]", " ", -1)
					tempDocLine = strings.Replace(tempDocLine, "[", " ", -1)
					tempDocLineArray := strings.Split(tempDocLine, " ")
					//routerURL = tempDocLineArray[1]
					routerMethod = tempDocLineArray[2]
				}

				if routerMethod == "post" {
					//requestCreateBodyName = requestBodyName
				} else if routerMethod == "put" {
					//requestUpdateBodyName = requestBodyName
				}
			}
		}
		name := strings.Replace(t.Name, "Controller", "", 1)
		cu := &modules.CircleUnit{
			Name: name,
			EnableControllerSource: true,
			IsCreateble:            isCreateble,
			IsUpdateble:            isUpdateble,
			IsGetAllable:           isGetAllable,
			IsGetOneable:           isGetOneable,
			IsDeleteble:            isDeleteble,
		}
		// TODO:
		// spew.Dump(t.Doc)
		// spew.Dump(t.Methods[0].Orig)
		// fmt.Println(t.Methods[0].Level)
		// fmt.Println(t.Methods[0].Recv)
		// fmt.Println(t.Methods[0].Name)

		cs.Units = append(cs.Units, cu)
	}
	return nil
}

func scanSourceModel(sourceType string, cs *modules.CircleSet, p *doc.Package) error {
	for _, t := range p.Types {
		if sourceType == "models" && strings.Index(t.Doc, "gen:qs") < 0 {
			continue
		}
		if _, ok := t.Decl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType); !ok {
			continue
		}

		structDecl := t.Decl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
		fields := structDecl.Fields.List
		cu := &modules.CircleUnit{
			Name:                 t.Name,
			EnableModelSource:    sourceType == "models",
			EnableRequestSource:  sourceType == "requests",
			EnableResponseSource: sourceType == "responses",
		}

		fmt.Println("Scan type...")
		fmt.Println("type.Name : ", t.Name)

		for _, field := range fields {
			isSystem := false
			fmt.Println("Scan field...")
			fmt.Println("field.Names : ", field.Names)
			if len(field.Names) <= 0 {
				continue
			}
			unitName := field.Names[0].Name
			switch unitName {
			case "ID", "CreatedAt", "UpdatedAt", "Name", "Description":
				isSystem = true
			}

			description := ""
			if field.Tag != nil {
				olnyTags := strings.Replace(field.Tag.Value, "`", "", -1)
				tags, err := structtag.Parse(olnyTags)
				if err != nil {
					return err
				}

				if descriptionTag, err := tags.Get("description"); err == nil {
					description = descriptionTag.Name
				} else {
					fmt.Println("설명이 없습니다.")
				}
			}

			typeName := ""
			isNull := false
			if _, ok := field.Type.(*ast.Ident); ok {
				typeName = field.Type.(*ast.Ident).Name
			} else if _, ok := field.Type.(*ast.SelectorExpr); ok {
				typeX := field.Type.(*ast.SelectorExpr).X.(*ast.Ident).Name
				typeSel := field.Type.(*ast.SelectorExpr).Sel.Name
				typeName = fmt.Sprintf("%s.%s", typeX, typeSel)
			} else if _, ok := field.Type.(*ast.StarExpr); ok {
				typeX := field.Type.(*ast.StarExpr).X.(*ast.SelectorExpr).X.(*ast.Ident).Name
				typeSel := field.Type.(*ast.StarExpr).X.(*ast.SelectorExpr).Sel.Name
				typeName = fmt.Sprintf("*%s.%s", typeX, typeSel)
				isNull = true
			} else {
				spew.Dump(field.Type)
			}

			cu.Properties = append(cu.Properties, &modules.CircleUnitProperty{
				Name:        unitName,
				Description: description,
				Type:        typeName,
				Nullable:    isNull,
				IsEnable:    true,
				IsManual:    true,
				IsSystem:    isSystem,
			})
		}
		cs.Units = append(cs.Units, cu)
	}
	return nil
}

func curWhere(currentWhere *string, l string) {
	tempCurrentWhere := ""
	if strings.Index(l, "// circle:system:start") == 0 {
		tempCurrentWhere = strings.Replace("circle:system", "circle:", "", -1)
	} else if strings.Index(l, "// circle:manual:start") == 0 {
		tempCurrentWhere = strings.Replace("circle:manual", "circle:", "", -1)
	} else if strings.Index(l, "// circle:auto:start") == 0 {
		tempCurrentWhere = strings.Replace("circle:auto", "circle:", "", -1)
	}
	if tempCurrentWhere != "" {
		*currentWhere = tempCurrentWhere
	}
}

func extract(isReaded bool, value string, line string, target string) (bool, string) {
	if isReaded {
		return true, value
	}
	if strings.Index(line, target) >= 0 {
		replace := strings.Replace(line, target, "", -1)
		return true, strings.TrimSpace(replace)
	}
	return false, ""
}

func getWord(s string, start string, end string) string {
	startIndex := strings.Index(s, start) + len(start)
	endIndex := strings.Index(s, end)
	if end == "" {
		endIndex = len(s)
	}

	if startIndex < 0 || endIndex < 0 {
		return ""
	}

	return s[startIndex:endIndex]
}

func removeRigth(s, indexChar string) string {
	if i := strings.Index(s, indexChar); i >= 0 {
		return s[0:i]
	}
	return s
}
