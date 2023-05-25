package service

import (
	"encoding/json"
	"fmt"
	"github.com/Lxb921006/Gin-bms/project/command/client"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Ws struct {
	Conn        *websocket.Conn `json:"-"`
	Ip          string          `json:"ip"`
	ProcessName string          `json:"name"`
}

func (ws *Ws) Run() (err error) {
	_, message, err := ws.Conn.ReadMessage()

	if err != nil {
		log.Println(err)
		return
	}

	if err = ws.Parse(message); err != nil {
		return
	}

	if err = ws.Send(); err != nil {
		return err
	}

	return
}

func (ws *Ws) Send() (err error) {
	server := fmt.Sprintf("%s:12306", ws.Ip)
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	cn := client.NewRpcClient(ws.ProcessName, ws.Conn, conn)
	if err = cn.Send(); err != nil {
		return err
	}

	return
}

func (ws *Ws) Parse(data []byte) (err error) {
	if err = json.Unmarshal(data, &ws); err != nil {
		return
	}

	return
}

func NewWs(Conn *websocket.Conn) *Ws {
	return &Ws{
		Conn: Conn,
	}
}
