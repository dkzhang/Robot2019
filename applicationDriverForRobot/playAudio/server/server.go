package main

import (
	"Robot2019/myUtil"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	//"os"
	//"os/exec"
	"time"

	pb "Robot2019/applicationDriverForRobot/playAudio/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50061"
)

type server struct {
	pb.UnimplementedPlayAudioServiceServer
}

func (s *server) ControlTheLifter(ctx context.Context, in *pb.PlayAudioRequest) (*pb.PlayAudioReply, error) {
	log.Printf("Received: %v", in.GetAudioName())

	//system call
	cmd := exec.Command("play", "BGPsong.wav")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		return &pb.PlayAudioReply{ErrorMessage: err.Error()}, nil
	} else {
		return &pb.PlayAudioReply{}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPlayAudioServiceServer(s, &server{})
	fmt.Printf("Begin to serve %s \n", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
