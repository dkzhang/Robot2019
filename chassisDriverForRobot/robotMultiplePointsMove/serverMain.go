package main

import (
	pb "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/grpc"
	MMoveServer "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/server"
	"Robot2019/myUtil"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":50072"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMultiplePointsMoveServer(s, &MMoveServer.Server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
