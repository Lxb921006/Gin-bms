package model

import (
	"github.com/Lxb921006/Gin-bms/project/dao"
	"gorm.io/gorm"
	"time"
)

type ProcessUpdateRecordModel struct {
	gorm.Model
	Ip         string    `json:"ip" gorm:"not null"`
	UpdateName string    `json:"update_name" gorm:"not null"`
	Project    string    `json:"project" gorm:"not null"`
	Status     string    `json:"status" gorm:"not null"`
	Operator   string    `json:"operator" gorm:"not null"`
	Progress   int32     `json:"progress" gorm:"not null"`
	CostTime   int32     `json:"cost_time" gorm:"not null"`
	Start      time.Time `json:"start" gorm:"-"`
	End        time.Time `json:"end" gorm:"-"`
}

func (pur *ProcessUpdateRecordModel) Add() (err error) {

	return
}

func (pur *ProcessUpdateRecordModel) Del(pid []int32) (err error) {
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
