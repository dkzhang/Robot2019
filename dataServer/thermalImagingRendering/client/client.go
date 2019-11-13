package main

import (
	"context"
	"log"
	"time"

	pb "Robot2019/dataServer/thermalImagingRendering/grpc"
	"google.golang.org/grpc"
)

func main() {
	/////////////////////////////////
	// Set up a connection to the server.
	address := "localhost:50061"

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
		DataArray: nil,
		Height:    8,
		Width:     32,
		Filepath:  "~/testTir",
		Filename:  "test001",
	})

	if err != nil {
		log.Fatalf("could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
