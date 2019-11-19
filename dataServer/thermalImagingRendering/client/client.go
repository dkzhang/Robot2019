package client

import (
	"Robot2019/applicationDriverForRobot/thermalImagingDataCollect/client"
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/dataServer/thermalImagingRendering/grpc"
	"google.golang.org/grpc"
)

func ThermalImagingCollectAndRender() {
	dataArray1, err := client.CollectThermalImagingData("")
	dataArray2, err := client.CollectThermalImagingData("")
	dataArray3, err := client.CollectThermalImagingData("")
	dataArray4, err := client.CollectThermalImagingData("")

	fmt.Printf("%v", err)

	dataArray := append(dataArray1, dataArray2...)
	dataArray = append(dataArray, dataArray3...)
	dataArray = append(dataArray, dataArray4...)

	ThermalImagingRender("", dataArray)

}

func ThermalImagingRender(address string, dataArray []float64) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
		Height:    8,
		Width:     32,
		Filepath:  "/ThermalImages/testTir/",
		Filename:  "test005",
	})

	if err != nil {
		log.Fatalf("could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
