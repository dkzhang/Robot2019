package main

import (
	"Robot2019/myUtil"
	"context"
	_ "context"
	"encoding/json"
	"log"
	_ "log"
	"testing"
	"time"

	pb "Robot2019/dataServer/robotStatusServer/grpc"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
)

func TestServer_GetRobotStatus(t *testing.T) {
	//连接redis容器，读取相关状态信息
	redisConn, err := redis.Dial("tcp", "myRedis001:6379")
	if err != nil {
		t.Fatalf("redis dial error: %v", err)
		return
	}
	defer redisConn.Close()

	theStatus := pb.RobotStatusReply{
		MoveStatus:      "Running",
		ChargeStatus:    false,
		SoftEstopStatus: false,
		HardEstopStatus: false,
		PowerPercent:    99,
		X:               1.1,
		Y:               2.2,
		Theta:           3.3,
		Datetime:        myUtil.FormatTime(time.Now()),
		ErrorMessage:    "",
	}

	theJSON, err := json.Marshal(theStatus)
	if err != nil {
		t.Fatalf("json marshal error: %v", theStatus)
	}

	result, err := redis.String(redisConn.Do("SET", "CurrentRobotStatus", theJSON))
	if err != nil {
		t.Fatalf("Set CurrentRobotStatus error: %v", err)
		return
	}
	t.Logf("SET result = %v", result)

	for i := 0; i < 20; i++ {
		t.Logf("sleep: %d \n", i)
	}

	/////////////////////////////////
	// Set up a connection to the server.
	address := "localhost:50061"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewRobotStatusServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetRobotStatus(ctx, &pb.RobotStatusRequest{Tag: myUtil.FormatTime(time.Now())})

	if err != nil {
		log.Fatalf("could not reply: %v", err)
	}
	log.Printf("reply = %v", r)

}
