package main

import (
	"Robot2019/chassisDriverForRobot/robotSinglePointMove/client"
	"Robot2019/dataServer/missionPlanningExecutionSystem/structure"
	"log"
)

func main() {
	mp := structure.MissionPlanning{}

	//读取任务计划
	mp.UnmarshalJSON("")

	//逐条执行计划
	for i, mission := range mp.Missions {
		log.Printf("Mission %d : %s is about to be executed!", i, mission.MoveMarker)

		client.MoveAndWaitForArrival(mission.MoveMarker)

		for j, sm := range mission.SubMissions {
			log.Printf("SubMission %d : %s is about to be executed!", j, mission.MoveMarker)

			switch sm.Name {
			case "Lifter":
				//
			case "Thermal":
				//

			}

			log.Printf("SubMission %d : %s is accomplished!", j, mission.MoveMarker)
		}

		log.Printf("Mission %d : %s is accomplished!", i, mission.MoveMarker)
	}
}
