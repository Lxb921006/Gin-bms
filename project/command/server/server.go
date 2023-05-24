package main

import (
	"bufio"
	"fmt"
	pb "github.com/Lxb921006/Gin-bms/project/command/command"
	"google.golang.org/grpc"
	"log"
	"net"
	"os/exec"
)

type server struct {
	pb.UnimplementedStreamUpdateProcessServiceServer
}

func (s *server) DockerUpdate(req *pb.StreamRequest, stream pb.StreamUpdateProcessService_DockerUpdateServer) (err error) {

	log.Println("rev run DockerUpdate")

	cmd := exec.Command("sh", "/root/shellscript/test.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		if err = stream.Send(&pb.StreamReply{Message: scanner.Text()}); err != nil {
			return
		}
	}

	if err = cmd.Wait(); err != nil {
		return
	}

	return
}

func (s *server) JavaUpdate(req *pb.StreamRequest, stream pb.StreamUpdateProcessService_JavaUpdateServer) (err error) {
	log.Println("rev run JavaUpdate")

	cmd := exec.Command("sh", "/root/shellscript/test2.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		if err = stream.Send(&pb.StreamReply{Message: scanner.Text()}); err != nil {
			return
		}
	}

	if err = cmd.Wait(); err != nil {
		return
	}

	return
}

func (s *server) DockerReload(req *pb.StreamRequest, stream pb.StreamUpdateProcessService_DockerReloadServer) (err error) {
	return
}

func (s *server) JavaReload(req *pb.StreamRequest, stream pb.StreamUpdateProcessService_JavaReloadServer) (err error) {
	return
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 12306))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStreamUpdateProcessServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
