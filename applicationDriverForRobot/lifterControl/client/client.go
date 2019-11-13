package main

import (
	"context"
	"log"
	"time"

	pb "Robot2019/applicationDriverForRobot/lifterControl/grpc"
	"google.golang.org/grpc"
)

func main() {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "140.143.16.113:50061"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewLifterControlServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()
	r, err := c.ControlTheLifter(ctx, &pb.LifterControlRequest{Para: 1000})

	if err != nil {
		log.Fatalf("could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
