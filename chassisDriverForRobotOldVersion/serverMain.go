package main

import (
	robotMove "Robot2019/chassisDriverForRobotOldVersion/robotMove/grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SinglePointMove(ctx context.Context, in *robotMove.SinglePointInfo) (*robotMove.MoveResponse, error) {
	log.Printf("Received: %v", in)
	return &robotMove.MoveResponse{
		Command: "the command",
		Uuid:    10001,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	robotMove.RegisterRobotMoveServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
