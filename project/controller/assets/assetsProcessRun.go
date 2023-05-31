package assets

import (
	"fmt"
	"github.com/Lxb921006/Gin-bms/project/command/client"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AssetsProcessRunForm struct {
	Ip         string     `form:"ip" json:"ip" gorm:"not null" binding:"required"`
	UpdateName string     `form:"update_name" json:"update_name" gorm:"not null" binding:"required"`
	Uuid       string     `form:"uuid" json:"uuid" gorm:"not null;unique" binding:"required"`
	Err        chan error `form:"err,omitempty" json:"-"`
}

func (apf *AssetsProcessRunForm) Run(ctx *gin.Context) (err error) {
	if err = ctx.ShouldBind(apf); err != nil {
		return
	}

	go func() {
		server := fmt.Sprintf("%s:12306", apf.Ip)
		conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			apf.Err <- err
			return
		} else {
			apf.Err <- nil
		}

		cn := client.NewRpcClient(apf.UpdateName, apf.Uuid, nil, conn)
		if err = cn.Send(); err != nil {
			apf.Err <- err
		} else {
			apf.Err <- nil
		}
	}()

	for {
		select {
		case err = <-apf.Err:
			return
		default:
		}
	}
}

func NewAssetsProcessRunForm() *AssetsProcessRunForm {
	return &AssetsProcessRunForm{
		Err: make(chan error, 1),
	}
}
