package client

import "testing"

func TestMMove(t *testing.T) {
	markers := []string{"R1x1", "R1x9", "R2x9", "R2x1",
		"R3x1", "R3x9", "R4x9", "R4x1",
		"R5x1", "R5x9", "R6x9", "R7x9",
		"R6x1", "R7x1",
		"charger"}
	MMove(markers)
}
