package main

import (
	"log"

	"github.com/Lxb921006/Gin-bms/project/config"
	"github.com/Lxb921006/Gin-bms/project/dao"
	"github.com/Lxb921006/Gin-bms/project/migrate"
	"github.com/Lxb921006/Gin-bms/project/router/root"
)

func main() {
	//初始化mysql
	err := dao.InitPoolMysql()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	//初始化redis连接池
	dao.InitPoolRds(config.RedisConAddre, config.RedisPwd, config.RedisUserDb)
	if dao.RdPool == nil {
		log.Fatalf(dao.ErrorRedisConnectFailed.Error())
		return
	}
	dao.Rds = dao.NewRedisDb(dao.RdPool, map[string]dao.Md{})

	//这里初始化数据库表
	migrate.InitTable()

	//初始化gin并启动
	t := root.SetupRouter()
	err = t.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}
