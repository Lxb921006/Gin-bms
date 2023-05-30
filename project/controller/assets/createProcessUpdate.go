package assets

import (
	"github.com/Lxb921006/Gin-bms/project/model"
	"github.com/gin-gonic/gin"
)

type CreateProcessUpdateForm struct {
	//Ip         string   `form:"ip" json:"ip"`
	//Uuid       string   `form:"uuid" json:"uuid"`
	//Project    string   `form:"project" json:"project"`
	//Operator   string   `form:"operator" json:"operator"`
	//UpdateName string   `form:"update_name" json:"update_name"`
}

func (c *CreateProcessUpdateForm) Create(ctx *gin.Context) (err error) {
	var cm model.AssetsProcessUpdateRecordModel
	if err = ctx.ShouldBindJSON(&cm); err != nil {
		return
	}

	if err = cm.Create(cm); err != nil {
		return
	}

	return
}
