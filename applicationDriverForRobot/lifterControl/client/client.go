package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/applicationDriverForRobot/lifterControl/grpc"
	"google.golang.org/grpc"
)

func LifterControl(para int64) (err error) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "192.168.10.21:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return fmt.Errorf(" fatal error! LifterControl did not connect: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewLifterControlServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	log.Printf("context.WithTimeout() OK!")

	defer cancel()
	r, err := c.ControlTheLifter(ctx, &pb.LifterControlRequest{Para: para})

	if err != nil {
		return fmt.Errorf(" fatal error! LifterControl could not reply: %v", err)
	} else if len(r.ErrorMessage) != 0 {
		return fmt.Errorf(" LifterControl reply Error Message: %s", r.ErrorMessage)
	} else {
		log.Printf("LifterControl reply = %v", r)
		return nil
	}
}
