package model

import (
	"github.com/Lxb921006/Gin-bms/project/dao"
	"github.com/Lxb921006/Gin-bms/project/service"
	"gorm.io/gorm"
	"time"
)

type AssetsModel struct {
	gorm.Model
	Ip       string    `json:"ip" gorm:"not null"`
	Project  string    `json:"project" gorm:"not null"`
	Status   string    `json:"status" gorm:"not null"`
	Operator string    `json:"operator" gorm:"not null"`
	Start    time.Time `json:"start" gorm:"-"`
	End      time.Time `json:"end" gorm:"-"`
}

func (o *AssetsModel) List(page int, am AssetsModel) (data *service.Paginate, err error) {
	var os []AssetsModel
	sql := dao.DB.Model(os).Or(am)
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

func (o *AssetsModel) Del(pid []int32) (err error) {
	tx := dao.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Where("id IN ?", pid).Delete(o).Error; err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}
