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

	pb "Robot2019/applicationDriverForRobot/lifterControl/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedLifterControlServiceServer
}

func (s *server) ControlTheLifter(ctx context.Context, in *pb.LifterControlRequest) (*pb.LifterControlReply, error) {
	log.Printf("Received: %v", in.GetPara())

	//system call
	//cmd := exec.Command("dir", "")
	//cmd.Stdout = os.Stdout
	//err := cmd.Run()

	//use go-rpio
	err := rpio.Open()
	if err != nil {
		return &pb.LifterControlReply{
			ErrorMessage: err.Error(),
		}, nil
	}
	defer rpio.Close()

	para := in.GetPara()
	if para > 0 {
		// Init
		pinDOWN := rpio.Pin(5)
		pinDOWN.Output() // Output mode
		pinDOWN.Low()    // Set pinDOWN High

		//上升
		pinUP := rpio.Pin(6)
		pinUP.Output() // Output mode
		pinUP.High()   // Set pinUP High
		time.Sleep(time.Millisecond * time.Duration(para))
		pinUP.Low() // Set pinUP Low
	} else {
		//Init
		pinUP := rpio.Pin(6)
		pinUP.Output() // Output mode
		pinUP.Low()    // Set pinUP High

		//下降
		pinDOWN := rpio.Pin(5)
		pinDOWN.Output() // Output mode
		pinDOWN.High()   // Set pinDOWN High
		time.Sleep(time.Millisecond * time.Duration(-para))
		pinDOWN.Low() // Set pinDOWN Low

	}

	return &pb.LifterControlReply{}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLifterControlServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s \n", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
