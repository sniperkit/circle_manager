// @APIVersion 1.0.0
// @Title Circle API
// @Description Circle API
// @Contact leejungju.go@gmail.com
// @TermsOfServiceUrl http://circle.com
// @License Private
// @SecurityDefinition userAPIKey apiKey X-USER-AUTH-TOKEN header "I love auto-generated docs
package routers

import (
	"github.com/astaxie/beego"
	"github.com/jungju/circle/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/githubCommits",
			beego.NSInclude(
				&controllers.GithubCommitController{},
			),
		),
		beego.NSNamespace("/githubReleases",
			beego.NSInclude(
				&controllers.GithubReleaseController{},
			),
		),
		beego.NSNamespace("/events",
			beego.NSInclude(
				&controllers.EventController{},
			),
		),
		beego.NSNamespace("/employees",
			beego.NSInclude(
				&controllers.EmployeeController{},
			),
		),
		beego.NSNamespace("/keyEvents",
			beego.NSInclude(
				&controllers.KeyEventController{},
			),
		),
		beego.NSNamespace("/projects",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),
		beego.NSNamespace("/todos",
			beego.NSInclude(
				&controllers.TodoController{},
			),
		),
		beego.NSNamespace("/teams",
			beego.NSInclude(
				&controllers.TeamController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
