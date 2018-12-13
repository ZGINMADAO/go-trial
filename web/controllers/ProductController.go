package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"go-trial/datamodels"
	"fmt"
	"go-trial/units"
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
	page, _ := my.Ctx.URLParamInt("page")
	size, _ := my.Ctx.URLParamInt("pageSize")
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

	build.Limit(size, (page-1)*size).Desc("Id").Find(&product)
	total, _ := (&countBuild).Count(new(datamodels.Product))

	type result struct {
		datamodels.Product
		CreatedAt string
	}
	results := make([]result, len(product))
	for inx, val := range product {
		results[inx].CreatedAt = val.CreatedAt.Format(units.CnFormat)
		results[inx].Id = val.Id
		results[inx].Title = val.Title
		results[inx].Image = val.Image
	}


	fmt.Println(product)
	my.Ctx.JSON(iris.Map{"rows": results, "total": total})
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
