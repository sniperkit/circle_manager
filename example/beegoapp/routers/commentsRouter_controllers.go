package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EmployeeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:EventController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubCommitController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:GithubReleaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:IcsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:KeyEventController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:SprintController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TeamController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TodoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:TrelloController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserAuthController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserAuthController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserAuthController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserAuthController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:UserController"],
		beego.ControllerComments{
			Method: "Me",
			Router: `/me`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:WebhookController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:WebhookController"],
		beego.ControllerComments{
			Method: "CalendarEvent",
			Router: `/calendar/events/google`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:WebhookController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:WebhookController"],
		beego.ControllerComments{
			Method: "PostGithub",
			Router: `/github`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:WebhookController"] = append(beego.GlobalControllerRouter["github.com/jungju/circle_manager/example/beegoapp/controllers:WebhookController"],
		beego.ControllerComments{
			Method: "Sync",
			Router: `/sync`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
