package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/web/controllers"
)

func main() {
	app := iris.New()
	//app.Logger().SetLevel("debug")
	app.StaticWeb("/public","./web/public")
	app.RegisterView(iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true))

	mvc.Configure(app.Party("/user"), UserMvc)
	app.Run(iris.Addr("0.0.0.0:8080"))
}

func UserMvc(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	app.Handle(new(controllers.UserController))
}
