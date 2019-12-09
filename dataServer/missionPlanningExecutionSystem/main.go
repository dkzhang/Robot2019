package main

import (
	lifter "Robot2019/applicationDriverForRobot/lifterControl/client"
	singleMove "Robot2019/chassisDriverForRobot/robotSinglePointMove/client"
	"fmt"

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
		log.Printf("Main Mission %d : (%s: %s) is about to be executed!",
			i, mission.TheMainMission.Name, mission.TheMainMission.Para)

		ExecuteMainMission(mission.TheMainMission)

		for j, sm := range mission.TheSubMissions {
			log.Printf("SubMission %d : %s is about to be executed!", j, sm.Name)

			switch sm.Name {
			case structure.MISSION_LifterControl:
				para, err := strconv.ParseInt(sm.Para, 10, 64)
				if err != nil {
					log.Printf(" fatal error! ParseInt error: %v", err)
				} else {
					lifter.LifterControl(para)
				}
			case structure.MISSION_ThermalImaging:
				//调用服务生成图像

				//根据图像名，生成一条记录写入redis数据库

			}

			log.Printf("SubMission %d : %s is accomplished!", j, mission.MoveMarker)
		}

		log.Printf("Mission %d : %s is accomplished!", i, mission.MoveMarker)
	}
}

func ExecuteMainMission(mm structure.MainMission) (err error) {
	switch mm.Name {
	case structure.MAIN_MISSION_sMoveWF:
		return singleMove.MoveAndWaitForArrival(mm.Para)
	default:
		return fmt.Errorf("unsupported Main Mission Name: %s", mm.Name)
	}
	return nil
}

func ExecuteSubMission(sm structure.SubMission) (err error) {
	switch sm.Name {
	case structure.SUB_MISSION_LifterControl:
		paraInt, err := strconv.ParseInt(sm.Para, 10, 64)
		if err != nil {
			log.Printf(" fatal error! ParseInt error: %v", err)
			return fmt.Errorf("fatal error! SUB_MISSION_LifterControl ParseInt error: %s, %v", sm.Para, err)
		} else {
			return lifter.LifterControl(paraInt)
		}
	case structure.SUB_MISSION_ThermalImaging:
		//调用服务生成图像

		//根据图像名，生成一条记录写入redis数据库
	default:
		return fmt.Errorf("unsupported Main Mission Name: %s", sm.Name)
	}
	return nil
}
