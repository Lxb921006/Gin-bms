package utils

import (
	"fmt"
	"github.com/Lxb921006/Gin-bms/project/api"
	"github.com/Lxb921006/Gin-bms/project/command/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

type Celery struct {
	Works chan api.CeleryInterface
	Limit chan struct{}
	Wg    sync.WaitGroup
	Err   chan error
}

func NewCelery() *Celery {
	//data := celery.Data()

	c := &Celery{
		Works: make(chan api.CeleryInterface),
	}

	go func() {
		for w := range c.Works {
			data := w.Data()
			server := fmt.Sprintf("%s:12306", data["ip"].(string))
			conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return
			}

			cn := client.NewRpcClient(data["update_name"].(string), data["uuid"].(string), nil, conn)
			go func() {
				if err = cn.Send(); err != nil {
					return
				}
			}()
		}

	}()

	return c
}

func (c *Celery) Task(task api.CeleryInterface) {
	c.Works <- task
}
