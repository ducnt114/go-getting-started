package router

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"go-getting-started/controller"
)

func InitRouter() {
	web.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	web.Router("/user", &controller.UserController{})
}
