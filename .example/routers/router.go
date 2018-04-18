// @APIVersion
// @Title
// @Description
// @Contact
// @TermsOfServiceUrl
// @License
// @SecurityDefinition
package routers

import (
	"github.com//controllers"
	"github.com/astaxie/beego"
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
