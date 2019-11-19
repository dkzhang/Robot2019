package main

import (
	"Robot2019/chassisDriverForRobot/common"
	"encoding/json"
	"fmt"
)

type MultipleMoveCommandResponse struct {
	common.CommandResponse
	TaskID string `json:"task_id"`
}

func (mmcr *MultipleMoveCommandResponse) UnmarshalJSON(strJSON string) (err error) {
	err = json.Unmarshal([]byte(strJSON), mmcr)
	if err != nil {
		return fmt.Errorf("MultipleMoveCommandResponse json.Unmarshal error: %v", err)
	} else {
		return nil
	}
}
