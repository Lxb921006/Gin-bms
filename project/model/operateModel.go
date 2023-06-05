package model

import (
	"time"

	"github.com/Lxb921006/Gin-bms/project/dao"
	"github.com/Lxb921006/Gin-bms/project/service"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type OperateLog struct {
	gorm.Model
	Url      string    `json:"url" gorm:"not null"`
	Operator string    `json:"operator" gorm:"not null"`
	Ip       string    `json:"ip" gorm:"not null"`
	Start    time.Time `json:"start" gorm:"-"`
	End      time.Time `json:"end" gorm:"-"`
}

func (o *OperateLog) OperateLogList(page int, op OperateLog) (data *service.Paginate, err error) {
	var os []OperateLog
	sql := dao.DB.Model(o).Where(op)
	pg := service.NewPaginate()
	data, err = pg.GetPageData(page, sql)
	if err != nil {
		return
	}

	if err = data.Gd.Find(&os).Error; err != nil {
		return
	}

	data.ModelSlice = os

	return
}

func (o *OperateLog) OperateLogListByDate(page int, op OperateLog) (data *service.Paginate, err error) {
	var os []OperateLog
	sql := dao.DB.Model(o).Or(op).Where("created_at between ? and ?", op.Start, op.End)
	pg := service.NewPaginate()
	data, err = pg.GetPageData(page, sql)
	if err != nil {
		return
	}

	if err = data.Gd.Find(&os).Error; err != nil {
		return
	}

	data.ModelSlice = os

	return
}

func (o *OperateLog) AddOperateLog(ctx *gin.Context) (err error) {
	o.Url = ctx.Request.URL.Path
	o.Operator = ctx.Query("user")
	o.Ip = ctx.RemoteIP()

	if err = dao.DB.Create(o).Error; err != nil {
		return
	}
	return
}
