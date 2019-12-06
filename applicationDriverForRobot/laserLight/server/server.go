package main

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"net"
	//"os"
	//"os/exec"
	"time"

	"github.com/stianeikeland/go-rpio"

	pb "Robot2019/applicationDriverForRobot/laserLight/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50071"
)

type server struct {
	pb.UnimplementedLaserLightServiceServer
}

func (s *server) SwitchLaserLight(ctx context.Context, in *pb.LaserLightRequest) (*pb.LaserLightReply, error) {
	log.Printf("Received: %v", in.GetTurnOn())

	//use go-rpio
	err := rpio.Open()
	if err != nil {
		return &pb.LaserLightReply{
			ErrorMessage: err.Error(),
		}, nil
	}
	defer rpio.Close()

	pin := rpio.Pin(4)
	pin.Output() // Output mode

	turnOn := in.GetTurnOn()
	if turnOn == true {
		pin.High()
	} else {
		pin.Low()
	}

	return &pb.LaserLightReply{}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLaserLightServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s \n", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
