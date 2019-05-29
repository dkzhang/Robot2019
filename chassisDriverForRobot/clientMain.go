package main

import (
	robotMove "Robot2019/chassisDriverForRobot/robotMove/grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := robotMove.NewRobotMoveClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SinglePointMove(ctx, &robotMove.SinglePointInfo{
		InfoMask: 1,
		Marker:   "2",
	})
	if err != nil {
		log.Fatalf("could not single point move: %v", err)
	}
	log.Printf("get response: %v", r)
}
