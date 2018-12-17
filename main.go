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

	//fileServer := app.StaticHandler("./web/public", false, false)
	//app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	//	path := r.URL.Path
	//	fmt.Println(path)
	//	//请注意，如果path的后缀为“index.html”，则会自动重定向到“/”，
	//	//所以我们的第一个处理程序将被执行。
	//	if !strings.Contains(path, ".") {
	//		//如果它不是资源，那么就像正常情况一样继续使用路由器. <-- IMPORTANT
	//		router(w, r)
	//		return
	//	}
	//
	//	fmt.Println("is here")
	//	//获取并释放上下文，以便使用它来执行我们的文件服务器
	//	//记得：我们使用net / http.Handler，因为我们在路由器本身之前处于“低级别”。
	//	ctx := app.ContextPool.Acquire(w, r)
	//	fileServer(ctx)
	//	app.ContextPool.Release(ctx)
	//})

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
