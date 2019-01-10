package admin

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/datamodels"
	adminService "go-trial/services/admin"
	"go-trial/units"
	"github.com/kataras/iris/sessions"
	"go-trial/web/controllers"
)

type AuthController struct {
	controllers.BaseController
	Session *sessions.Session
}

func (my *AuthController) GetLogin() mvc.View {
	return mvc.View{
		Name: "admin/login.html",
		Data: iris.Map{"url": "/admin/login"},
	}
}

func (my *AuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Dependencies().Add(adminService.NewAuth())
}

func (my *AuthController) BeginRequest(ctx iris.Context) {
	//fmt.Println("BeginRequest被调用了啊哈啊哈哈")
}

func (my *AuthController) EndRequest(ctx iris.Context) {
	//fmt.Println("EndRequest被调用了啊哈啊哈哈")
}

func (my *AuthController) PostLogin() {

	username := my.Ctx.PostValue("username")
	password := my.Ctx.PostValue("password")

	if len(username) < 5 || len(password) < 5 {
		my.ReturnJson(false, nil, "用户或密码不能少于5位")
	} else {
		//apiResource(true, nil, "登录成功")
		var admin datamodels.Admin
		admin.Username = username
		admin.Password = password
		ok, adminErr := my.DB.Get(&admin)
		if adminErr != nil {
			my.Ctx.Application().Logger().Warn(adminErr.Error())
			return
		}

		if ok == true {
			my.Session.Set("admin_session_profile", admin)
			my.Ctx.JSON(iris.Map{"status": true, "data": nil, "message": "登录成功"})
		} else {
			my.ReturnJson(false, nil, "登录失败")
		}
	}
}

func (my *AuthController) Get() mvc.View {
	admin := my.Session.Get("admin_session_profile")
	var result []datamodels.Tree
	my.DB.Find(&result)
	deepResult := make([]units.List, 0)
	units.Recursive(result, 0, &deepResult)
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{"Tree": deepResult, "AdminInfo": admin},
	}
}
