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

func (s *server) AnalyzeThermalImaging(ctx context.Context, in *pb.ThermalImagingAnalysisRequest) (*pb.ThermalImagingAnalysisReply, error) {
	log.Printf("Received: %v", *in)

	return &pb.ThermalImagingAnalysisReply{}, nil
}

func AnalyzeThermalImaging(dataArray []float64) (level, report string, err error) {
	if len(dataArray) <= 1 {
		return "", "", fmt.Errorf("dataArray length <%d> is too small", len(dataArray))
	}

	const UpperAlarmLimit = 60.0
	const LowerAlarmLimit = 0.0

	maxValue := dataArray[0]
	minValue := dataArray[0]
	sumValue := 0.0
	UpperCount, LowerCount := 0, 0

	for i := 0; i < len(dataArray); i++ {
		if dataArray[i] > maxValue {
			maxValue = dataArray[i]
		} else if dataArray[i] < minValue {
			minValue = dataArray[i]
		}

		sumValue += dataArray[i]

		if dataArray[i] > UpperAlarmLimit {
			UpperCount++
		} else if dataArray[i] < LowerAlarmLimit {
			LowerCount++
		}
	}

	if UpperCount >= 3 || LowerCount >= 3 {
		level = "Warning"
	} else {
		level = "Info"
	}

	report = fmt.Sprintf("统计温度点数量:%d；其中最大值为%f，最小值为%f, 平均值为%f；其中超过上限%f的点数为%d，低于下限%f的点数为%d \n",
		len(dataArray), maxValue, minValue, sumValue/float64(len(dataArray)), UpperAlarmLimit, UpperCount, LowerAlarmLimit, LowerCount)
	return
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
