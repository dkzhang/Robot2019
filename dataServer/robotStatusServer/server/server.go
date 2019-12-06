package main

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

	robotStatusWriter "Robot2019/chassisDriverForRobot/subscribeRobotStatusWriter/server"
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

	robotStatus := robotStatusWriter.RobotStatusTopic{}
	err = json.Unmarshal([]byte(result), &robotStatus)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %v", err)
	} else {
		return &pb.RobotStatusReply{
			MoveStatus:      robotStatus.Results.MoveStatus,
			ChargeStatus:    robotStatus.Results.ChargeState,
			SoftEstopStatus: robotStatus.Results.SoftEStopState,
			HardEstopStatus: robotStatus.Results.HardEStopState,
			PowerPercent:    int64(robotStatus.Results.PowerPercent),
			X:               robotStatus.Results.CurrentPose.X,
			Y:               robotStatus.Results.CurrentPose.Y,
			Theta:           robotStatus.Results.CurrentPose.Theta,
			Datetime:        robotStatus.TimeStamp,
			ErrorMessage:    "",
		}, nil
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
