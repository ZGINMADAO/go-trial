package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	mvc.Configure(app.Party("/user"), UserMvc)
	app.Run(iris.Addr("0.0.0.0:8080"))
}

func UserMvc(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
	})
	app.Handle(new(controllers.UserController))
}
