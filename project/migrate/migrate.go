package migrate

import (
	"github.com/Lxb921006/Gin-bms/project/dao"
	"github.com/Lxb921006/Gin-bms/project/model"
)

func InitTable() {
	dao.DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.OperateLog{},
	)
}
