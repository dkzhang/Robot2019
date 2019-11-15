package main

import (
	robotSinglePointMove "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	"testing"
)

func TestGenerateMoveCommand(t *testing.T) {
	spi := &robotSinglePointMove.SinglePointInfo{
		InfoMask:             15,
		Marker:               "point123xx",
		LocationX:            1.1,
		LocationY:            2.2,
		LocationTheta:        3.3,
		MaxContinuousRetries: 4,
		DistanceTolerance:    5.5,
		ThetaTolerance:       6.6,
	}

	cmd, uuid := GenerateMoveCommand(spi)
	t.Logf("cmd=%s, uuid=%s", cmd, uuid)
}
