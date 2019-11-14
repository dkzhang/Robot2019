package basicStructure

import "testing"

func TestUnmarshalJSON(t *testing.T) {
	strJSON := `
	{
		"type":"response",
		"command":"/api/move",
		"uuid":"12345",
		"status":"OK",
		"error_message":"",
		"task_id":"xxx"
	}
	`
	sr, err := UnmarshalJSON(strJSON)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	} else {
		t.Logf("UnmarshalJSON success: %v", sr)
		t.Logf("type = %s, command = %s, uuid = %s", sr.Type, sr.Command, sr.UUID)
		t.Logf("status = %s, error_message = %s, task_id = %s", sr.Status, sr.ErrorMessage, sr.TaskID)
	}
}
