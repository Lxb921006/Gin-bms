package migrate

import (
	"github.com/Lxb921006/Gin-bms/project/dao"
	"github.com/Lxb921006/Gin-bms/project/model"
	"log"
)

func InitTable() {
	err := dao.DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.OperateLog{},
	)

	if err != nil {
		log.Println("InitTable esg >>>", err)
	}
}
