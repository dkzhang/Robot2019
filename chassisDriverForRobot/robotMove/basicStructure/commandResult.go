package basicStructure

import "Robot2019/chassisDriverForRobot/generalCommandTransceiver"

type CommandProcessResult struct {
	generalCommandTransceiver.BasicCommandResult
	TaskID string `json:"task_id"`
}

var CommandResultProcessChanSize = 16

func GenerateCallbackFunc() (func(*generalCommandTransceiver.CommandResult) error, chan CommandProcessResult) {
	commandResultProcessChan := make(chan CommandProcessResult, CommandResultProcessChanSize)

	cbf := func(cr *generalCommandTransceiver.CommandResult) error {
		commandResult, err := UnmarshalJSON(cr.StrJSON)
		if err != nil {
			return err
		} else {
			commandResultProcessChan <- commandResult
			return nil
		}
	}

	return cbf, commandResultProcessChan
}
