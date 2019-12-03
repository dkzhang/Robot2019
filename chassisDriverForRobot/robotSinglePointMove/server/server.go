package main

import (
	"Robot2019/myUtil"
	"fmt"
	"log"
	"net"
	//"os"
	//"os/exec"
	"time"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSinglePointMoveServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSinglePointMoveServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
