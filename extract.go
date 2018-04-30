package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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

func (cm *CircleManager) ImportCircle() error {
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

	return nil
}

func (cm *CircleManager) ImportCircleUnit(name string) error {
	flagRead := &FlagRead{}
	cu := &modules.CircleUnit{
		Name: name,
	}

	inFile, _ := os.Open(
		filepath.Join(
			cm.MapTemplateSets["models"].SourcePath,
			fmt.Sprintf("%s.go", name),
		),
	)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	currentWhere := ""
	for scanner.Scan() {
		l := scanner.Text()
		l = strings.TrimSpace(l)

		scanLineForModel(flagRead, cu, &currentWhere, l)
	}

	return nil
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
	case "system", "manual", "auto":
		name := getWord(l, "&modules.", "Controller{},")
		if name == "" {
			return
		}
		cs.Units = append(cs.Units, modules.CircleUnit{
			Name:     name,
			IsSystem: *currentWhere == "system",
			IsManual: *currentWhere == "manual",
		})
	}
}

func scanLineForController(flagRead *FlagRead, cu *modules.CircleUnit, currentWhere *string, l string) {
	if *currentWhere == "end_api" && strings.Index(l, fmt.Sprintf("type %s struct {", cu.Name)) >= 0 {
		*currentWhere = "in_api"
		return
	}

	if *currentWhere == "in_api" && strings.Index(l, "}") >= 0 {
		*currentWhere = "end_api"
		return
	}

	if *currentWhere != "end_api" {
		return
	}

	tags := ""
	re := regexp.MustCompile("\\`(.*)`")
	if match := re.FindStringSubmatch(l); len(match) >= 1 {
		tags = match[1]
	}

	sl := strings.Replace(l, "  ", " ", -1)
	sl = strings.Replace(sl, "  ", " ", -1)
	sl = strings.Replace(sl, "  ", " ", -1)
	slArr := strings.Split(sl, " ")

	if len(slArr) >= 2 {
		isSystem := false
		switch slArr[0] {
		case "ID", "CreatedAt", "UpdatedAt", "Name", "Description":
			isSystem = true
		}
		description := ""
		if len(slArr) >= 3 {
			description = getWord(tags, "description:\"", "")
			description = removeRigth(description, "\"")
		}
		cu.Properties = append(cu.Properties, modules.CircleUnitProperty{
			Name:        slArr[0],
			Description: description,
			//CircleUnit                     CircleUnit
			//CircleUnitID                   uint
			Type: slArr[1],
			//Nullable                       bool
			IsEnable: true,
			IsManual: true,
			IsSystem: isSystem,
		})
	} else {
		fmt.Println("알수없는 형식 : ", sl)
	}
}

func scanLineForModel(flagRead *FlagRead, cu *modules.CircleUnit, currentWhere *string, l string) {
	if *currentWhere == "end_model" && strings.Index(l, fmt.Sprintf("type %s struct {", cu.Name)) >= 0 {
		*currentWhere = "in_model"
		return
	}

	if *currentWhere == "in_model" && strings.Index(l, "}") >= 0 {
		*currentWhere = "end_model"
		return
	}

	if *currentWhere != "in_model" {
		return
	}

	tags := ""
	re := regexp.MustCompile("\\`(.*)`")
	if match := re.FindStringSubmatch(l); len(match) >= 1 {
		tags = match[1]
	}

	sl := strings.Replace(l, "  ", " ", -1)
	sl = strings.Replace(sl, "  ", " ", -1)
	sl = strings.Replace(sl, "  ", " ", -1)
	slArr := strings.Split(sl, " ")

	if len(slArr) >= 2 {
		isSystem := false
		switch slArr[0] {
		case "ID", "CreatedAt", "UpdatedAt", "Name", "Description":
			isSystem = true
		}
		description := ""
		if len(slArr) >= 3 {
			description = getWord(tags, "description:\"", "")
			description = removeRigth(description, "\"")
		}
		cu.Properties = append(cu.Properties, modules.CircleUnitProperty{
			Name:        slArr[0],
			Description: description,
			//CircleUnit                     CircleUnit
			//CircleUnitID                   uint
			Type: slArr[1],
			//Nullable                       bool
			IsEnable: true,
			IsManual: true,
			IsSystem: isSystem,
		})
	} else {
		fmt.Println("알수없는 형식 : ", sl)
	}
}

func scanLineForAdmin(flagRead *FlagRead, cs *modules.CircleSet, currentWhere *string, l string) {
	switch *currentWhere {
	case "system", "manual", "auto":
		name := getWord(l, "addResourceAndMenu(&models.", "{}, \"")
		if name == "" {
			return
		}
		cs.Units = append(cs.Units, modules.CircleUnit{
			Name:      name,
			MenuGroup: getWord(l, "\", \"", "\", anyoneAllow, -1)"),
			MenuName:  getWord(l, "{}, \"", "\", \""),
			IsSystem:  *currentWhere == "system",
			IsManual:  *currentWhere == "manual",
		})
	}
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
