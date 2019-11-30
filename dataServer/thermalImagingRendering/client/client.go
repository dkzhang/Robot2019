package client

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/dataServer/thermalImagingRendering/grpc"
	"google.golang.org/grpc"
)

func ThermalImagingRender(address string, dataArray []float64, width, height int) (filepath string, filename string, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return "", "", fmt.Errorf("grpc.Dial error: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewThermalImagingRenderingServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()
	r, err := c.ThermalImagingRender(ctx, &pb.ThermalImagingRenderingRequest{
		DataArray: dataArray,
		Height:    int32(height),
		Width:     int32(width),
		Filepath:  "/ThermalImages/",
		Filename:  myUtil.FormatTime(time.Now()),
	})

	if err != nil {
		log.Fatalf("could not reply: %v", err)
		return "", "", fmt.Errorf("ThermalImagingRender reply error: %v", err)
	}
	log.Printf("reply = %v", r)
	return filepath, filename, nil
}

//docker run -p 50061:50061 -v /home/dkzhang/tirImages:/ThermalImages -d registry.cn-beijing.aliyuncs.com/dkzhang/robot2019-dataserver-tir:3.0
