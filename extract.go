package main

// type AppInfo struct {
// 	APIVersion         string
// 	Title              string
// 	Description        string
// 	Contact            string
// 	TermsOfServiceUrl  string
// 	License            string
// 	SecurityDefinition string
// }

// func (cm *CircleManager) GeneateSourceBySet(cs *CircleSet) error {
// 	circleSet = cs
// 	cm.prepare()

// 	// mapDBUnit := map[string]CircleUnit{}
// 	// for _, dbUnit := range cs.Units {
// 	// 	mapDBUnit[dbUnit.Name] = dbUnit
// 	// }

// 	// newCS := &CircleSet{
// 	// 	Units: []CircleUnit{},
// 	// }
// 	// if cm.ByType == "db" {

// 	// } else if cm.ByType == "source" {

// 	// }

// 	// mapUpdateUnit := map[string]CircleUnit{}
// 	// for _, circleTemplateSet := range cm.MapTemplateSets {
// 	// 	for _, sourceUnit := range circleTemplateSet.Extract() {
// 	// 		if circleTemplateSet.IsMulti {

// 	// 		}

// 	// 		var newUnit CircleUnit
// 	// 		if dbUnit, ok := mapDBUnit[sourceUnit.Name]; ok {
// 	// 			if cm.ByType == "db" {
// 	// 				newUnit = merge(dbUnit, sourceUnit)
// 	// 			} else if cm.ByType == "source" {
// 	// 				newUnit = merge(sourceUnit, dbUnit)
// 	// 			}
// 	// 		} else {
// 	// 			newUnit = sourceUnit
// 	// 		}
// 	// 		mapUpdateUnit[newUnit.Name] = newUnit
// 	// 	}

// 	// }

// 	// mapUpdateUnit := map[string]CircleUnit{}
// 	// for _, circleTemplateSet := range []CircleTemplateSet{
// 	// 	cm.QORAdminTemplateSet,
// 	// 	cm.RouterTemplateSet,
// 	// } {
// 	// 	for _, sourceUnit := range circleTemplateSet.Extract() {
// 	// 		var newUnit CircleUnit
// 	// 		if dbUnit, ok := mapDBUnit[sourceUnit.Name]; ok {
// 	// 			if cm.ByType == "db" {
// 	// 				newUnit = merge(dbUnit, sourceUnit)
// 	// 			} else if cm.ByType == "source" {
// 	// 				newUnit = merge(sourceUnit, dbUnit)
// 	// 			}
// 	// 		} else {
// 	// 			newUnit = sourceUnit
// 	// 		}
// 	// 		mapUpdateUnit[newUnit.Name] = newUnit
// 	// 	}

// 	// }

// 	// gen(cm.ModelsDir, makePath("models.tmpl"), cs)
// 	// gen(cm.ControllersDir, makePath("controllers.tmpl"), cs)
// 	// gen(cm.RequestsBodyDir, makePath("requests.tmpl"), cs)
// 	// gen(cm.ResponseBodyDir, makePath("responses.tmpl"), cs)

// 	return nil
// }

// func merge(baseCU modules.CircleUnit, changeCU modules.CircleUnit) modules.CircleUnit {
// 	setOnlyExistValue := func(baseString, newValue string) string {
// 		if newValue == "" {
// 			return BaseString
// 		}
// 		return newValue
// 	}
// 	setOnlyExistValueForBool := func(baseString, newValue bool) bool {
// 		if newValue == false {
// 			return BaseString
// 		}
// 		return newValue
// 	}
// 	changeCU.Name = setOnlyExistValue(changeCU.Name, BaseCU.Name)
// 	changeCU.Description = setOnlyExistValue(changeCU.Description, BaseCU.Description)
// 	changeCU.ControllerName = setOnlyExistValue(changeCU.ControllerName, BaseCU.ControllerName)
// 	changeCU.Import = setOnlyExistValue(changeCU.Import, BaseCU.Import)
// 	changeCU.URL = setOnlyExistValue(changeCU.URL, BaseCU.URL)
// 	changeCU.MenuName = setOnlyExistValue(changeCU.MenuName, BaseCU.MenuName)
// 	changeCU.MenuGroup = setOnlyExistValue(changeCU.MenuGroup, BaseCU.MenuGroup)
// 	changeCU.IsEnable = setOnlyExistValueForBool(changeCU.IsEnable, BaseCU.IsEnable)
// 	changeCU.IsManual = setOnlyExistValueForBool(changeCU.IsManual, BaseCU.IsManual)
// 	changeCU.IsSystem = setOnlyExistValueForBool(changeCU.IsSystem, BaseCU.IsSystem)
// 	return changeCU
// }