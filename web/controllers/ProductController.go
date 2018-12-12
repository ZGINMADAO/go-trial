package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"go-trial/datamodels"
	"fmt"
)

type ProductController struct {
	BaseController
}

func (my *ProductController) GetType() mvc.View {
	return mvc.View{
		Name: "admin/product/type-lists.html",
		Data: iris.Map{"url": "/admin/login"},
	}
}

func (my *ProductController) GetTypeList() {

	const SIZE = 2
	page, _ := my.Ctx.URLParamInt("page")
	if page < 1 {
		page = 1
	}
	var product []datamodels.Product
	my.DB.Limit(SIZE, (page-1)*SIZE).Find(&product)
	fmt.Println(product)
	my.Ctx.JSON(apiResource(false, product, "登录失败"))
}
