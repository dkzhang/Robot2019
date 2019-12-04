package server

import (
	"Robot2019/chassisDriverForRobot/common"
	"encoding/json"
	"fmt"
)

type SingleMoveCommandResponse struct {
	common.CommandResponse
	TaskID string `json:"task_id"`
}

func (smcr *SingleMoveCommandResponse) UnmarshalJSON(strJSON string) (err error) {
	err = json.Unmarshal([]byte(strJSON), smcr)
	if err != nil {
		return fmt.Errorf("SingleMoveCommandResponse json.Unmarshal error: %v", err)
	} else {
		return nil
	}
}
