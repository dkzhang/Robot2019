package mainProcessor

import (
	"Robot2019/chassisDriverForRobot/mainProcessor/auxiliary"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"Robot2019/chassisDriverForRobot/typeNotificationProcessor"
	"Robot2019/chassisDriverForRobot/typeNotificationProcessor/typeNotificationStructure"
	"log"
	"sync"
)

type NotificationMainProcessor struct {
	serverIPandPort string

	cancelChan       chan interface{}
	notificationProc *typeNotificationProcessor.TypeNotificationProcessor

	runOnce sync.Once
}

var commandChanSize = 1

func (nmp *NotificationMainProcessor) run() {

	var commandChan chan socketCommunication.CommandStruct = make(chan socketCommunication.CommandStruct, commandChanSize)

	psm := socketCommunication.SocketManagementFactory(nmp.serverIPandPort, commandChan)
	resultChan, _ := psm.GetResultAndFeedbackChan()
	notificationOutputChan := nmp.notificationProc.GetInChan()

	for {
		var resultMsg string
		select {
		case resultMsg = <-resultChan:

		case <-nmp.cancelChan:
			return
		}
		resultType, err := auxiliary.UnmarshalJSON(resultMsg)

		if err != nil {
			log.Fatalf("auxiliary.UnmarshalJSON <%s> error: %v", resultMsg, err)
			continue
		}

		if resultType.Type == "notification" {
			notificationOutputChan <- resultMsg
		} else {
			log.Fatalf("unexpected resultMsg type <%s> : %s", resultType.Type, resultMsg)
		}
	}

}

func (nmp *NotificationMainProcessor) RegisterRobotStatusListener(name string, notificationChan chan typeNotificationStructure.Notification) (err error) {
	return nmp.notificationProc.RegisterRobotStatusListener(name, notificationChan)
}

func (nmp *NotificationMainProcessor) UnregisterRobotStatusListener(name string) {
	nmp.notificationProc.UnregisterRobotStatusListener(name)
}

func (nmp *NotificationMainProcessor) IsRunning() bool {
	select {
	case <-nmp.cancelChan:
		return false
	default:
		return true
	}
}

func (nmp *NotificationMainProcessor) Cancel() {
	close(nmp.cancelChan)
}

func (nmp *NotificationMainProcessor) GoRun() {
	nmp.runOnce.Do(func() {
		go nmp.run()
	})
}

func notificationMainProcessorFactory(serverIPandPort string) *NotificationMainProcessor {
	ptr := &NotificationMainProcessor{
		serverIPandPort:  serverIPandPort,
		cancelChan:       make(chan interface{}),
		notificationProc: typeNotificationProcessor.TypeNotificationProcessorFactory(),
	}
	return ptr
}
