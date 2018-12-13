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
	keyWord := my.Ctx.URLParam("keyWord")
	if page < 1 {
		page = 1
	}

	var product []datamodels.Product
	build := my.DB.Where("")
	if len(keyWord) > 0 {
		fmt.Println(keyWord)
		build = my.DB.Where("title like ?", "%"+keyWord+"%")
	}

	countBuild := *build

	build.Limit(SIZE, (page-1)*SIZE).Find(&product)
	total, _ := (&countBuild).Count(new(datamodels.Product))

	s1 := make([]datamodels.Product, 0)
	if product == nil {
		product = s1
	}

	fmt.Println(product)
	my.Ctx.JSON(iris.Map{"rows": product, "total": total})
}
