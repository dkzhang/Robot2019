package main

import (
	"context"
	"log"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
)

func (s *server) MoveAndWaitForArrival(ctx context.Context, in *pb.SinglePointInfo) (*pb.MoveAndWaitForArrivalResponse, error) {
	log.Printf("Received: %v", *in)

	//实例化一个通信模块

	//构造定期获取机器人状态命令（附随机数）并发送

	//构造单点移动命令（附随机数）并发送

	//循环接收传回的消息
	//检查是否为所发命令的回复
	//如果是，且提示错误，则返回错误
	//其余情况继续循环

	//检查是否为订阅消息，且完成移动
	//如果是，则返回成功
	//如果尚在移动，则继续循环；如果移动出错，则返回错误

	return &pb.MoveAndWaitForArrivalResponse{}, nil
}
