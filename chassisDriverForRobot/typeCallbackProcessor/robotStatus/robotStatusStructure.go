package robotStatus

type CallbackTopic struct {
	Type    string `json:"type"`
	Topic   string `json:"topic"`
	Results Result `json:"results"`
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
	ErrorCode    uint64    `json:"error_code"`
	//`json:" "`
}

type RobotPose struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Theta float64 `json:"theta"`
}
