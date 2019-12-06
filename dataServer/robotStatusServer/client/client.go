package client

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/dataServer/robotStatusServer/grpc"
	"google.golang.org/grpc"
)

func GetRobotStatus() (reply *pb.RobotStatusReply, err error) {

	address := "192.168.10.27:50080"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Printf(" fatal error! did not connect: %v", err)
		return nil, fmt.Errorf("grpc CollectThermalImagingData grpc.Dial error: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewRobotStatusServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.GetRobotStatus(ctx, &pb.RobotStatusRequest{
		Tag: myUtil.FormatTime(time.Now()),
	})

	if err != nil {
		log.Printf(" fatal error! could not reply: %v", err)
		return nil, fmt.Errorf("grpc GetRobotStatus could not reply: %v", err)
	}

	log.Printf("reply = %v", r)
	if r.ErrorMessage != "" {
		log.Printf(" fatal error! reply error message: %v", err)
		return nil, fmt.Errorf("grpc GetRobotStatus Reply error message = : %v", err)
	}

	return r, nil
}
