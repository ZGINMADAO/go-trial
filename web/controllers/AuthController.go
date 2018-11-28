package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	//"fmt"
	"fmt"
	"go-trial/services"
)

type AuthController struct {
	Ctx  iris.Context
	auth services.AuthService
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
	data := my.auth.Get(1)
	fmt.Println(data)
	fmt.Println(username, password)
}
