package server

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "Robot2019/dataServer/thermalImagingAnalysis/grpc"

	"google.golang.org/grpc"
)

const (
	port = ":50061"
)

type server struct {
	pb.UnimplementedThermalImagingAnalysisServiceServer
}

func (s *server) CollectRenderAnalyze(ctx context.Context, in *pb.ThermalImagingAnalysisRequest) (*pb.ThermalImagingAnalysisReply, error) {
	log.Printf("Received: %v", *in)

	return &pb.ThermalImagingAnalysisReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterThermalImagingAnalysisServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
