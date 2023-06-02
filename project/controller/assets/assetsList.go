package assets

import (
	"github.com/Lxb921006/Gin-bms/project/model"
	"github.com/Lxb921006/Gin-bms/project/service"
	"github.com/Lxb921006/Gin-bms/project/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//资产列表的增删改查

type AssetsListForm struct {
	Ip      string `form:"ip,omitempty" json:"ip" gorm:"not null"`
	Project string `form:"project,omitempty" json:"project" gorm:"not null"`
	Page    int    `form:"page" validate:"min=1" binding:"required"`
}

func (a *AssetsListForm) List(ctx *gin.Context) (data *service.Paginate, err error) {
	var al model.AssetsModel
	if err = ctx.ShouldBind(a); err != nil {
		return
	}

	validate := validator.New()
	vd := NewValidateData(validate)
	if err = vd.ValidateStruct(a); err != nil {
		return
	}

	if err = utils.CopyStruct(a, &al); err != nil {
		return
	}

	data, err = al.List(a.Page, al)
	if err != nil {
		return
	}

	return
}
