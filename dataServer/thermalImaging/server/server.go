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
	imageRender "Robot2019/dataServer/thermalImagingRendering/client"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedThermalImagingServiceServer
}

func (s *server) CollectRenderAnalyze(ctx context.Context, in *pb.ThermalImagingRequest) (*pb.ThermalImagingReply, error) {
	log.Printf("Received: %v", in.GetTag())

	return CollectRenderAnalyze()
}

func CollectRenderAnalyze() (*pb.ThermalImagingReply, error) {

	//分别从两个树莓派收集数据
	mdata1, err := dataCollect.CollectThermalImagingData("192.168.10.23:50061")
	if err != nil {
		return nil, fmt.Errorf("data collect 1 error: %v", err)
	}
	mdata2, err := dataCollect.CollectThermalImagingData("192.168.10.25:50061")
	if err != nil {
		return nil, fmt.Errorf("data collect 2 error: %v", err)
	}

	var a1, a2, a3, a4 []float64
	const h, w = 8, 8
	//根据热红外测温矩阵实际部署位置调整数组顺序
	if mdata1[0].Id == 0x69 {
		a1 = mdata1[0].Data
		a2 = mdata1[1].Data
	}else{
		a1 = mdata1[1].Data
		a2 = mdata1[0].Data
	}

	if mdata2[0].Id == 0x69 {
		a3 = mdata2[0].Data
		a4 = mdata2[1].Data
	}else{
		a3 = mdata2[1].Data
		a4 = mdata2[0].Data
	}

	//合成数值数组
	dataArray, newWidth, newHeight, err := MergeThermalArray(a1, a2, a3, a4, w, h)
	if err != nil {
		return nil, fmt.Errorf("MergeThermalArray error: %v", err)
	}

	fmt.Printf("dataArray(%d,%d) = %v", newWidth, newHeight,dataArray)
	//调用绘图服务绘图
	filepath, filename, err := imageRender.ThermalImagingRender("localhost:50061", dataArray, newWidth, newHeight)
	if err != nil {
		return nil, fmt.Errorf("ThermalImagingRender error: %v", err)
	}

	fmt.Printf("%s \n", filepath)
	fmt.Printf("%s \n", filename)

	/*
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
	return &pb.ThermalImagingReply{
		Filepath:       filepath,
		Filename:       filename,
		DataArray:      dataArray,
		Height:         int32(newHeight),
		Width:          int32(newWidth),
		AnalysisReport: "analysisReport",
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

func MergeThermalArray(a1, a2, a3, a4 []float64, w, h int) (r []float64, newWidth, newHeight int, err error) {

	if len(a1) != w*h || len(a1) != w*h || len(a1) != w*h || len(a1) != w*h {
		return nil, -1, -1,
			fmt.Errorf("illegal array length: %d,%d,%d,%d <=> %d, %d", len(a1), len(a2), len(a3), len(a4), w, h)
	}

	fmt.Printf("a1,a2,a3,a4: %d,%d,%d,%d <=> %d, %d \n", len(a1), len(a2), len(a3), len(a4), w, h)

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
