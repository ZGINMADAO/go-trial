package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/web/controllers"
	"go-trial/services"
)

func main() {

	app := iris.New()

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.WriteString("404")
	})

	app.OnErrorCode(500, func(ctx iris.Context) {
		ctx.WriteString("500")
	})

	//app.Logger().SetLevel("debug")
	app.Favicon("./web/public/favicon.ico")
	app.StaticWeb("/public", "./web/public")
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))

	mvc.Configure(app.Party("/admin"), AdminMvc)
	mvc.Configure(app.Party("/home"), HomeMvc)
	mvc.Configure(app.Party("/test"), TestMvc)
	app.Run(iris.Addr("0.0.0.0:8080"))
}

func AdminMvc(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Admin Path: %s", ctx.Path())
		ctx.Next()
	})

	authService := services.NewAuth()
	app.Register(authService)

	app.Handle(new(controllers.AuthController))
}

func HomeMvc(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Home Path: %s", ctx.Path())
		ctx.Next()
	})
	app.Handle(new(controllers.UserController))
}

func TestMvc(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Test Path: %s", ctx.Path())
		ctx.Next()
	})
	app.Handle(new(controllers.TestController))
}
