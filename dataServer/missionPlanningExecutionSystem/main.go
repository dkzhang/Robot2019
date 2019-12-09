package main

import (
	laserLight "Robot2019/applicationDriverForRobot/laserLight/client"
	lifter "Robot2019/applicationDriverForRobot/lifterControl/client"
	singleMove "Robot2019/chassisDriverForRobot/robotSinglePointMove/client"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"Robot2019/dataServer/missionPlanningExecutionSystem/structure"

	"log"
	"strconv"
)

func main() {
	mp := structure.MissionPlanning{}

	//读取任务计划
	f, err := os.OpenFile("./missionPlanning.json", os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		log.Printf("fatal error! Open JSON file error: %v", err)
		return
	} else {
		contentByte, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("fatal error! Read JSON file ioutil.ReadAll error: %v", err)
			return
		}
		strJSON := string(contentByte)
		mp.UnmarshalJSON(strJSON)
	}

	for { //无限循环执行
		//逐条执行计划
		for i, mission := range mp.Missions {
			log.Printf("Main Mission %d : (%s: %s) is about to be executed!",
				i, mission.TheMainMission.Name, mission.TheMainMission.Para)

			err := ExecuteMainMission(mission.TheMainMission)
			if err != nil {
				log.Printf("fatal error! ExecuteMainMission error: %v", err)
				continue
			}

			for j, sm := range mission.TheSubMissions {
				log.Printf("SubMission %d : %s is about to be executed!", j, sm.Name)
				err := ExecuteSubMission(sm)
				if err != nil {
					log.Printf("fatal error! ExecuteSubMission error: %v", err)
					continue
				}
				log.Printf("SubMission %d : %s is accomplished!", j, sm.Name)
			}

			log.Printf("Mission %d : %s is accomplished!", i, mission.TheMainMission.Name)
		}
	}
}

func ExecuteMainMission(mm structure.MainMission) (err error) {
	switch mm.Name {
	case structure.MAIN_MISSION_sMoveWF:
		return singleMove.MoveAndWaitForArrival(mm.Para)
	case structure.MAIN_MISSION_Wait:
		paraInt, err := strconv.ParseInt(mm.Para, 10, 64)
		if err != nil {
			log.Printf(" fatal error! ParseInt error: %v", err)
			return fmt.Errorf("fatal error! MAIN_MISSION_Wait ParseInt error: %s, %v", mm.Para, err)
		} else {
			time.Sleep(time.Second * time.Duration(paraInt))
			return nil
		}
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
	case structure.SUB_MISSION_LaserLight:
		paraBool, err := strconv.ParseBool(sm.Para)
		if err != nil {
			log.Printf(" fatal error! paraBool error: %v", err)
			return fmt.Errorf("fatal error! SUB_MISSION_LaserLight ParseBool error: %s, %v", sm.Para, err)
		} else {
			laserLight.SwitchLaserLight(paraBool)
			return nil
		}
	case structure.SUB_MISSION_Wait:
		paraInt, err := strconv.ParseInt(sm.Para, 10, 64)
		if err != nil {
			log.Printf(" fatal error! ParseInt error: %v", err)
			return fmt.Errorf("fatal error! SUB_MISSION_Wait ParseInt error: %s, %v", sm.Para, err)
		} else {
			time.Sleep(time.Second * time.Duration(paraInt))
			return nil
		}

	default:
		return fmt.Errorf("unsupported Main Mission Name: %s", sm.Name)
	}
	return nil
}
