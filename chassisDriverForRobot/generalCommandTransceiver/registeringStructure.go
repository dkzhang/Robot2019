package generalCommandTransceiver

import (
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"encoding/json"
)

type RegisteringStructure struct {
	Command  *socketCommunication.CommandStruct
	Filter   func(string) (bool, *CommandResult)
	Callback func(*CommandResult) (bool, error)
}

func DefaultCommandFilter(uuid string) func(string) (bool, *CommandResult) {
	return func(strJSON string) (bool, *CommandResult) {
		//先判断消息类型是否为response
		var resultType ResultType
		err := json.Unmarshal([]byte(strJSON), &resultType)
		if err != nil {
			return false, nil
		}
		if resultType.Type != "response" {
			return false, nil
		}

		//再判断uuid是否符合
		var commandResult CommandResult
		err = json.Unmarshal([]byte(strJSON), &commandResult)
		if err != nil {
			return false, nil
		} else if commandResult.BasicInfo.UUID != uuid {
			return false, &commandResult
		} else {
			return true, &commandResult
		}
	}
}
