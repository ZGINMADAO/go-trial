package controllers

import "github.com/kataras/iris"

type UserController struct {
	Ctx iris.Context
}

func (user *UserController) postAdd() string {
	return "success"
}
