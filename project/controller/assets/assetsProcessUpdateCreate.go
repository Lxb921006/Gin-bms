package assets

import (
	"github.com/Lxb921006/Gin-bms/project/model"
	"github.com/gin-gonic/gin"
)

type ProcessUpdateForm struct {
}

func (c *ProcessUpdateForm) Create(ctx *gin.Context) (err error) {
	var cm model.AssetsProcessUpdateRecordModel
	if err = ctx.ShouldBindJSON(&cm); err != nil {
		return
	}

	if err = cm.Create(cm); err != nil {
		return
	}

	return
}
