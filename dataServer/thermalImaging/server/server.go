package server

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	//dataCollect "Robot2019/applicationDriverForRobot/thermalImagingDataCollect/client"
	pb "Robot2019/dataServer/thermalImaging/grpc"
	//dataAnalysis "Robot2019/dataServer/thermalImagingAnalysis/client"
	//imageRender "Robot2019/dataServer/thermalImagingRendering/client"

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

	/*
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

	*/
	return nil, nil
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

func MergeThermalArray(a1, a2, a3, a4 []float64, w, h int) (r []float64, newWidth, newHeight int, err error) {

	if len(a1) != w*h || len(a1) != w*h || len(a1) != w*h || len(a1) != w*h {
		return nil, -1, -1, fmt.Errorf("illegal array length")
	}

	r = make([]float64, w*h*4)
	//横向合并
	newWidth = w * 4
	newHeight = h

	for i := 0; i < h; i++ {
		iw := i * w
		iwn := i * newWidth
		copy(r[iwn+0*w:iwn+1*w], a1[iw:iw+w])
		copy(r[iwn+1*w:iwn+2*w], a2[iw:iw+w])
		copy(r[iwn+2*w:iwn+3*w], a3[iw:iw+w])
		copy(r[iwn+3*w:iwn+4*w], a4[iw:iw+w])
	}

	return
}
