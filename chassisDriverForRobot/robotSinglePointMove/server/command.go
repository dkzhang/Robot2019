package main

import (
	"fmt"
	"math/rand"
	"time"

	pb "Robot2019/chassisDriverForRobot/robotSinglePointMove/grpc"
)

func GenerateMoveCommand(spi *pb.SinglePointInfo) (cmd string, uuid string) {
	cmd = "/api/move?"

	if spi.InfoMask&16 != 0 {
		//marker
		cmd += "marker=" + spi.Marker
	} else {
		//location
		cmd += fmt.Sprintf("location=%f,%f,%f", spi.LocationX, spi.LocationY, spi.LocationTheta)
	}

	// max_continuous_retries
	if spi.InfoMask&4 != 0 {
		cmd += fmt.Sprintf("&max_continuous_retries=%d", spi.MaxContinuousRetries)
	}

	// distance_tolerance
	if spi.InfoMask&2 != 0 {
		cmd += fmt.Sprintf("&distance_tolerance=%f", spi.DistanceTolerance)
	}

	// theta_tolerance
	if spi.InfoMask&1 != 0 {
		cmd += fmt.Sprintf("&theta_tolerance=%f", spi.ThetaTolerance)
	}

	rand.Seed(time.Now().Unix())
	uuid = fmt.Sprintf("%X", rand.Uint32())
	cmd += fmt.Sprintf("&uuid=%s", uuid)

	return cmd, uuid
}
