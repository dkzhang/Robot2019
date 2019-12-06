package main

import (
	_ "encoding/json"
	_ "github.com/gomodule/redigo/redis"
	"testing"
)

func TestServer_GetRobotStatus(t *testing.T) {

	server := server{}

	reply, err := server.GetRobotStatus(nil, nil)
	if err != nil {
		t.Errorf("GetRobotStatus error: %v", err)
	} else {
		t.Logf("GetRobotStatus reply = %v", *reply)
	}
}

/*
func TestServer_GetRobotStatus2(t *testing.T) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "140.143.16.113:50061"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Printf(" fatal error! did not connect: %v", err)
	}
	t.Logf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewRobotStatusServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Logf("context.WithTimeout() OK!")
	defer cancel()
	r, err := c.GetRobotStatus(ctx, &pb.RobotStatusRequest{Tag: myUtil.FormatTime(time.Now())})

	if err != nil {
		t.Fatalf("could not reply: %v", err)
	}
	t.Logf("reply = %v", r)
}*/
