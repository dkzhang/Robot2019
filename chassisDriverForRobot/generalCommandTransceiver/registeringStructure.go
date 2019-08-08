package generalCommandTransceiver

import (
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"encoding/json"
)

type RegisteringStructure struct {
	Command    *socketCommunication.CommandStruct
	Filter     func(string) bool
	Callback   func(string) (bool, error)
	resultChan chan string
}

func DefaultCommandFilter(uuid string) func(string) bool {
	return func(strJSON string) bool {
		//先判断消息类型是否为response
		var resultType ResultType
		err := json.Unmarshal([]byte(strJSON), &resultType)
		if err != nil {
			return false
		}
		if resultType.Type != "response" {
			return false
		}

		//再判断uuid是否符合
		var commandResult CommandResult
		err = json.Unmarshal([]byte(strJSON), &commandResult)
		if err != nil {
			return false
		} else if commandResult.BasicInfo.UUID != uuid {
			return false
		} else {
			return true
		}
	}
}
