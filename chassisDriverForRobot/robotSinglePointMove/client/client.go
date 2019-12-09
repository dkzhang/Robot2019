package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	"google.golang.org/grpc"
)

func MoveAndWaitForArrival(marker string) (err error) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "localhost:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return fmt.Errorf(" fatal error! MoveAndWaitForArrival did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewSinglePointMoveClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.MoveAndWaitForArrival(ctx, &pb.SinglePointInfo{
		InfoMask: 16,
		Marker:   marker,
	})

	if err != nil {
		return fmt.Errorf(" fatal error! MoveAndWaitForArrival could not reply: %v", err)
	} else if len(r.ErrorMessage) != 0 {
		return fmt.Errorf(" MoveAndWaitForArrival reply Error Message: %s", r.ErrorMessage)
	} else {
		log.Printf("MoveAndWaitForArrival reply = %v", r)
		return nil
	}
}

func Move(marker string) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "localhost:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Printf(" fatal error! did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewSinglePointMoveClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.Move(ctx, &pb.SinglePointInfo{
		InfoMask: 16,
		Marker:   marker,
	})

	if err != nil {
		log.Printf(" fatal error! could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
