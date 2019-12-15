package structure

import (
	"os"
	"testing"
)

func TestMissionPlanning_MarshalJSON(t *testing.T) {
	mp := MissionPlanning{}

	mp.Missions = append(mp.Missions, Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_sMoveWF, Para: "R1x1"},
		TheSubMissions: []SubMission{},
	})

	mp.Missions = append(mp.Missions, Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_sMoveWF, Para: "R1x9"},
		TheSubMissions: []SubMission{},
	})

	mp.Missions = append(mp.Missions, Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_sMoveWF, Para: "R2x9"},
		TheSubMissions: []SubMission{},
	})

	mp.Missions = append(mp.Missions, Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_sMoveWF, Para: "R2x2"},
		TheSubMissions: []SubMission{
			{
				Name: SUB_MISSION_LaserLight,
				Para: "true",
			},
			{
				Name: SUB_MISSION_LifterControl,
				Para: "15000",
			},
			{
				Name: SUB_MISSION_Wait,
				Para: "3",
			},
			{
				Name: SUB_MISSION_LifterControl,
				Para: "-16000",
			},
			{
				Name: SUB_MISSION_LaserLight,
				Para: "false",
			},
		},
	})

	mp.Missions = append(mp.Missions, Mission{
		TheMainMission: MainMission{Name: MAIN_MISSION_sMoveWF, Para: "charger"},
		TheSubMissions: []SubMission{},
	})

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
