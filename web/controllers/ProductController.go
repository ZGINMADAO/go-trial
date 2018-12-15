package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"go-trial/datamodels"
	"fmt"
	"go-trial/units"
	"go-trial/services"
)

type ProductController struct {
	BaseController
}

func (my *ProductController) Get() mvc.View {
	return mvc.View{
		Name: "admin/product/lists.html",
	}
}

func (my *ProductController) GetList() {

	service := services.NewProduct()

	page := my.Ctx.URLParam("page")
	size := my.Ctx.URLParam("size")
	keyWord := my.Ctx.URLParam("keyWord")
	var requestData = make(map[string]string)
	requestData["KeyWord"] = keyWord
	requestData["Size"] = size
	requestData["Page"] = page
	total, results := service.ProductWithType(my.DB, requestData)
	my.Ctx.JSON(iris.Map{"rows": results, "total": total})
}

func (my *ProductController) GetEdit() mvc.View {

	var productTypes []datamodels.ProductType

	my.DB.Find(&productTypes)

	return mvc.View{
		Name: "admin/product/edit.html",
		Data: map[string][]datamodels.ProductType{"ProductTypes": productTypes},
	}
}

func (my *ProductController) Post() {
	service := services.NewProduct()
	result := service.Request(my.Ctx)
	my.DB.Insert(result)
}

func (my *ProductController) Put() {

}

func (my *ProductController) GetType() mvc.View {
	return mvc.View{
		Name: "admin/product/type-lists.html",
	}
}

func (my *ProductController) GetTypeList() {
	page, _ := my.Ctx.URLParamInt("page")
	size, _ := my.Ctx.URLParamInt("pageSize")
	keyWord := my.Ctx.URLParam("keyWord")

	if page < 1 {
		page = 1
	}

	var productType []datamodels.ProductType
	build := my.DB.Where("")
	if len(keyWord) > 0 {
		fmt.Println(keyWord)
		build = my.DB.Where("title like ?", "%"+keyWord+"%")
	}

	countBuild := *build

	build.Limit(size, (page-1)*size).Desc("Id").Find(&productType)
	total, _ := (&countBuild).Count(new(datamodels.ProductType))

	type result struct {
		datamodels.Product
		CreatedAt string
	}
	results := make([]result, len(productType))
	for inx, val := range productType {
		results[inx].CreatedAt = val.CreatedAt.Format(units.CnFormat)
		results[inx].Id = val.Id
		results[inx].Title = val.Title
		results[inx].Image = val.Image
	}
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
	var productType datamodels.ProductType
	productType.Title = title
	productType.Image = image
	my.DB.Insert(&productType)
	my.Ctx.JSON(apiResource(true, nil, "操作成功"))
}

func (my *ProductController) PutTypeBy(id int) {

	title := my.Ctx.PostValueTrim("title")
	image := my.Ctx.PostValueTrim("image")

	if len(title) < 1 || len(image) < 1 {
		my.Ctx.JSON(apiResource(false, nil, "名称和图片不能为空"))
		return
	}
	productType := new(datamodels.ProductType)
	productType.Title = title
	productType.Image = image
	affected, _ := my.DB.Id(id).Update(productType)
	if affected < 1 {
		my.Ctx.JSON(apiResource(false, nil, "删除失败"))
		return
	}
	my.Ctx.JSON(apiResource(true, nil, "操作成功"))
}

func (my *ProductController) DeleteTypeBy(id int) {
	productType := new(datamodels.ProductType)
	affected, _ := my.DB.Id(id).Delete(productType)
	if affected < 1 {
		my.Ctx.JSON(apiResource(false, nil, "删除失败"))
		return
	}
	my.Ctx.JSON(apiResource(true, nil, "操作成功"))
}
