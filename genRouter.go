package circle_manager

import (
	"fmt"
	"strings"
)

type RouterFile struct {
	APIVersion         string
	Title              string
	Description        string
	Contact            string
	TermsOfServiceUrl  string
	License            string
	SecurityDefinition string
	RootURL            string
	ImportRoot         string
	MapRouterItems     map[string][]RouterItem
	RawFilePath        string
	CircleSet          *CircleSet
	TemplatePath       string
}

type RouterItem struct {
	Name           string
	ControllerName string
	URL            string
	Type           string
}

func (fm *RouterFile) Export(unit *CircleUnit) error {
	fm.MapRouterItems["auto"] = append(fm.MapRouterItems["auto"], RouterItem{
		ControllerName: fmt.Sprintf("%sController", unit.Name),
		URL:            unit.URL,
		Type:           "auto",
	})
	return nil
}

func (fm *RouterFile) Import(status, l string) {
	fm.APIVersion = replaceAndTrim(fm.APIVersion, l, "// @APIVersion")
	fm.Title = replaceAndTrim(fm.Title, l, "// @Title")
	fm.Description = replaceAndTrim(fm.Description, l, "// @Description")
	fm.Contact = replaceAndTrim(fm.Contact, l, "// @Contact")
	fm.TermsOfServiceUrl = replaceAndTrim(fm.TermsOfServiceUrl, l, "// @TermsOfServiceUrl")
	fm.License = replaceAndTrim(fm.License, l, "// @License")
	fm.SecurityDefinition = replaceAndTrim(fm.SecurityDefinition, l, "// @SecurityDefinition")

	if targetLine(l, "beego.NSNamespace(") {
		fm.MapRouterItems[status] = append(fm.MapRouterItems[status], RouterItem{
			ControllerName: getWord(l, "beego.NSInclude(&controllers.", "Controller{}))"),
			URL:            getWord(l, "beego.NSNamespace(\"", "\", beego.NSInclude(&controllers."),
		})
	}
}

func replaceAndTrim(target string, s string, old string) string {
	if target != "" || strings.Index(s, old) != 0 {
		return target
	}

	ret := strings.Replace(s, old, "", -1)
	return strings.TrimSpace(ret)
}

func getWord(s string, start string, end string) string {
	startIndex := strings.Index(s, start) + len(start)
	endIndex := strings.Index(s, end)

	return s[startIndex:endIndex]
}

func setStatus(rawSatatus, l, targetLine, newStatus string) string {
	if strings.Index(l, targetLine) == 0 {
		return strings.Replace(newStatus, "circle:", "", -1)
	}
	return rawSatatus
}

func targetLine(l, target string) bool {
	return strings.Index(l, target) == 0
}
