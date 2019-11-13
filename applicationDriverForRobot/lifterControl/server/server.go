package main

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

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
	cmd := exec.Command("dir", "")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	theReply := pb.LifterControlReply{}
	if err != nil {
		theReply.ErrorMessage = err.Error()
	}

	return &theReply, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLifterControlServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
