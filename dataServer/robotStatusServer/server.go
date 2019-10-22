package main

import (
	"Robot2019/myUtil"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "Robot2019/dataServer/robotStatusServer/grpc"
	"google.golang.org/grpc"

	"github.com/gomodule/redigo/redis"
)

const (
	port = ":50061"
)

type server struct {
	pb.UnimplementedRobotStatusServiceServer
}

func (s *server) GetRobotStatus(ctx context.Context, in *pb.RobotStatusRequest) (*pb.RobotStatusReply, error) {
	log.Printf("Received: %v", in.GetTag())
	//连接redis容器，读取相关状态信息
	c, err := redis.Dial("tcp", "myRedis001:6379")
	if err != nil {
		return nil, fmt.Errorf("redis dial error: %v", err)
	}
	defer c.Close()

	result, err := redis.String(c.Do("GET", "CurrentRobotStatus"))
	if err != nil {
		return nil, fmt.Errorf("Get CurrentRobotStatus error: %v", err)
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
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRobotStatusServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
