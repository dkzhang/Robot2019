package server

import (
	"Robot2019/chassisDriverForRobot/common"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"Robot2019/chassisDriverForRobot/subscribeRobotStatus"
	"context"
	"fmt"
	"log"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
)

func (s *Server) MoveAndWaitForArrival(ctx context.Context, in *pb.SinglePointInfo) (*pb.MoveAndWaitForArrivalResponse, error) {
	//实例化一个通信模块
	log.Printf("Received: %v", *in)

	//实例化一个通信模块
	serverIPandPort := "192.168.10.10:31001"
	psm := socketCommunication.SocketManagementFactory(serverIPandPort)
	defer psm.Cancel()

	//构造单点移动命令（附随机数）并发送
	cmdMoveStruct := socketCommunication.CommandStruct{Name: "Single Move"}
	cmdMoveStruct.Command, cmdMoveStruct.UUID = GenerateMoveCommand(in)
	psm.CommandChan <- &cmdMoveStruct

	//构造定期获取机器人状态命令（附随机数）并发送
	cmdSubscribeStruct := socketCommunication.CommandStruct{Name: "Subscribe Robot Status"}
	cmdSubscribeStruct.Command, cmdSubscribeStruct.UUID = subscribeRobotStatus.GenerateSubscribeRobotStatusCommand(nil)
	psm.CommandChan <- &cmdSubscribeStruct

	//循环接收传回的消息
	cmdMoveFlag := false
	cmdSubscribeFlag := false
	var err error

	for {
		select {
		case result := <-psm.ResultChan:
			if cmdMoveFlag == false {
				//命令解析
				//检查是否为所发命令的回复
				cmdMoveFlag, err = CmdResponseParse(result, cmdMoveStruct.UUID)

				if err != nil {
					return &pb.MoveAndWaitForArrivalResponse{
						ErrorMessage: fmt.Sprintf("Cmd Move Response Parse error: %v", err),
					}, nil
				} else if cmdMoveFlag == true {
					continue
				}
			}

			if cmdSubscribeFlag == false {
				//命令解析
				//检查是否为所发命令的回复
				cmdSubscribeFlag, err = CmdResponseParse(result, cmdSubscribeStruct.UUID)

				if err != nil {
					return &pb.MoveAndWaitForArrivalResponse{
						ErrorMessage: fmt.Sprintf("Cmd Subcrible Response Parse error: %v", err),
					}, nil
				} else if cmdSubscribeFlag == true {
					continue
				}
			}

			if cmdMoveFlag == true && cmdSubscribeFlag == true {
				//订阅消息解析
				//检查是否为订阅消息，且完成移动
				//如果是，则返回成功
				//如果尚在移动，则继续循环；如果移动出错，则返回错误
				if SubscribeResponseParse(result) == true {
					return &pb.MoveAndWaitForArrivalResponse{
						ErrorMessage: "",
					}, nil
				}
			}

		case feedback := <-psm.FeedbackChan:
			log.Printf("socketCommunication feedback: %v", feedback)
		}
	}

	return &pb.MoveAndWaitForArrivalResponse{}, nil
}

// 解析命令的响应报文
// 如果收到的不是所发命令对应的响应报文，且没有出错，返回false，nil
// 如果是对应的响应报文，且状态正确，返回true，nil
// 如果出错，返回error
func CmdResponseParse(result string, uuid string) (bool, error) {
	// 检查是否为所发命令的回复
	pcr, err := common.CommandDetection(result, uuid)
	//检查是否出错
	if err != nil {
		return false, fmt.Errorf("CommandDetection error: %v", err)
	} else {
		if pcr != nil {
			//是所发命令的回复,则进一步解析
			smcr := SingleMoveCommandResponse{}
			if smcr.UnmarshalJSON(result) == nil {
				//如果解析成功，则用smcr，含TaskID
				log.Printf("Single Move respone: %v", smcr)
				if smcr.Status == "OK" {
					return true, nil
				} else {
					return false, fmt.Errorf("Single Move respone error: %v", smcr)
				}
			} else {
				//如果未解析成功，也按成功返回，用pcr，不含TaskID
				if pcr.Status == "OK" {
					return true, nil
				} else {
					return false, fmt.Errorf("Single Move respone error: %v", pcr)
				}
			}
		}
	}
	return false, nil
}

// 解析命令的响应报文
// 如果收到的不是callback消息订阅，或者不是robot_status主题，或者运动没有完成，返回false
// 如果是对应的响应报文，且指示移动命令已完成，返回true
// 如果出错，由于订阅消息的可重复性，直接忽略，返回false
func SubscribeResponseParse(result string) bool {

	pct, err := common.CallbackTopicDetection(result, "robot_status")

	//检查是否出错
	if err != nil {
		log.Printf("CallbackTopicDetection error: %v", err)
		return false
	} else {
		if pct != nil {
			//是robot status的订阅消息,则进一步解析
			robotStatus := subscribeRobotStatus.RobotStatusTopic{}
			err = robotStatus.UnmarshalJSON(result)
			if err != nil {
				return false
			} else {
				if robotStatus.Results.MoveStatus == "succeeded" {
					return true
				}
			}
		}
	}
	return false
}
