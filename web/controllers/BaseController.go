package controllers

import (
	"github.com/kataras/iris"
	"github.com/go-xorm/xorm"
)

type BaseController struct {
	Ctx iris.Context
	DB  *xorm.Engine
}
