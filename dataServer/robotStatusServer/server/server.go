package server

import (
	"Robot2019/cache"
	"Robot2019/myUtil"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "Robot2019/dataServer/robotStatusServer/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50071"
)

type server struct {
	pb.UnimplementedRobotStatusServiceServer
}

func (s *server) GetRobotStatus(ctx context.Context, in *pb.RobotStatusRequest) (*pb.RobotStatusReply, error) {
	log.Printf("Received: %v", in.GetTag())

	opts := &cache.RedisOpts{
		Host: cache.RedisHost,
	}
	theRedis := cache.NewRedis(opts)

	result, err := theRedis.ListIndex("RobotStatus", -1)
	if err != nil {
		return nil, fmt.Errorf("theRedis.ListIndex: %v", err)
	}

	theReply := pb.RobotStatusReply{}
	err = json.Unmarshal([]byte(result), &theReply)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %v", err)
	} else {
		return &theReply, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRobotStatusServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
