package client

import "testing"

func TestMMove(t *testing.T) {
	markers := []string{"R1x1", "R6x1", "charger"}
	MMove(markers)
}
