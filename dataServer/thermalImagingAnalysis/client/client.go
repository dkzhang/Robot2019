package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/dataServer/thermalImagingAnalysis/grpc"
	"google.golang.org/grpc"
)

func ThermalImagingAnalyze(address string, dataArray []float64) (level, analysisReport string, err error) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return "", "", fmt.Errorf("grpc CollectThermalImagingData grpc.Dial error: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewThermalImagingAnalysisServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.AnalyzeThermalImaging(ctx, &pb.ThermalImagingAnalysisRequest{
		DataArray: dataArray,
	})
	if err != nil {
		log.Fatalf("could not reply: %v", err)
		return "", "", fmt.Errorf("grpc AnalyzeThermalImaging could not reply: %v", err)
	}

	log.Printf("reply = %v", r)
	if r.ErrorMessage != "" {
		log.Fatalf("reply error message: %v", err)
		return "", "", fmt.Errorf("grpc AnalyzeThermalImaging Reply error message = : %v", err)
	}

	return r.Level, r.AnalysisReport, nil
}
