package client

import (
	"context"
	"log"
	"time"

	pb "Robot2019/applicationDriverForRobot/laserLight/grpc"
	"google.golang.org/grpc"
)

func SwitchLaserLight(turnOn bool) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "192.168.1.109:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Printf(" fatal error! did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewLaserLightServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()
	r, err := c.SwitchLaserLight(ctx, &pb.LaserLightRequest{TurnOn: turnOn})
	if err != nil {
		log.Printf(" fatal error! could not reply: %v", err)
	}
	log.Printf("reply = %v", r)
}
