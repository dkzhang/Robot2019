package client

import "testing"

func TestMMove(t *testing.T) {
	markers := []string{"R1x1", "R1x2", "charger"}
	MMove(markers)
}
