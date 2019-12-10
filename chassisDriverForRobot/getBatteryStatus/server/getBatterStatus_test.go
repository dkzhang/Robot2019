package server

import "testing"

func TestGetBatteryStatus(t *testing.T) {
	_, err := GetBatteryStatus(nil, nil)
	if err != nil {
		t.Errorf("GetBatteryStatus error: %v", err)
	}
}
