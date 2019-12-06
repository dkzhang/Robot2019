package server

import (
	"Robot2019/chassisDriverForRobot/common"
	"encoding/json"
	"fmt"
)

type RobotStatusTopic struct {
	common.CallbackTopic
	Results   Result `json:"results"`
	TimeStamp string
}

type Result struct {
	MoveTarget     string `json:"move_target"`
	MoveStatus     string `json:"move_status"`
	RunningStatus  string `json:"running_status"`
	MoveRetryTimes int    `json:"move_retry_times"`

	ChargeState    bool `json:"charge_state"`
	SoftEStopState bool `json:"soft_estop_state"`
	HardEStopState bool `json:"hard_estop_state"`
	EStopState     bool `json:"estop_state"`
	PowerPercent   int  `json:"power_percent"`

	CurrentPose  RobotPose `json:"current_pose"`
	CurrentFloor int       `json:"current_floor"`
	ErrorCode    string    `json:"error_code"`
}

type RobotPose struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Theta float64 `json:"theta"`
}

func (rst *RobotStatusTopic) UnmarshalJSON(strJSON string) (err error) {
	err = json.Unmarshal([]byte(strJSON), rst)
	if err != nil {
		return fmt.Errorf("robot status topic json.Unmarshal error: %v", err)
	} else {
		return nil
	}
}
