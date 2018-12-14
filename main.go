package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/web/controllers"
	"go-trial/datasource"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

func main() {
	DB = datasource.Instance()

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

	tpl := iris.HTML("./web/views", ".html").Reload(true)
	app.RegisterView(tpl)

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

	app.Register(DB)
	app.Router.Get("/mailbox", func(ctx iris.Context) {
		ctx.View("admin/mailbox.html")
	})
	app.Party("/auth").Handle(new(controllers.AuthController))
	app.Party("/product").Handle(new(controllers.ProductController))
	app.Party("/tool").Handle(new(controllers.ToolController))
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
