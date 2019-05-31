package robotStatus

import "testing"

func TestUnmarshalJSON(t *testing.T) {
	strJSON := `
	{
    "type": "callback",
    "topic": "robot_status",    
    "results": {
        "move_target": "target_name", 
        "move_status": "running", 
        "running_status": "running", 
        "move_retry_times": 3, 

        "charge_state": bool, 
        "soft_estop_state": bool, // 通过API接口设置的软急停状态, true->急停中，false->非急停中
        "hard_estop_state": bool, // 通过硬件急停按钮设置的硬急停状态, true->急停中，false->非急停中
        "estop_state": bool,  // hard_estop_state || sofpt_estop_state, true->急停中，false->非急停中
        "power_percent": 100, //电量百分比，单位：%
        "current_pose": {
            "x": 11.0,     // 单位：m
            "y": 11.0,       // 单位：m
            "theta": 0.5, //单位：rad
        }
        "current_floor": 16,
        "error_code": "00000000"   // v0.7.7新增，16进制错误码，总共8个字节表示，非0表示机器人异常
    }
}
	`
	sr, err := UnmarshalJSON(strJSON)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	} else {
		t.Logf("UnmarshalJSON success: %v", sr)
		t.Logf("type = %s, topic = %s", sr.Type, sr.Topic)
		//t.Logf("status = %s, error_message = %s, task_id = %s", sr.Status, sr.ErrorMessage, sr.TaskID)
	}
}
