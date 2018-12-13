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

	build.Limit(SIZE, (page-1)*SIZE).Desc("Id").Find(&product)
	total, _ := (&countBuild).Count(new(datamodels.Product))

	s1 := make([]datamodels.Product, 0)
	if product == nil {
		product = s1
	}

	fmt.Println(product)
	my.Ctx.JSON(iris.Map{"rows": product, "total": total})
}

func (my *ProductController) PostType() {

	title := my.Ctx.PostValueTrim("title")
	image := my.Ctx.PostValueTrim("image")

	fmt.Println(title, image)
	if len(title) < 1 || len(image) < 1 {
		my.Ctx.JSON(apiResource(false, nil, "名称和图片不能为空"))
		return
	}
	var product datamodels.Product
	product.Title = title
	product.Image = image
	my.DB.Insert(&product)
	fmt.Println(product)
	my.Ctx.JSON(apiResource(true, nil, "操作成功"))
}

func (my *ProductController) PutTypeBy(id int) {

	title := my.Ctx.PostValueTrim("title")
	image := my.Ctx.PostValueTrim("image")

	if len(title) < 1 || len(image) < 1 {
		my.Ctx.JSON(apiResource(false, nil, "名称和图片不能为空"))
		return
	}
	product := new(datamodels.Product)
	product.Title = title
	product.Image = image
	affected, _ := my.DB.Id(id).Update(product)
	if affected < 1 {
		my.Ctx.JSON(apiResource(false, nil, "删除失败"))
		return
	}
	my.Ctx.JSON(apiResource(true, nil, "操作成功"))
}

func (my *ProductController) DeleteTypeBy(id int) {
	product := new(datamodels.Product)
	affected, _ := my.DB.Id(id).Delete(product)
	if affected < 1 {
		my.Ctx.JSON(apiResource(false, nil, "删除失败"))
		return
	}
	my.Ctx.JSON(apiResource(true, nil, "操作成功"))
}
