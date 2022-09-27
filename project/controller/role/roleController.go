package role

import (
	"fmt"
	"net/http"

	"github.com/Lxb921006/Gin-bms/project/model"

	"github.com/Lxb921006/Gin-bms/project/logic/role"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CreateRoleForm struct {
	RoleName string `form:"rolename" binding:"required"`
}

type DeleteRoleJson struct {
	Rid []uint `form:"rid" binding:"required"`
}

type RoleListQuery struct {
	RoleName string `form:"rolename"`
	Page     int    `form:"page" validate:"min=1" binding:"required"`
}

type OperatePermsJson struct {
	Rid      uint   `form:"rid" binding:"required"`
	Pid      []uint `form:"pid" binding:"required"`
	RoleName string `form:"rolename" binding:"required"`
}

type UserPermsQuery struct {
	Uid uint `form:"uid" binding:"required"`
}

type RolePermsQuery struct {
	Rid uint `form:"rid" binding:"required"`
}

func CreateRole(ctx *gin.Context) {
	var r model.Role
	var cr CreateRoleForm

	if err := ctx.ShouldBindWith(&cr, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	r.RoleName = cr.RoleName

	if err := r.CreateRole(r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("[%s] 创建失败, errMsg: %s", cr.RoleName, err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("[%s] 创建成功", cr.RoleName),
	})
}

func DeleteRoles(ctx *gin.Context) {
	var r model.Role
	var dr DeleteRoleJson

	if err := ctx.ShouldBindJSON(&dr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := r.DeleteRole(dr.Rid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v删除成功", dr.Rid),
	})
}

func GetRolesInfo(ctx *gin.Context) {
	var r model.Role
	data, err := r.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func AllotPermsToRole(ctx *gin.Context) {
	var ap OperatePermsJson

	if err := ctx.ShouldBindJSON(&ap); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := role.UpdateUserPerms(ap.Pid, ap.Rid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("[%s]添加权限pid为:%v的成功", ap.RoleName, ap.Pid),
	})
}

func RemoveRolePerms(ctx *gin.Context) {
	var r model.Role
	var rp OperatePermsJson

	if err := ctx.ShouldBindJSON(&rp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := r.RemovePerms(rp.Rid, rp.Pid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	mp := r.FormatUserPerms(data, 0)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("[%s]移除权限pid为:%v成功", rp.RoleName, rp.Pid),
		"data":    mp,
	})
}

func GetRolesList(ctx *gin.Context) {
	var r model.Role
	var rp RoleListQuery
	if err := ctx.ShouldBindQuery(&rp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()
	vd := NewValidateData(validate)
	vd.ValidateStruct(rp)

	r.RoleName = rp.RoleName

	data, err := r.GetRolesList(rp.Page, r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":     data.ModelSlice,
		"total":    data.Total,
		"pageSize": data.PageSize,
	})
}

func GetUserPerms(ctx *gin.Context) {
	var r model.Role
	var up UserPermsQuery

	if err := ctx.ShouldBindQuery(&up); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := r.GetUserPerms(up.Uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var fdata []role.Menu

	if len(data) != 0 {
		fdata = role.FormatUserPerms(data, 0)

	} else {
		fdata = []role.Menu{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": fdata,
	})
}

func GetRolePerms(ctx *gin.Context) {
	var r model.Role
	var up RolePermsQuery
	var pidList []uint

	if err := ctx.ShouldBindQuery(&up); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := r.GetRolePerms(up.Rid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var fdata []role.Menu

	if len(data) != 0 {
		fdata = role.FormatUserPerms(data, 0)
		for _, v := range data {
			pidList = append(pidList, v.ID)
		}
	} else {
		fdata = []role.Menu{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    fdata,
		"pidList": pidList,
	})
}

func GetAllFormatPerms(ctx *gin.Context) {
	var r model.Role
	data, err := r.GetAllFormatPerms()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var fdata []role.Menu

	if len(data) != 0 {
		fdata = role.FormatUserPerms(data, 0)

	} else {
		fdata = []role.Menu{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": fdata,
	})
}
