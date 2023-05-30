package assets

import (
	"fmt"
	"github.com/Lxb921006/Gin-bms/project/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ProcessWs(ctx *gin.Context) {
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err)
		return
	}

	defer conn.Close()

	ws := service.NewWs(conn)

	if err = ws.Run(); err != nil {
		if err = ws.Conn.WriteMessage(1, []byte(fmt.Sprintf("%s", err.Error()))); err != nil {
			return
		}
		return
	}
}

func GetMissionStatus(ctx *gin.Context) {
	var ps ProcessStatusForm
	data, err := ps.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": 10000,
	})
}

func AddAssets(ctx *gin.Context) {

}

func CreateUpdateProcess(ctx *gin.Context) {
	var create CreateProcessUpdateForm
	if err := create.Create(ctx); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10002,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"code":    10000,
	})
}
