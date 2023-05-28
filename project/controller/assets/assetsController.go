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

}
