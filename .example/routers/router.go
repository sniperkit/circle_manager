// @APIVersion 10.1.1
// @Title Circle
// @Description wow
// @Contact myapp@myapp.com
// @TermsOfServiceUrl http://circle.circle
// @License MIT
// @SecurityDefinition "userAPIKey apiKey X-USER-AUTH-TOKEN header "I love auto-generated docs"
package routers

import (
	"github.com/astaxie/beego"
	"github.com/jungju/circle/controllers"
)

func init() {
	// circle:system:start
	ns := beego.NewNamespace("/v1",
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
		beego.NSNamespace("/webhooks",
			beego.NSInclude(
				&controllers.WebhookController{},
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
