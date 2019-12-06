package client

import "testing"

func TestGetRobotStatus(t *testing.T) {
	reply, err := GetRobotStatus()
	if err != nil {
		t.Errorf("GetRobotStatus error: %v", err)
	} else {
		t.Logf("GetRobotStatus reply = %v", *reply)
	}
}
