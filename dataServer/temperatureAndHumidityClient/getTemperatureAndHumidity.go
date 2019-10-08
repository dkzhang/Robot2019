package temperatureAndHumidityClient

import (
	pb "Robot2019/dataServer/temperatureAndHumidityClient/grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetTemperatureAndHumidity() ([]*TemperatureAndHumidityInfo, error) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTemperatureAndHumidityQueryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetTemperatureAndHumidity(ctx,
		&pb.TemperatureAndHumidityRequest{Datetime: time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Fatalf("could not get correct reply: %v", err)
		return nil, err
	}

	if len(r.GetErrorMessage()) != 0 {
		return nil, fmt.Errorf("%s", r.GetErrorMessage())
	}

	result := make([]*TemperatureAndHumidityInfo, len(r.GetThInfo()))
	for i := 0; i < len(r.GetThInfo()); i++ {
		result[i] = &TemperatureAndHumidityInfo{
			Temperature: r.GetThInfo()[i].GetTemperature(),
			Humidity:    r.GetThInfo()[i].GetHumidity(),
			Datetime:    r.GetThInfo()[i].GetDatetime(),
		}
	}

	return result, nil
}

type TemperatureAndHumidityInfo struct {
	Temperature float64
	Humidity    float64
	Datetime    string
}
