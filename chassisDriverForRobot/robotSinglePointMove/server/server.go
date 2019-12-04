package server

import (
	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
)

type Server struct {
	pb.UnimplementedSinglePointMoveServer
}
