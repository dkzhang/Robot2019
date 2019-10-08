package temperatureAndHumidityClient

import (
	pb "Robot2019/dataServer/temperatureAndHumidityClient/grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

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

func main() {
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

func GetTemperatureAndHumidity() ([]TemperatureAndHumidityInfo, error) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

}

type TemperatureAndHumidityInfo struct {
	Temperature float64
	Humidity    float64
	Datetime    string
}
