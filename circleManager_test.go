package circle_manager

import (
	"testing"

	"github.com/jinzhu/inflection"
	"github.com/jungju/circle/utils"
)

func TestGo(t *testing.T) {
	cm := &CircleManager{}
	base = "test"

	cm.GeneateSourceBySet(&CircleSet{
		Name: "Office1",
		Units: []CircleUnit{
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
	})
}

func makeCircleUnit(name string, menuName string, menuGroup string, properties ...CircleUnitProperty) CircleUnit {
	return CircleUnit{
		Name:         name,
		VariableName: utils.MakeFirstLowerCase(name),
		URL:          inflection.Plural(utils.MakeFirstLowerCase(name)),
		MenuName:     menuName,
		MenuGroup:    menuGroup,
		Properties:   properties,
	}
}

func makeCircleUnitProperty(name string, typeName string) CircleUnitProperty {
	return CircleUnitProperty{
		Name: name,
		Type: typeName,
	}
}
