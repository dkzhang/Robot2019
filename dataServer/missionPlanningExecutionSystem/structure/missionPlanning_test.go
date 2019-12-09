package structure

import (
	"os"
	"testing"
)

func TestMissionPlanning_MarshalJSON(t *testing.T) {
	mp := MissionPlanning{}

	mission1 := Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_Move, Para: "R1x1"},
		TheSubMissions: []SubMission{},
	}
	mp.Missions = append(mp.Missions, mission1)

	mission2 := Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_Move, Para: "R1x9"},
		TheSubMissions: []SubMission{},
	}
	mp.Missions = append(mp.Missions, mission2)

	mission3 := Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_Move, Para: "charger"},
		TheSubMissions: []SubMission{},
	}
	mp.Missions = append(mp.Missions, mission3)

	strJSON, err := mp.MarshalJSON()
	if err != nil {
		t.Fatalf("mp.MarshalJSON error: %v", err)
	} else {
		t.Logf("mp.MarshalJSON success: %s", strJSON)
	}

	f, err := os.Create("missionPlanning.json")
	defer f.Close()
	if err != nil {
		t.Errorf("os.Create file error: %v", err)
	} else {
		_, err = f.Write([]byte(strJSON))
		if err != nil {
			t.Errorf("file write error: %v", err)
		}
	}
}
