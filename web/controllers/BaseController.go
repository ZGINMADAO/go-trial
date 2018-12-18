package controllers

import (
	"github.com/kataras/iris"
	"github.com/go-xorm/xorm"
)

type BaseController struct {
	Ctx iris.Context
	DB  *xorm.Engine
}

func (my *BaseController) ReturnJson(status bool, data interface{}, message string) {
	type apiJson struct {
		Status  bool        `json:"status"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data"`
	}
	my.Ctx.JSON(&apiJson{Status: status, Data: data, Message: message})
}
