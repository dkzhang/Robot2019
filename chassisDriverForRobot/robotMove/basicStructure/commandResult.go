package basicStructure

import "Robot2019/chassisDriverForRobot/generalCommandTransceiver"

type CommandResult struct {
	generalCommandTransceiver.BasicCommandResult
	TaskID string `json:"task_id"`
}

var CommandResultChanSize = 16

func GenerateCallbackFunc() (func(*generalCommandTransceiver.CommandResult) error, chan CommandResult) {
	commandResultChan := make(chan CommandResult, CommandResultChanSize)

	cbf := func(cr *generalCommandTransceiver.CommandResult) error {
		commandResult, err := UnmarshalJSON(cr.StrJSON)
		if err != nil {
			return err
		} else {
			commandResultChan <- commandResult
			return nil
		}
	}

	return cbf, commandResultChan
}
