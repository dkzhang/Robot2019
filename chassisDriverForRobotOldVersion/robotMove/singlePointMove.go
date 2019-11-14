package robotMove

import (
	"Robot2019/chassisDriverForRobotOldVersion/generalCommandTransceiver"
	"Robot2019/chassisDriverForRobotOldVersion/socketCommunication"
)

type SinglePointMove struct {
	Processor  generalCommandTransceiver.GeneralCommandTransceiver
	SocketComm socketCommunication.SocketManagement
}

func (spm *SinglePointMove) Move() {

}

func (spm *SinglePointMove) MoveAndWaitForArrival() {

}
