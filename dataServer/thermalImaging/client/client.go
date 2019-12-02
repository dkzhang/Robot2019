package client

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/dataServer/thermalImaging/grpc"
	"google.golang.org/grpc"
)

func ThermalImaging(address string) (err error) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return fmt.Errorf("grpc CollectThermalImagingData grpc.Dial error: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewThermalImagingServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.CollectRenderAnalyze(ctx, &pb.ThermalImagingRequest{Tag: myUtil.FormatTime(time.Now())})
	if err != nil {
		log.Fatalf("could not reply: %v", err)
		return fmt.Errorf("grpc CollectThermalImagingData Reply error: %v", err)
	}
	log.Printf("reply = %v", r)

	//将结果写入redis

	return nil
}
