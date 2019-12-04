package myUtil

import "testing"

func TestSplitJSON(t *testing.T) {
	strJSON := `strJSON = {"command":"/api/request_data","error_message":"","status":"OK","type":"response","uuid":"3CC7D568"}      
  {"code":"01001","data":{"target":"R1x9"},"description":"The move task is started.","level":"info","type":"notification"}, `

	ss := SplitJSON(strJSON)

	if len(ss) != 2 {
		t.Errorf("SplitJSON error: %v", ss)
	} else {
		t.Logf("SplitJSON success: %v", ss)
	}
}
