package main

import (
	"Robot2019/chassisDriverForRobot/common"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"context"
	"fmt"
	"log"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
)

func (s *server) Move(ctx context.Context, in *pb.SinglePointInfo) (*pb.MoveResponse, error) {
	log.Printf("Received: %v", *in)

	//实例化一个通信模块
	serverIPandPort := "192.168.10.10:31001"
	psm := socketCommunication.SocketManagementFactory(serverIPandPort)
	//test
	fmt.Printf("%v", psm)

	//构造单点移动命令（附随机数）并发送
	cmdStruct := socketCommunication.CommandStruct{Name: "Single Move"}
	cmdStruct.Command, cmdStruct.UUID = GenerateMoveCommand(in)
	psm.CommandChan <- &cmdStruct

	//循环接收传回的消息
	errorCount := 0
	const errorMax = 5
	for {
		select {
		case result := <-psm.ResultChan:
			// 检查是否为所发命令的回复
			pcr, err := common.CommandDetection(result, cmdStruct.UUID)

			//检查是否出错并计数
			if err != nil {
				errorCount++
				if errorCount > errorMax {
					return nil, fmt.Errorf("CommandDetection error too many times")
				}
			} else {
				if pcr != nil {
					//是所发命令的回复,则进一步解析
					smcr := SingleMoveCommandResponse{}
					if smcr.UnmarshalJSON(result) != nil {
						return &pb.MoveResponse{
							Command:              "",
							Uuid:                 0,
							Status:               "",
							ErrorMessage:         "",
							TaskId:               0,
							XXX_NoUnkeyedLiteral: struct{}{},
							XXX_unrecognized:     nil,
							XXX_sizecache:        0,
						}, nil
					}

				}
			}

		case feedback := <-psm.FeedbackChan:
			log.Printf("socketCommunication feedback: %v", feedback)
		}

	}

	return &pb.MoveResponse{}, nil
}
