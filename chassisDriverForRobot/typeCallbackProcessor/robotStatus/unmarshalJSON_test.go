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
        	"charge_state": true, 
        	"soft_estop_state": true, 
        	"hard_estop_state": false, 
        	"estop_state": false,  
        	"power_percent": 100, 
        	"current_pose": {
            	"x": 11.0,     
           	 	"y": 11.0,      
            	"theta": 0.5
        	},
			"current_floor": 16,
			"error_code": "00000000"   
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
