package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/web/controllers"
	"go-trial/datasource"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/sessions"
	"time"
	//"github.com/kataras/iris/hero"
	"go-trial/web/middleware"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/websocket"
	//"go-trial/web/socket"
)

var DB *xorm.Engine

var SessionManage *sessions.Sessions

func main() {

	SessionManage = sessions.New(sessions.Config{
		Cookie:       "site_session_id",
		Expires:      60 * time.Minute,
		AllowReclaim: true,
	})

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

	tpl := iris.HTML("./web/views", ".html").Layout("admin/layout.html").Reload(true)
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

	//app.Use(middleware.BasicAuth)
	mvc.Configure(app.Party("/admin"), AdminMvc)
	mvc.Configure(app.Party("/test"), TestMvc)
	mvc.Configure(app.Party("/socket"), SocketMvc)

	//socket.StartSocket(app)

	app.Run(iris.Addr("0.0.0.0:8080"))
}

func AdminMvc(app *mvc.Application) {

	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Admin Path: %s", ctx.Path())
		ctx.Next()
	})

	session := SessionManage.Start

	hero.Register(session)
	app.Router.Use(hero.Handler(middleware.SessionAuth))

	app.Register(DB, session)
	//app.Register(SessionManage.Start)
	app.Router.Get("/mailbox", func(ctx iris.Context) {
		ctx.View("admin/mailbox.html")
	})
	app.Party("/auth").Handle(new(controllers.AuthController))
	app.Party("/product").Handle(new(controllers.ProductController))
	app.Party("/tool").Handle(new(controllers.ToolController))
	app.Party("/order").Handle(new(controllers.OrderController))
	app.Party("/system").Handle(new(controllers.SystemController))
}


func TestMvc(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Test Path: %s", ctx.Path())
		ctx.Next()
	})
	app.Handle(new(controllers.TestController))
}

func SocketMvc(app *mvc.Application) {

	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})

	app.Router.Any("/iris-ws.js", websocket.ClientHandler())

	//app.Router.Get("/echo", ws.Handler())
	//app.Router.Get("/demo", ws.Handler())
	//ws.OnConnection(handleConnection)

	app.Register(ws.Upgrade)
	app.Handle(new(controllers.SocketController))
}



//func handleConnection(c websocket.Connection) {
//
//	fmt.Println("onConnect")
//	// Read events from browser
//	c.On("chat", func(msg string) {
//		// Print the message to the console, c.Context() is the iris's http context.
//		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
//		// Write message back to the client message owner with:
//		// c.Emit("chat", msg)
//		// Write message to all except this client with:
//		c.To(websocket.Broadcast).Emit("chat", msg)
//	})
//}
