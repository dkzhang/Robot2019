package robotMove

import (
	"Robot2019/chassisDriverForRobot/generalCommandTransceiver"
	"Robot2019/chassisDriverForRobot/socketCommunication"
)

type SinglePointMove struct {
	Processor  generalCommandTransceiver.GeneralCommandTransceiver
	SocketComm socketCommunication.SocketManagement
}

func (spm *SinglePointMove) Move() {

}

func (spm *SinglePointMove) MoveAndWaitForArrival() {

}
