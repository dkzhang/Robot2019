package typeNotificationStructure

import "testing"

func TestUnmarshalJSON(t *testing.T) {
	strJSON := `
	{
    	"type": "notification",
    	"code": "01001",
   	 	"level": "info",
    	"description": "The move task is started.",
    	"data":{"target":"room_205",
			"markers":["m1","m2"],
			"distance_tolerance":1}
	}
	`
	sr, err := UnmarshalJSON(strJSON)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	} else {
		t.Logf("UnmarshalJSON success: %v", sr)
		//t.Logf("type = %s, command = %s, uuid = %s", sr.Type, sr.Command, sr.UUID)
		//t.Logf("status = %s, error_message = %s, task_id = %s", sr.Status, sr.ErrorMessage, sr.TaskID)
	}
}
