package server

import (
	robotSinglePointMove "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
	"testing"
)

func TestGenerateMoveCommand(t *testing.T) {
	spi := &robotSinglePointMove.SinglePointInfo{
		InfoMask:             16,
		Marker:               "R1x1",
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

func TestServer_Move(t *testing.T) {
	s := Server{}

	spi := &robotSinglePointMove.SinglePointInfo{
		InfoMask: 16,
		Marker:   "charger",
		//Marker:               "R7x1",
		LocationX:            1.1,
		LocationY:            2.2,
		LocationTheta:        3.3,
		MaxContinuousRetries: 4,
		DistanceTolerance:    5.5,
		ThetaTolerance:       6.6,
	}

	r, err := s.MoveAndWaitForArrival(nil, spi)
	if err != nil {
		t.Errorf("MoveAndWaitForArrival error: %v", err)
	} else {
		t.Logf("MoveAndWaitForArrival response = %v", *r)
	}
}
