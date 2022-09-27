package perms

import (
	"fmt"
	"net/http"

	"github.com/Lxb921006/Gin-bms/project/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CreatePermMenuForm struct {
	Path     string `form:"path" binding:"required"`
	Title    string `form:"title" binding:"required"`
	ParentId uint   `form:"parentid"`
	Level    uint   `form:"level" binding:"required"`
}

type DeletePermsJson struct {
	Pid []uint `form:"pid" binding:"required"`
}

type PermsListQuery struct {
	Page int `form:"page" validate:"min=1" binding:"required"`
}

func CreatePermMenu(ctx *gin.Context) {
	var p model.Permission
	var pd CreatePermMenuForm

	if err := ctx.ShouldBindWith(&pd, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	p.Path = pd.Path
	p.Title = pd.Title
	p.ParentId = pd.ParentId
	p.Level = pd.Level

	if err := p.CreatePerms(p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("[%s] 创建失败, errMsg: %s", pd.Title, err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("[%s] 创建成功", pd.Title),
	})

}

func DeletePermsMenu(ctx *gin.Context) {
	var p model.Permission
	var dp DeletePermsJson

	if err := ctx.ShouldBindJSON(&dp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := p.DeletePerms(dp.Pid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v删除成功", dp.Pid),
	})

}

func GetPermsList(ctx *gin.Context) {
	var p model.Permission
	var pp PermsListQuery
	if err := ctx.ShouldBindQuery(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()
	vd := NewValidateData(validate)
	vd.ValidateStruct(pp)

	data, err := p.GetPermsList(pp.Page)
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
