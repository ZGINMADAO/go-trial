package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/services"
)

type AuthController struct {
	ApiJson
	Ctx  iris.Context
	Auth services.AuthService
}

func (_ *AuthController) GetLogin() mvc.View {
	return mvc.View{
		Name: "admin/login.html",
		Data: iris.Map{"url": "/admin/login"},
	}
}

func (my *AuthController) PostLogin() {

	username := my.Ctx.PostValue("username")
	password := my.Ctx.PostValue("password")

	ok := my.Auth.Search(username, password)
	if ok == true {
		my.Ctx.JSON(apiResource(true, nil, "登录成功"))
	} else {
		my.Ctx.JSON(apiResource(false, nil, "登录失败"))
	}

}
