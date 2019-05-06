package robotMove

import "Robot2019/chassisDriverForRobot"

type StructReturn struct {
	chassisDriverForRobot.BasicStructReturn
	TaskID string `json:"task_id"`
}
