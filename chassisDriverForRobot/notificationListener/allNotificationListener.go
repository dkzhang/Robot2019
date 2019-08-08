package notificationListener

import (
	"Robot2019/chassisDriverForRobot/generalCommandTransceiver"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"encoding/json"
)

type AllNotificationListener struct {
	Processor  generalCommandTransceiver.GeneralCommandTransceiver
	SocketComm socketCommunication.SocketManagement
}

func (nl *AllNotificationListener) Init(p generalCommandTransceiver.GeneralCommandTransceiver,
	s socketCommunication.SocketManagement) {
	//1. 先配置收发器和socket
	nl.Processor = p
	nl.SocketComm = s

	//2. 然后注册处理模块到收发器
	var reg generalCommandTransceiver.RegisteringStructure

	//无需发送命令
	reg.Command = nil

	//接收type = notification的消息
	reg.Filter = func(strJSON string) bool {
		//先判断消息类型是否为notification
		var resultType generalCommandTransceiver.ResultType
		err := json.Unmarshal([]byte(strJSON), &resultType)
		if err != nil {
			return false
		}
		if resultType.Type != "notification" {
			return false
		}
		return true
	}

	//定义Callback函数
	var NotificationProcessChanSize = 16
	NotificationProcessChan := make(chan Notification, NotificationProcessChanSize)
	reg.Callback = func(strJSON string) (b bool, e error) {
		notification, err := UnmarshalJSON(strJSON)
		if err != nil {
			return false, err
		} else {
			NotificationProcessChan <- notification
			return false, nil
		}
	}

}
