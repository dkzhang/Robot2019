package structure

import (
	"encoding/json"
	"fmt"
)

// 任务集
type MissionPlanning struct {
	Missions []Mission `json:"missions"`
}

// 前往指定位置，执行指定的一系列子任务（如果有）
type Mission struct {
	MoveMarker  string       `json:"move_marker"`
	SubMissions []SubMission `json:"submissions"`
}

// 机器人到达指定位置后，执行的子任务
// 目前有：升降杆上升/下降，获取热红外影像，获取温湿度值
type SubMission struct {
	Name string `json:"name"`
	Para string `json:"para"`
}

func (mp *MissionPlanning) UnmarshalJSON(strJSON string) (err error) {
	err = json.Unmarshal([]byte(strJSON), mp)
	if err != nil {
		return fmt.Errorf("MissionPlanning json.Unmarshal error: %v", err)
	} else {
		return nil
	}
}
