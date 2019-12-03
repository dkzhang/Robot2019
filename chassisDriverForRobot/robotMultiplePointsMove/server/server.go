package server

import (
	"Robot2019/chassisDriverForRobot/common"
	pb "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/grpc"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"context"
	"fmt"
	"log"
)

type Server struct {
	pb.UnimplementedMultiplePointsMoveServer
}

func (s *Server) Move(ctx context.Context, in *pb.MultiplePointsInfo) (*pb.MoveResponse, error) {
	log.Printf("Received: %v", *in)

	//实例化一个通信模块
	serverIPandPort := "192.168.10.10:31001"
	psm := socketCommunication.SocketManagementFactory(serverIPandPort)
	defer psm.Cancel()

	//构造多点移动命令（附随机数）并发送
	cmdStruct := socketCommunication.CommandStruct{Name: "Multiple Move"}
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
					mmcr := MultipleMoveCommandResponse{}
					if mmcr.UnmarshalJSON(result) == nil {
						//如果解析成功，则用smcr填写，含TaskID
						return &pb.MoveResponse{
							Command:      mmcr.Command,
							Status:       mmcr.Status,
							ErrorMessage: mmcr.ErrorMessage,
							TaskId:       mmcr.TaskID,
						}, nil
					} else {
						//如果未解析成功，也按成功返回，用pcr填写，不含TaskID
						return &pb.MoveResponse{
							Command:      pcr.Command,
							Status:       pcr.Status,
							ErrorMessage: pcr.ErrorMessage,
							TaskId:       "",
						}, nil
					}
				}
			}

		case feedback := <-psm.FeedbackChan:
			log.Printf("socketCommunication feedback: %v", feedback)
		}
	}

	//return &pb.MoveResponse{}, nil
}

func (s *Server) MoveAndWaitForArrival(ctx context.Context, in *pb.MultiplePointsInfo) (*pb.MoveAndWaitForArrivalResponse, error) {
	// 暂时不启用该函数
	return &pb.MoveAndWaitForArrivalResponse{}, nil
}
