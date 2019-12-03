package main

import (
	"Robot2019/myUtil"
	"fmt"
	"log"
	"net"
	//"os"
	//"os/exec"
	"time"

	pb "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50072"
)

type server struct {
	pb.UnimplementedMultiplePointsMoveServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMultiplePointsMoveServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
