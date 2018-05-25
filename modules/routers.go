package modules

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func RegistRouterForSystemStatus() error {
	beego.Get("/healthcheck", func(ctx *context.Context) {
		ctx.Output.Body([]byte("ok"))
	})
	return nil
}

func RegistRouterForEnvs(prodEnv string, version string) error {
	beego.Get("/system/envs", func(ctx *context.Context) {
		ctx.Output.Body([]byte(prodEnv))
	})
	return nil
}
