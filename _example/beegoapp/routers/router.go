// @APIVersion 0.1.10
// @Title Circle
// @Description wow
// @Contact leejungju.go@gmail.com
// @TermsOfServiceUrl http://circle.land
// @License MIT
// @SecurityDefinition userAPIKey apiKey X-USER-AUTH-TOKEN header "I love auto-generated docs"
package routers

import (
	"github.com/astaxie/beego"
	"github.com/jungju/circle_manager/_example/beegoapp/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1")
	beego.NSNamespace("/webhooks",
		beego.NSInclude(
			&controllers.WebhookController{},
		),
	)
	// // circle:system:end

	// // circle:manual:start
	// // circle:manual:end

	// // circle:auto:start

	// // circle:auto:end

	beego.AddNamespace(ns)
}
