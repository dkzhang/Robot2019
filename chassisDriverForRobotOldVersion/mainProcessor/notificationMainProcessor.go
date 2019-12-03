package mainProcessor

import (
	"Robot2019/chassisDriverForRobotOldVersion/mainProcessor/auxiliary"
	"Robot2019/chassisDriverForRobotOldVersion/socketCommunication"
	"Robot2019/chassisDriverForRobotOldVersion/typeNotificationProcessor"
	"Robot2019/chassisDriverForRobotOldVersion/typeNotificationProcessor/typeNotificationStructure"
	"log"
	"sync"
)

type NotificationMainProcessor struct {
	serverIPandPort  string
	commandChan      chan socketCommunication.CommandStruct
	cancelChan       chan interface{}
	notificationProc *typeNotificationProcessor.TypeNotificationProcessor
	socketManagement *socketCommunication.SocketManagement

	runOnce sync.Once
}

func (nmp *NotificationMainProcessor) run() {
	resultChan, _ := nmp.socketManagement.GetResultAndFeedbackChan()
	notificationOutputChan := nmp.notificationProc.GetInChan()

	//启动socket收发goroutine
	nmp.socketManagement.GoRun()
	//启动notification分发goroutine
	nmp.GoRun()

	for {
		var resultMsg string
		select {
		case resultMsg = <-resultChan:

		case <-nmp.cancelChan:
			return
		}
		resultType, err := auxiliary.UnmarshalJSON(resultMsg)

		if err != nil {
			log.Printf(" fatal error! auxiliary.UnmarshalJSON <%s> error: %v", resultMsg, err)
			continue
		}

		if resultType.Type == "notification" {
			notificationOutputChan <- resultMsg
		} else {
			log.Printf(" fatal error! unexpected resultMsg type <%s> : %s", resultType.Type, resultMsg)
		}
	}

}

func (nmp *NotificationMainProcessor) RegisterNotificationListener(name string, notificationChan chan typeNotificationStructure.Notification) (err error) {
	return nmp.notificationProc.RegisterListener(name, notificationChan)
}

func (nmp *NotificationMainProcessor) UnregisterNotificationListener(name string) {
	nmp.notificationProc.UnregisterListener(name)
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
	commandChanSize := 1

	ptr := &NotificationMainProcessor{
		serverIPandPort:  serverIPandPort,
		commandChan:      make(chan socketCommunication.CommandStruct, commandChanSize),
		cancelChan:       make(chan interface{}),
		notificationProc: typeNotificationProcessor.TypeNotificationProcessorFactory(),
	}
	ptr.socketManagement = socketCommunication.SocketManagementFactory(ptr.serverIPandPort, ptr.commandChan)
	return ptr
}
