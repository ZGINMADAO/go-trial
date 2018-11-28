package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"fmt"
)

type AuthController struct {
	Ctx iris.Context
}

func (_ *AuthController) GetLogin() mvc.View {
	return mvc.View{
		Name: "admin/login.html",
		Data: iris.Map{"url": "/admin/login"},
	}
}

func (my *AuthController) PostLogin() {
	username := my.Ctx.Params().Get("username")
	password := my.Ctx.Params().Get("password")
	fmt.Println(username, password)
}
