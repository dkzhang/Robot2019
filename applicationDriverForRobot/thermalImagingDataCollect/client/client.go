package client

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/applicationDriverForRobot/thermalImagingDataCollect/grpc"
	"google.golang.org/grpc"
)

func CollectThermalImagingData(address string) (data []*pb.ModelData, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, fmt.Errorf("grpc CollectThermalImagingData grpc.Dial error: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewThermalImagingDataCollectServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.CollectThermalImagingData(ctx, &pb.ThermalImagingDataCollectRequest{Tag: myUtil.FormatTime(time.Now())})
	if err != nil {
		log.Fatalf("could not reply: %v", err)
		return nil, fmt.Errorf("grpc CollectThermalImagingData Reply error: %v", err)
	}
	log.Printf("reply = %v", r)

	return r.Mdata, nil
}
