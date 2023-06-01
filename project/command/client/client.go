package client

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/Lxb921006/Gin-bms/project/command/command"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"io"
	"log"
)

type RpcClient struct {
	Name    string
	Uuid    string
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
		if err = rc.JavaUpdate(); err != nil {
			return err
		}
	default:
		return errors.New("无效操作")
	}

	return
}

func (rc *RpcClient) DockerUpdate() (err error) {
	c := pb.NewStreamUpdateProcessServiceClient(rc.RpcConn)
	stream, err := c.DockerUpdate(context.Background(), &pb.StreamRequest{Uuid: rc.Uuid})
	if err != nil {
		return
	}

	defer rc.RpcConn.Close()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		log.Println(resp.Message)

		if rc.WsConn != nil {
			if err = rc.WsConn.WriteMessage(1, []byte(fmt.Sprintf("%s\n", resp.Message))); err != nil {
				return err
			}
		}
	}

	return
}

func (rc *RpcClient) JavaUpdate() (err error) {
	c := pb.NewStreamUpdateProcessServiceClient(rc.RpcConn)
	stream, err := c.JavaUpdate(context.Background(), &pb.StreamRequest{})
	if err != nil {
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if rc.WsConn != nil {
			if err = rc.WsConn.WriteMessage(1, []byte(fmt.Sprintf("%s\n", resp.Message))); err != nil {
				return err
			}
		}

	}

	return
}

func NewRpcClient(name, uuid string, ws *websocket.Conn, rc *grpc.ClientConn) *RpcClient {
	return &RpcClient{
		Name:    name,
		Uuid:    uuid,
		WsConn:  ws,
		RpcConn: rc,
	}
}
