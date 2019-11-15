package main

import (
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	//"os"
	//"os/exec"
	"time"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSinglePointMoveServer
}

func (s *server) Move(ctx context.Context, in *pb.SinglePointInfo) (*pb.MoveResponse, error) {
	log.Printf("Received: %v", *in)

	//实例化一个通信模块
	serverIPandPort := "192.168.10.10:31001"
	psm := socketCommunication.SocketManagementFactory(serverIPandPort)
	//test
	fmt.Printf("%v", psm)
	//构造单点移动命令（附随机数）并发送

	//循环接收传回的消息
	// 检查是否为所发命令的回复

	//如果是，则返回；如果不是，则继续循环接收

	return &pb.MoveResponse{}, nil
}

func GenerateMoveCommand(spi *pb.SinglePointInfo) (cmd string, uuid string) {
	cmd = "/api/move?"

	if spi.InfoMask&16 != 0 {
		//marker
		cmd += "marker=" + spi.Marker
	} else {
		//location
		cmd += fmt.Sprintf("location=%f,%f,%f", spi.LocationX, spi.LocationY, spi.LocationTheta)
	}

	// max_continuous_retries
	if spi.InfoMask&4 != 0 {
		cmd += fmt.Sprintf("&max_continuous_retries=%d", spi.MaxContinuousRetries)
	}

	// distance_tolerance
	if spi.InfoMask&2 != 0 {
		cmd += fmt.Sprintf("&distance_tolerance=%f", spi.DistanceTolerance)
	}

	// theta_tolerance
	if spi.InfoMask&1 != 0 {
		cmd += fmt.Sprintf("&theta_tolerance=%f", spi.ThetaTolerance)
	}

	rand.Seed(time.Now().Unix())
	uuid = fmt.Sprintf("%X", rand.Uint32())
	cmd += fmt.Sprintf("&uuid=%s", uuid)

	return cmd, uuid
}

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

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSinglePointMoveServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
