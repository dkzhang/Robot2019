package client

import (
	pb "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/grpc"
	"google.golang.org/grpc"

	"context"
	"log"
	"time"
)

func Move(marker string) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "192.168.10.27:50072"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewMultiplePointsMoveClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.Move(ctx, &pb.MultiplePointsInfo{
		InfoMask: 0,
		Markers:  []string{},
	})
	if err != nil {
		log.Fatalf("could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
