package main

import (
	lifter "Robot2019/applicationDriverForRobot/lifterControl/client"
	singleMove "Robot2019/chassisDriverForRobot/robotSinglePointMove/client"

	"Robot2019/dataServer/missionPlanningExecutionSystem/structure"

	"log"
	"strconv"
)

func main() {
	mp := structure.MissionPlanning{}

	//读取任务计划
	mp.UnmarshalJSON("")

	//逐条执行计划
	for i, mission := range mp.Missions {
		log.Printf("Mission %d : %s is about to be executed!", i, mission.MoveMarker)

		singleMove.MoveAndWaitForArrival(mission.MoveMarker)

		for j, sm := range mission.SubMissions {
			log.Printf("SubMission %d : %s is about to be executed!", j, mission.MoveMarker)

			switch sm.Name {
			case "Lifter":
				para, err := strconv.ParseInt(sm.Para, 10, 64)
				if err != nil {
					log.Printf(" fatal error! ParseInt error: %v", err)
				} else {
					lifter.LifterControl(para)
				}
			case "Thermal":
				//调用服务生成图像

				//根据图像名，生成一条记录写入redis数据库

			}

			log.Printf("SubMission %d : %s is accomplished!", j, mission.MoveMarker)
		}

		log.Printf("Mission %d : %s is accomplished!", i, mission.MoveMarker)
	}
}
