package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type UserController struct {
	Ctx iris.Context
}


func (_ *UserController) Get() mvc.View {
	return mvc.View{
		Name: "user/register.html",
		Data: iris.Map{"Title": "User Index"},
	}
}
func (_ *UserController) PostAdd() string {
	return "success"
}
