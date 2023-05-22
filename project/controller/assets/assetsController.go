package assets

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
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

	messageType, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error during message reading:", err)
		return
	}
	log.Printf("Received: %s", message)

	Process(conn, messageType)

}

func Process(conn *websocket.Conn, messageType int) {
	count := 0
	for {
		count++
		text := fmt.Sprintf("The server has received websocket data %d!\n", count)
		err := conn.WriteMessage(messageType, []byte(text))
		if err != nil {
			log.Println("Error during message writing >>>", err)
			break
		}

		time.Sleep(time.Second / 5)
	}
}
