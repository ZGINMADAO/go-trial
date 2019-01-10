package admin

import (
	"go-trial/web/controllers"
	"go-trial/datamodels"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"fmt"
)

type SupplierController struct {
	controllers.BaseController
}

func (my *SupplierController) Get() mvc.View {
	return mvc.View{
		Name: "admin/supplier/list.html",
	}
}

func (my *SupplierController) GetList() {

	total, _ := my.DB.Count(new(datamodels.Supplier))
	var supplierList []datamodels.Supplier
	err := my.DB.Find(&supplierList)
	if err != nil {
		my.Ctx.Application().Logger().Warn(err.Error())
		return
	}
	my.Ctx.JSON(iris.Map{"rows": supplierList, "total": total})
}

const PostageFreeKey = "postage_free"
const PostageSetKey = "postage_set"

func (my *SupplierController) GetPostage() mvc.View {
	var postageFree, postSet datamodels.Setting

	postageFree.Key = PostageFreeKey
	my.DB.Get(&postageFree)

	postSet.Key = PostageSetKey
	my.DB.Get(&postSet)

	return mvc.View{
		Name: "admin/supplier/postage.html",
		Data: iris.Map{"PostageFree": postageFree.Val, "PostSet": postSet.Val},
	}
}

func (my *SupplierController) PutPostage() {

	var data = struct {
		PostageFree string
		PostageSet  string
	}{}

	my.Ctx.ReadForm(&data)

	var postageFree, postSet datamodels.Setting
	postageFree.Key = PostageFreeKey
	has, freeErr := my.DB.Get(&postageFree)
	if freeErr != nil {
		my.Ctx.Application().Logger().Warn(freeErr.Error())
		my.ReturnJson(false, "", freeErr.Error()+"123")
		return
	}
	temp := new(datamodels.Setting)
	temp.Val = data.PostageFree
	if has {
		_, updateErr := my.DB.Update(temp, datamodels.Setting{Key: PostageFreeKey})
		if updateErr != nil {
			fmt.Println(updateErr.Error())
			return
		}
	} else {
		temp.Key = PostageFreeKey
		my.DB.Insert(temp)
	}

	postSet.Key = PostageSetKey
	hasSet, setErr := my.DB.Get(&postSet)

	if setErr != nil {
		my.Ctx.Application().Logger().Warn(setErr.Error())
		my.ReturnJson(false, "", setErr.Error()+"123")
		return
	}

	tempSet := new(datamodels.Setting)
	tempSet.Val = data.PostageSet

	if hasSet {
		_, updateSetErr := my.DB.Update(tempSet, datamodels.Setting{Key: PostageSetKey})
		if updateSetErr != nil {
			fmt.Println(updateSetErr.Error())
			return
		}
	} else {
		tempSet.Key = PostageSetKey
		my.DB.Insert(tempSet)
	}
	my.ReturnJson(true, "", "操作成功")

}
