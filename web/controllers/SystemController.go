package controllers

import (
	"github.com/kataras/iris/mvc"
	"fmt"
	"go-trial/units"
	"github.com/kataras/iris"
	"go-trial/datamodels"
	//"strconv"
	"strconv"
)

type SystemController struct {
	BaseController
}

func (my *SystemController) GetRole() mvc.View {
	return mvc.View{
		Name: "admin/system/order_lists.html",
	}
}

func (my *SystemController) GetRoleList() {
	page, _ := my.Ctx.URLParamInt("page")
	size, _ := my.Ctx.URLParamInt("pageSize")
	keyWord := my.Ctx.URLParam("keyWord")

	if page < 1 {
		page = 1
	}

	var roles []datamodels.Role
	build := my.DB.Where("")
	if len(keyWord) > 0 {
		fmt.Println(keyWord)
		build = my.DB.Where("title like ?", "%"+keyWord+"%")
	}

	countBuild := *build

	build.Limit(size, (page-1)*size).Desc("Id").Find(&roles)
	total, _ := (&countBuild).Count(new(datamodels.ProductType))

	type result struct {
		datamodels.Role
		CreatedAt string
	}
	results := make([]result, len(roles))
	for inx, val := range roles {
		results[inx].CreatedAt = val.CreatedAt.Format(units.CnFormat)
		results[inx].Id = val.Id
		results[inx].Title = val.Title
	}
	my.Ctx.JSON(iris.Map{"rows": results, "total": total})
}

func (my *SystemController) GetPermissions() {
	var permissions []datamodels.Tree
	my.DB.Find(&permissions)
	my.ReturnJson(true, permissions, "")
}

func (my *SystemController) PutRoleByPermissions(roleId int) {

	idList := my.Ctx.PostValues("idList")

	//var role datamodels.Role
	//ok, err := my.DB.Id(roleId).Get(&role)
	//if err != nil || !ok {
	//	my.ReturnJson(false, nil, "角色信息不存在")
	//	return
	//}

	var permissions []datamodels.RolePermission
	pErr := my.DB.Where("role_id=?", roleId).Select("permission_id").Find(&permissions)
	if pErr != nil {
		my.ReturnJson(false, nil, pErr.Error())
		return
	}

	permissionIds := make([]int, len(permissions))
	for key, val := range permissions {
		permissionIds[key] = val.PermissionId
	}

	addData := make([]datamodels.RolePermission, 0)

	for _, val := range idList {
		isIn, _ := units.InArray(val, permissionIds)
		if !isIn {
			temp, _ := strconv.Atoi(val)
			addData = append(addData, datamodels.RolePermission{RoleId: roleId, PermissionId: temp})
		}
	}
	my.DB.Insert(&addData)

	deleteList := make([]int, 0)
	for _, val := range permissionIds {
		isIn, _ := units.InArray(val, idList)
		if !isIn {
			deleteList = append(deleteList, val)
		}
	}

	//var deleteData []datamodels.RolePermission
	//my.DB.Where("role_id=?", roleId).In("permission_id", deleteList).Delete(&deleteData)

	my.ReturnJson(true, nil, "操作成功")
}

func (my *SystemController) PutAdminRole() {
	//idList := my.Ctx.PostValues("idList")
	//
	//idIntList := make([]int, len(idList))
	//for key, val := range idList {
	//	v, err := strconv.Atoi(val)
	//	if err != nil {
	//		v = 0
	//	}
	//	idIntList[key] = v
	//}
	//addData := make([]datamodels.RolePermission, len(idList))
	//my.DB.Insert(addData)
	//my.ReturnJson(true, nil, "操作成功")
}
