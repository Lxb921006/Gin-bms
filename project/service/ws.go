package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

type ProcessData struct {
	Ip          string `json:"ip"`
	ProcessName string `json:"name"`
}

type Ws struct {
	Conn *websocket.Conn
}

func (ws *Ws) Receive() (err error) {
	_, message, err := ws.Conn.ReadMessage()
	if err != nil {
		return
	}

	_, err = ws.Parse(message)
	if err != nil {
		return
	}

	return
}

func (ws *Ws) Send(pd ProcessData) (err error) {
	return
}

func (ws *Ws) Parse(data []byte) (pd ProcessData, err error) {
	if err = json.Unmarshal(data, &pd); err != nil {
		return
	}
	return
}
