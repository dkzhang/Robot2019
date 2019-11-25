package server

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	dataCollect "Robot2019/applicationDriverForRobot/thermalImagingDataCollect/client"
	pb "Robot2019/dataServer/thermalImaging/grpc"
	dataAnalysis "Robot2019/dataServer/thermalImagingAnalysis/client"
	imageRender "Robot2019/dataServer/thermalImagingRendering/client"

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
	dataArray1, err := dataCollect.CollectThermalImagingData("")
	if err != nil {
		return nil, fmt.Errorf("data collect 1 error: %v", err)
	}
	dataArray2, err := dataCollect.CollectThermalImagingData("")
	if err != nil {
		return nil, fmt.Errorf("data collect 2 error: %v", err)
	}

	//合成数值数组
	dataArray := append(dataArray1, dataArray2...)
	const height = 8
	const width = 32

	//调用绘图服务绘图
	filepath, filename, err := imageRender.ThermalImagingRender("", dataArray)
	if err != nil {
		return nil, fmt.Errorf("ThermalImagingRender error: %v", err)
	}

	//调用分析服务进行热点分析
	analysisReport, err := dataAnalysis.ThermalImagingAnalyze("", &dataAnalysis.ThermalImagingDataStruct{
		DataArray: dataArray,
		Height:    height,
		Width:     width,
	})
	if err != nil {
		return nil, fmt.Errorf("dataAnalysis error: %v", err)
	}

	return &pb.ThermalImagingReply{
		Filepath:       filepath,
		Filename:       filename,
		DataArray:      dataArray,
		Height:         height,
		Width:          width,
		AnalysisReport: analysisReport,
		ErrorMessage:   "",
	}, nil
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
