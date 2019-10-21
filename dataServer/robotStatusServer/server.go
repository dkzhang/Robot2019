package main

import (
	"context"
	"log"
	"net"

	pb "Robot2019/dataServer/robotStatusServer/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50061"
)

type server struct {
	pb.UnimplementedRobotStatusServiceServer
}

func (s *server) GetRobotStatus(ctx context.Context, in *pb.RobotStatusRequest) (*pb.RobotStatusReply, error) {
	log.Printf("Received: %v", in.GetTag())
	//TODO: 连接redis容器，读取相关状态信息
	return &pb.RobotStatusReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRobotStatusServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
