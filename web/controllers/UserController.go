package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-trial/datasource"
	"go-trial/repositories"
	"fmt"
)

type UserController struct {
	Ctx iris.Context
}

func (_ *UserController) Get() mvc.View {
	return mvc.View{
		Name:   "user/register.html",
		Layout: "shared/layout.html",
		Data:   iris.Map{"Title": "User Index"},
	}
}
func (my *UserController) PostAdd() string {

	temp := repositories.NewAdmin(datasource.Instance())
	fmt.Println(temp)
	return "add"
}
