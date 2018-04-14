// @APIVersion
// @Title
// @Description
// @Contact
// @TermsOfServiceUrl
// @License
// @SecurityDefinition
package routers

import (
	"github.com/astaxie/beego"
	"github.com/jungju/demo/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/cars",
			beego.NSInclude(
				&controllers.CarController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
