// @APIVersion 10.1.		// circle:auto:end
		beego.NSNamespace("/tests",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
		// circle:auto:end
	)

	beego.AddNamespace(ns)
}
