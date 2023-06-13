package utils

import (
	"github.com/Lxb921006/Gin-bms/project/api"
	"github.com/gin-gonic/gin"
	"sync"
)

type Celery struct {
	Works chan api.CeleryInterface
	Limit chan struct{}
	Wg    sync.WaitGroup
}

func NewCelery(workers int64) *Celery {
	c := &Celery{
		Works: make(chan api.CeleryInterface),
		Limit: make(chan struct{}, workers),
	}

	go func() {

	}()

	return c
}

func (c *Celery) Run(ctx *gin.Context) (err error) {

	return
}
