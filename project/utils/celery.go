package utils

import (
	"fmt"
	"github.com/Lxb921006/Gin-bms/project/api"
	"github.com/Lxb921006/Gin-bms/project/command/client"
	"github.com/Lxb921006/Gin-bms/project/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Celery struct {
	Works chan api.CeleryInterface
}

func NewCelery() *Celery {
	var aprm model.AssetsProcessUpdateRecordModel
	var dataModel = make(map[string]interface{})

	c := &Celery{
		Works: make(chan api.CeleryInterface),
	}

	l := NewLogger("C:\\Users\\Administrator\\Desktop\\celery.log")
	wf := l.WriteLog()
	log.SetOutput(wf)

	go func() {
		for w := range c.Works {
			data, err := w.Data()
			dataModel["uuid"] = data["uuid"].(string)
			dataModel["status"] = 400

			if err != nil {
				log.Println("获取grpc参数失败: ", err)
				return
			}

			server := fmt.Sprintf("%s:12306", data["ip"].(string))
			conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				if err = aprm.Update(dataModel); err != nil {
					log.Println("更新失败-connect: ", err)
				}

				log.Println("连接grpc失败: ", err)

				return
			}

			cn := client.NewRpcClient(data["update_name"].(string), data["uuid"].(string), nil, conn)
			go func() {
				if err = cn.Send(); err != nil {
					if err = aprm.Update(dataModel); err != nil {
						log.Println("更新失败-send: ", err)
					}
					log.Println("grpc发送数据失败: ", err)
				}
			}()
		}

	}()

	return c
}

func (c *Celery) Task(task api.CeleryInterface) {
	c.Works <- task
}
