package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/applicationDriverForRobot/laserLight/grpc"
	"google.golang.org/grpc"
)

func SwitchLaserLight(turnOn bool) (err error) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "192.168.10.23:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return fmt.Errorf(" fatal error! SwitchLaserLight did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewLaserLightServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()
	r, err := c.SwitchLaserLight(ctx, &pb.LaserLightRequest{TurnOn: turnOn})
	if err != nil {
		return fmt.Errorf(" fatal error! SwitchLaserLight could not reply: %v", err)
	} else if len(r.ErrorMessage) != 0 {
		return fmt.Errorf(" SwitchLaserLight reply Error Message: %s", r.ErrorMessage)
	} else {
		log.Printf("SwitchLaserLight reply = %v", r)
		return nil
	}
}
