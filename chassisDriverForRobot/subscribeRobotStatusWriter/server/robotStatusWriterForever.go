package server

import (
	"Robot2019/cache"
	"Robot2019/chassisDriverForRobot/common"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"Robot2019/myUtil"
	"fmt"
	"log"
	"time"
)

func RobotStatusWriterForever() {

	//实例化一个通信模块
	serverIPandPort := "192.168.10.10:31001"
	psm := socketCommunication.SocketManagementFactory(serverIPandPort)
	defer psm.Cancel()

	//构造定期获取机器人状态命令（附随机数）并发送
	cmdSubscribeStruct := socketCommunication.CommandStruct{Name: "Subscribe Robot Status"}
	cmdSubscribeStruct.Command, cmdSubscribeStruct.UUID = GenerateSubscribeRobotStatusCommand(nil)
	psm.CommandChan <- &cmdSubscribeStruct

	//循环接收传回的消息
	cmdSubscribeFlag := false
	var err error

	opts := &cache.RedisOpts{
		Host: cache.RedisHost,
	}
	theRedis := cache.NewRedis(opts)

	for {
		select {
		case strResult := <-psm.ResultChan:

			for _, strJSON := range myUtil.SplitJSON(strResult) {

				if cmdSubscribeFlag == false {
					//命令解析
					//检查是否为所发命令的回复
					cmdSubscribeFlag, err = CmdResponseParse(strJSON, cmdSubscribeStruct.UUID)

					if err != nil {
						log.Printf("CmdResponseParse error: %v", err)
					} else if cmdSubscribeFlag == true {
						continue
					}
				}

				//订阅消息解析
				//检查是否为订阅消息robot status，如果是，则存入redis
				log.Printf("check for subscribe: %s", strJSON)

				if err = SubscribeResponseParse(theRedis, strJSON); err != nil {
					log.Printf("SubscribeResponseParse error: %v", err)
				}

			}

		case feedback := <-psm.FeedbackChan:
			log.Printf("socketCommunication feedback: %v", feedback)
		}
	}

	return
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
			if pcr.Status == "OK" {
				log.Printf("Command OK: %v", pcr)
				return true, nil
			} else {
				return false, fmt.Errorf("respone error: %v", pcr)
			}
		}
	}
	return false, nil
}

// 解析命令的响应报文
// 如果收到的不是callback消息订阅，或者不是robot_status主题，或者运动没有完成，返回false
// 如果是对应的响应报文，且指示移动命令已完成，返回true
// 如果出错，由于订阅消息的可重复性，直接忽略，返回false
func SubscribeResponseParse(theRedis *cache.Redis, result string) error {

	pct, err := common.CallbackTopicDetection(result, "robot_status")

	//检查是否出错
	if err != nil {
		return fmt.Errorf("CallbackTopicDetection error: %v", err)
	} else {
		if pct != nil {
			//是robot status的订阅消息,则进一步解析
			robotStatus := RobotStatusTopic{}
			err = robotStatus.UnmarshalJSON(result)
			if err != nil {
				return fmt.Errorf("robotStatus.UnmarshalJSON error: %v", err)
			} else {
				robotStatus.TimeStamp = myUtil.FormatTime(time.Now())
				//write to redis
				theRedis.ListMaxLenRPush("RobotStatus", robotStatus, 256)
			}
		}
	}
	return nil
}
