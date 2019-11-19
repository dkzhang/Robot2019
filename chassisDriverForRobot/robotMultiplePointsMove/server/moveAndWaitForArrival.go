package main

import (
	pb "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/grpc"
	"context"
)

func (s *server) MoveAndWaitForArrival(ctx context.Context, in *pb.MultiplePointsInfo) (*pb.MoveAndWaitForArrivalResponse, error) {
	// 暂时不启用该函数
	return &pb.MoveAndWaitForArrivalResponse{}, nil
}
