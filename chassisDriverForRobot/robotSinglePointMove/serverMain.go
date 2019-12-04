package main

import (
	MMoveServer "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/server"
	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	SMoveServer "Robot2019/chassisDriverForRobot/robotSinglePointMove/server"
	"Robot2019/myUtil"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":50071"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSinglePointMoveServer(s, &SMoveServer.Server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
