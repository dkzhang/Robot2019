package temperatureAndHumidityClient

import (
	pb "Robot2019/dataServer/temperatureAndHumidityClient/grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"
)

func TestGetTemperatureAndHumidity(t *testing.T) {
	go mockServer()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 3)
		ths, err := GetTemperatureAndHumidity()
		if err != nil {
			t.Fatalf("GetTemperatureAndHumidity error: %v", err)
		}

		for i := 0; i < len(ths); i++ {
			t.Logf("GetTemperatureAndHumidity result [%d]: %v", i, *(ths[i]))
		}
	}
}

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedTemperatureAndHumidityQueryServiceServer
}

func (s *server) GetTemperatureAndHumidity(ctx context.Context, in *pb.TemperatureAndHumidityRequest) (*pb.TemperatureAndHumidityReply, error) {

	reply := pb.TemperatureAndHumidityReply{
		ThInfo:       make([]*pb.TemperatureAndHumidityInfo, 3),
		ErrorMessage: "",
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		reply.ThInfo[i] = &pb.TemperatureAndHumidityInfo{
			Temperature: 20 + rand.Float64()*10,
			Humidity:    50 + rand.Float64()*20,
			Datetime:    time.Now().Format("2006-01-02 15:04:05"),
		}
	}

	return &reply, nil
}

func mockServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTemperatureAndHumidityQueryServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
