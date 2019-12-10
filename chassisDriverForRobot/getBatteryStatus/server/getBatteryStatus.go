package server

import (
	"Robot2019/chassisDriverForRobot/common"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"context"
	"fmt"
	"log"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
)

func GetBatteryStatus(ctx context.Context, in *pb.SinglePointInfo) (*pb.MoveResponse, error) {
	//log.Printf("Received: %v", *in)

	//实例化一个通信模块
	serverIPandPort := "192.168.10.10:31001"
	psm := socketCommunication.SocketManagementFactory(serverIPandPort)
	defer psm.Cancel()

	//构造查询电池状态命令（不带随机数）并发送
	cmdStruct := socketCommunication.CommandStruct{}
	cmdStruct.Name, cmdStruct.Command = GenerateGetCommand()
	psm.CommandChan <- &cmdStruct

	//循环接收传回的消息
	for {
		select {
		case result := <-psm.ResultChan:
			fmt.Printf("receive from socker: %s", result)
			// 检查是否为所发命令的回复
			pcr, err := common.CommandDetection(result, cmdStruct.UUID)

			//检查是否出错并计数
			if err != nil {
				return nil, fmt.Errorf("CommandDetection error: %v", err)
			} else if pcr != nil {
				//是所发命令的回复,则进一步解析
				fmt.Printf("reply = %s", result)
				return &pb.MoveResponse{
					Command:              "",
					Status:               "",
					ErrorMessage:         "",
					TaskId:               "",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				}, nil
			}

		case feedback := <-psm.FeedbackChan:
			log.Printf("socketCommunication feedback: %v", feedback)
		}
	}

	//return &pb.MoveResponse{}, nil
}
