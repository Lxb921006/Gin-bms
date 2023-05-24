package client

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/Lxb921006/Gin-bms/project/command/command"
	"google.golang.org/grpc"
	"io"
)
import "github.com/gorilla/websocket"

type RpcClient struct {
	Name    string
	RpcConn *grpc.ClientConn
	WsConn  *websocket.Conn
}

func (rc *RpcClient) Send() (err error) {
	switch rc.Name {
	case "dockerUpdate":
		if err = rc.DockerUpdate(); err != nil {
			return err
		}
	case "javaUpdate":
		if err = rc.DockerUpdate(); err != nil {
			return err
		}
	default:
		return errors.New("无效操作")
	}

	return
}

func (rc *RpcClient) DockerUpdate() (err error) {
	c := pb.NewStreamUpdateProcessServiceClient(rc.RpcConn)
	stream, err := c.DockerUpdate(context.Background(), &pb.StreamRequest{})
	if err != nil {
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err = rc.WsConn.WriteMessage(1, []byte(fmt.Sprintf("%s\n", resp.Message))); err != nil {
			return err
		}
	}

	return
}

func NewRpcClient(name string, ws *websocket.Conn, rc *grpc.ClientConn) *RpcClient {
	return &RpcClient{
		Name:    name,
		WsConn:  ws,
		RpcConn: rc,
	}
}
