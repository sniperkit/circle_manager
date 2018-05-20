package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jungju/circle_manager/modules"
	"github.com/stretchr/testify/assert"
)

func TestCleanRouterSource(t *testing.T) {
	regenRouterSource := cleanRouterSource(routerSource)
	i := strings.Index(regenRouterSource, CIRCLE_AUTO_START_WORD+"\n\t\t"+CIRCLE_AUTO_END_WORD)
	assert.NotEqual(t, -1, i)
	regenRouterSource = strings.TrimSpace(strings.Replace(regenRouterSource, "\n", "", -1))
	assert.Equal(t, "}", regenRouterSource[len(regenRouterSource)-1:len(regenRouterSource)])
}

func TestRemoveRouterSource(t *testing.T) {
	regenRouterSource := removeRouterSource(routerSource, "Team")
	//fmt.Println(regenRouterSource)
	i := strings.Index(regenRouterSource, "Team")
	assert.Equal(t, -1, i)
	i = strings.Index(regenRouterSource, "beego.AddNamespace(ns)")
	assert.NotEqual(t, -1, i)
}

func TestSaveRouterSource(t *testing.T) {
	unit := &modules.CircleUnit{
		Name: "Robot",
	}

	source, err := saveRouterSource(routerTemplate, unit)
	assert.Nil(t, err)
	assert.NotEqual(t, -1, strings.Index(source, fmt.Sprintf("%sController", unit.Name)))
	assert.NotEqual(t, -1, strings.Index(source, "robots"))
}

func TestGenerateRouter(t *testing.T) {
	envs = &Envs{
		AppDir: "_example",
		Mode:   "gen",
	}
	newRouterSource, err := generateRouter(routerSource, testCS)
	assert.Nil(t, err)
	assert.NotEqual(t, -1, strings.Index(newRouterSource, testCS.Units[0].Name))
}

func TestGo(t *testing.T) {
	cm := &CircleManager{}
	envs = &Envs{
		AppDir: "_example",
	}

	err := cm.GenerateSource(testCS)
	if err != nil {
		assert.Fail(t, err.Error())
	}
}

func makeCircleUnit(name string, menuName string, menuGroup string, properties ...*modules.CircleUnitProperty) *modules.CircleUnit {
	return &modules.CircleUnit{
		Name:                   name,
		MenuName:               menuName,
		MenuGroup:              menuGroup,
		Properties:             properties,
		EnableAdminSource:      true,
		EnableModelSource:      true,
		EnableControllerSource: true,
		EnableRequestSource:    true,
		EnableResponseSource:   true,
	}
}

func makeCircleUnitProperty(name string, typeName string) *modules.CircleUnitProperty {
	return &modules.CircleUnitProperty{
		Name: name,
		Type: typeName,
	}
}

var testCS = &modules.CircleSet{
	Name:                  "Office1",
	Import:                "jungju/circle",
	AppVersion:            "10.1.1",
	AppTitle:              "Circle",
	AppDescription:        "wow",
	AppContact:            "myapp@myapp.com",
	AppTermsOfServiceUrl:  "http://circle.circle",
	AppLicense:            "MIT",
	AppSecurityDefinition: `"userAPIKey apiKey X-USER-AUTH-TOKEN header "I love auto-generated docs"`,

	Units: []*modules.CircleUnit{
		makeCircleUnit(
			"GithubCommit", "Commits", "이벤트관리",
			makeCircleUnitProperty("RepoName", "string"),
			makeCircleUnitProperty("Comments", "string"),
			makeCircleUnitProperty("UserName", "string"),
			makeCircleUnitProperty("BranchName", "string"),
		),
		makeCircleUnit(
			"GithubRelease", "Releases", "이벤트관리",
			makeCircleUnitProperty("RepoName", "string"),
			makeCircleUnitProperty("TagName", "string"),
			makeCircleUnitProperty("UserName", "string"),
			makeCircleUnitProperty("PreRelease", "bool"),
			makeCircleUnitProperty("Message", "string"),
		),
		makeCircleUnit(
			"Event", "이벤트", "이벤트관리",
			makeCircleUnitProperty("EventCreated", "time.Time"),
			makeCircleUnitProperty("EventEnds", "*time.Time"),
			makeCircleUnitProperty("Summary", "string"),
			makeCircleUnitProperty("Organizer", "string"),
			makeCircleUnitProperty("EventUser", "string"),
			makeCircleUnitProperty("EventBegins", "time.Time"),
			makeCircleUnitProperty("EventID", "string"),
			makeCircleUnitProperty("Location", "string"),
			makeCircleUnitProperty("Source", "string"),
			makeCircleUnitProperty("Attendees", "string"),
			makeCircleUnitProperty("GithubRelease", "models.GithubRelease"),
			makeCircleUnitProperty("GithubReleaseID", "uint"),
		),
		makeCircleUnit(
			"Employee", "직원", "이벤트관리",
			makeCircleUnitProperty("OriginName", "string"),
		),
		makeCircleUnit(
			"KeyEvent", "주요일정", "이벤트관리",
			makeCircleUnitProperty("EventDate", "time.Time"),
		),
		makeCircleUnit(
			"Project", "프로젝트", "이벤트관리",
			makeCircleUnitProperty("Status", "string"),
		),
		makeCircleUnit(
			"Todo", "할일", "이벤트관리",
			makeCircleUnitProperty("ListID", "string"),
			makeCircleUnitProperty("ListName", "string"),
			makeCircleUnitProperty("Status", "string"),
			makeCircleUnitProperty("CardID", "string"),
			makeCircleUnitProperty("BoardID", "string"),
			makeCircleUnitProperty("BoardName", "string"),
			makeCircleUnitProperty("Source", "string"),
		),
		makeCircleUnit(
			"Team", "팀", "이벤트관리",
		),
	},
}

var routerSource = `// @APIVersion 0.1.10
// @Title Circle
// @Description wow
// @Contact leejungju.go@gmail.com
// @TermsOfServiceUrl http://circle.circle
// @License MIT
// @SecurityDefinition userAPIKey apiKey X-USER-AUTH-TOKEN header "I love auto-generated docs"
package routers

import (
	"github.com/astaxie/beego"
	"github.com/jungju/circle_manager/modules"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// circle:system:start
		beego.NSNamespace("/circleSets",
			beego.NSInclude(
				&modules.CircleSetController{},
			),
		),
		beego.NSNamespace("/circleUnits",
			beego.NSInclude(
				&modules.CircleUnitController{},
			),
		),
		beego.NSNamespace("/circleUnitProperties",
			beego.NSInclude(
				&modules.CircleUnitPropertyController{},
			),
		),
		beego.NSNamespace("/notifications",
			beego.NSInclude(
				&modules.NotificationController{},
			),
		),
		beego.NSNamespace("/notificationTypes",
			beego.NSInclude(
				&modules.NotificationTypeController{},
			),
		),
		// circle:system:end

		// circle:manual:start
		// circle:manual:end

		// circle:auto:start
		// circle:auto:end
	)
	beego.AddNamespace(ns)
}


`
