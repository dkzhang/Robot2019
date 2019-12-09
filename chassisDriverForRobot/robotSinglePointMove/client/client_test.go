package client

import "testing"

func TestMoveAndWaitForArrival(t *testing.T) {
	err := MoveAndWaitForArrival("R7x1")
	if err != nil {
		t.Errorf("MoveAndWaitForArrival error: %v", err)
	}
}
