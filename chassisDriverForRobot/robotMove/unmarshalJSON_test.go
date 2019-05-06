package robotMove

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
	}
}
