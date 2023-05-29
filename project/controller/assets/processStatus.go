package assets

import (
	"github.com/Lxb921006/Gin-bms/project/dao"
	"github.com/gin-gonic/gin"
)

type ProcessStatusForm struct {
	Result string `form:"result" binding:"required"`
}

func (ps *ProcessStatusForm) Get(ctx *gin.Context) (data map[string]string, err error) {
	if err = ctx.ShouldBind(ps); err != nil {
		return
	}

	data, err = dao.Rds.GetProcessStatus()
	//log.Println(data)
	if err != nil {
		return
	}

	return
}
