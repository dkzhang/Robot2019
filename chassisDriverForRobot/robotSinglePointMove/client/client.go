package main

import (
	"context"
	"log"
	"time"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	"google.golang.org/grpc"
)

func main() {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "localhost:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewSinglePointMoveClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r1, err := c.Move(ctx, &pb.SinglePointInfo{})
	r2, err := c.MoveAndWaitForArrival(ctx, &pb.SinglePointInfo{})

	if err != nil {
		log.Fatalf("could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
