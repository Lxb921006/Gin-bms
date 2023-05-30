package model

import (
	"github.com/Lxb921006/Gin-bms/project/dao"
	"gorm.io/gorm"
	"time"
)

type AssetsProcessUpdateRecordModel struct {
	ID         int64     `form:"id,omitempty" json:"id" gorm:"primaryKey"`
	Ip         string    `form:"ip" json:"ip" gorm:"not null" binding:"required"`
	Uuid       string    `form:"uuid" json:"uuid" gorm:"not null;unique" binding:"required"`
	UpdateName string    `form:"update_name" json:"update_name" gorm:"not null" binding:"required"`
	Project    string    `form:"project" json:"project" gorm:"not null" binding:"required"`
	Operator   string    `form:"operator" json:"operator" gorm:"not null" binding:"required"`
	Progress   int32     `form:"progress,omitempty" json:"progress" gorm:"default:0;nullable"`
	CostTime   int32     `form:"cost_time,omitempty" json:"cost_time" gorm:"default:0;nullable"`
	Start      time.Time `form:"start,omitempty" json:"start" gorm:"default:CURRENT_TIMESTAMP;nullable"`
	End        time.Time `form:"end,omitempty" json:"end" gorm:"default:CURRENT_TIMESTAMP;nullable"`
}

func (pur *AssetsProcessUpdateRecordModel) Create(data AssetsProcessUpdateRecordModel) (err error) {
	if err = dao.DB.Create(&data).Error; err != nil {
		return
	}
	return
}

func (pur *AssetsProcessUpdateRecordModel) Del(pid []int32) (err error) {
	tx := dao.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Where("id IN ?", pid).Delete(pur).Error; err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error

}

func (pur *AssetsProcessUpdateRecordModel) BeforeSave(tx *gorm.DB) (err error) {
	if pur.Start.IsZero() {
		pur.Start = time.Now()
	}

	if pur.End.IsZero() {
		pur.End = time.Now()
	}

	return
}
