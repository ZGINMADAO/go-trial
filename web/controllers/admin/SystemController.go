package admin

import (
	"github.com/kataras/iris/mvc"
	"fmt"
	"go-trial/units"
	"github.com/kataras/iris"
	"go-trial/datamodels"
	"strconv"
	"go-trial/web/controllers"
)

type SystemController struct {
	controllers.BaseController
}

func (my *SystemController) GetRole() mvc.View {

	var permissions []datamodels.Tree
	my.DB.Find(&permissions)
	deepResult := make([]units.List, 0)
	units.Recursive(permissions, 0, &deepResult)

	//var role datamodels.Role
	//ok, err := my.DB.Get(&role)
	//if !ok || err != nil {
	//
	//}

	//var rolePermissions []datamodels.RolePermission
	//my.DB.Where("role_id=?", Id).Select("permission_id").Find(&rolePermissions)

	return mvc.View{
		Name: "admin/system/role_lists.html",
		Data: struct {
			Permissions interface{}
		}{deepResult},
	}
}
func (my *SystemController) GetRoleByPermissions(id int) {
	var rolePermissions []datamodels.RolePermission
	my.DB.Where("role_id=?", id).Select("permission_id").Find(&rolePermissions)
	my.ReturnJson(true, rolePermissions, "")
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

func (my *SystemController) GetPermissionList() {
	//var permissions []datamodels.Permission

}

func (my *SystemController) PutRoleByPermissions(roleId int) {

	idList := my.Ctx.PostValues("idList[]")
	idIntList := make([]int, len(idList))
	for key, val := range idList {
		temp, _ := strconv.Atoi(val)
		idIntList[key] = temp
	}

	var role datamodels.Role
	fmt.Println(roleId)
	ok, err := my.DB.Where("id=?",roleId).Get(&role)
	fmt.Println(role)
	if err != nil || !ok {
		my.ReturnJson(false, nil, err.Error())
		return
	}

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

	for _, val := range idIntList {
		isIn := units.IntInArray(val, permissionIds)
		if !isIn {
			addData = append(addData, datamodels.RolePermission{RoleId: roleId, PermissionId: val})
		}
	}
	my.DB.Insert(&addData)

	deleteList := make([]int, 0)
	for _, val := range permissionIds {
		isIn := units.IntInArray(val, idIntList)
		if !isIn {
			deleteList = append(deleteList, val)
		}
	}

	deleteData := new(datamodels.RolePermission)
	my.DB.Where("role_id=?", roleId).In("permission_id", deleteList).Delete(deleteData)
	fmt.Println(deleteData)
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
