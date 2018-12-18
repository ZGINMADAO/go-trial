package controllers

import (
	"fmt"
	"github.com/kataras/iris/mvc"
	"go-trial/services"
)

type OrderController struct {
	BaseController
}

func (my *OrderController) Get() mvc.View {
	return mvc.View{
		Name: "admin/order/lists.html",
	}
}

func (my *OrderController) GetList() {

	service := services.NewOrder()

	page := my.Ctx.URLParam("page")
	size := my.Ctx.URLParam("size")
	keyWord := my.Ctx.URLParam("keyWord")
	var requestData = make(map[string]string)
	requestData["KeyWord"] = keyWord
	requestData["Size"] = size
	requestData["Page"] = page
	total := service.OrderList(my.DB, requestData)
	fmt.Println(total)
	//my.Ctx.JSON(iris.Map{"rows": results, "total": total})
}
