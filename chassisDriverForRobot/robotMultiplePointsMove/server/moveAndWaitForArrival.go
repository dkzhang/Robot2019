package main

import (
	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	"context"
)

func (s *server) MoveAndWaitForArrival(ctx context.Context, in *pb.SinglePointInfo) (*pb.MoveAndWaitForArrivalResponse, error) {
	// 暂时不启用该函数
	return &pb.MoveAndWaitForArrivalResponse{}, nil
}
