package main

import (
	"fmt"
	"math/rand"
	"time"

	pb "Robot2019/chassisDriverForRobot/robotMultiplePointsMove/grpc"
)

func GenerateMoveCommand(mpi *pb.MultiplePointsInfo) (cmd string, uuid string) {
	cmd = "/api/move?"

	cmd += mpi.Markers[0]

	for i := 1; i < len(mpi.Markers); i++ {
		cmd += "," + mpi.Markers[i]
	}

	// count
	if mpi.InfoMask&4 != 0 {
		cmd += fmt.Sprintf("&count=%d", mpi.Count)
	}

	// distance_tolerance
	if mpi.InfoMask&2 != 0 {
		cmd += fmt.Sprintf("&distance_tolerance=%f", mpi.DistanceTolerance)
	}

	// max_continuous_retries
	if mpi.InfoMask&1 != 0 {
		cmd += fmt.Sprintf("&max_continuous_retries=%d", mpi.MaxContinuousRetries)
	}

	rand.Seed(time.Now().Unix())
	uuid = fmt.Sprintf("%X", rand.Uint32())
	cmd += fmt.Sprintf("&uuid=%s", uuid)

	return cmd, uuid
}
