package server

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "Robot2019/dataServer/thermalImaging/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50061"
)

type server struct {
	pb.UnimplementedThermalImagingServiceServer
}

func (s *server) CollectRenderAnalyze(ctx context.Context, in *pb.ThermalImagingRequest) (*pb.ThermalImagingReply, error) {
	log.Printf("Received: %v", in.GetTag())

	//分别从两个树莓派收集数据

	//合成数值数组

	//调用绘图服务绘图

	//调用分析服务进行热点分析

	return &pb.ThermalImagingReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterThermalImagingServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
