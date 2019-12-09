package client

import (
	"Robot2019/chassisDriverForRobot/configuration"
	pb "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/grpc"
	"google.golang.org/grpc"

	"context"
	"log"
	"time"
)

func MMove(markers []string) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := configuration.MultiplePointsMove_ADDRESS

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Printf(" fatal error! did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewMultiplePointsMoveClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.Move(ctx, &pb.MultiplePointsInfo{
		InfoMask: 0,
		Markers:  markers,
	})
	if err != nil {
		log.Printf(" fatal error! could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
