package structure

import (
	"encoding/json"
	"fmt"
)

// 任务集
type MissionPlanning struct {
	Missions []Mission `json:"missions"`
}

// 每个任务由一个主任务和一组子任务组成。
// 主任务有且只有一个，子任务有0个到多个。
type Mission struct {
	TheMainMission MainMission  `json:"mainMissions"`
	TheSubMissions []SubMission `json:"submissions"`
}

// 主任务
type MainMission struct {
	Name string `json:"name"`
	Para string `json:"para"`
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

func (mp *MissionPlanning) MarshalJSON() (strJSON string, err error) {
	data, err := json.MarshalIndent(mp, "", "    ")
	if err != nil {
		return "", fmt.Errorf("MissionPlanning MarshalJSON error: %v", err)
	} else {
		strJSON = string(data)
		return strJSON, nil
	}
}
