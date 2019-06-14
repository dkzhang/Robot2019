package mainProcessor

import (
	"Robot2019/chassisDriverForRobot/mainProcessor/auxiliary"
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"Robot2019/chassisDriverForRobot/typeNotificationProcessor"
	"Robot2019/chassisDriverForRobot/typeNotificationProcessor/typeNotificationStructure"

	"log"
	"sync"
)

type GeneralMainProcessor struct {
	serverIPandPort  string
	commandChan      chan socketCommunication.CommandStruct
	cancelChan       chan interface{}
	notificationProc *typeNotificationProcessor.TypeNotificationProcessor
	socketManagement *socketCommunication.SocketManagement

	runOnce sync.Once
}

func (gmp *GeneralMainProcessor) run() {
	resultChan, feedbackChan := gmp.socketManagement.GetResultAndFeedbackChan()
	notificationOutputChan := gmp.notificationProc.GetInChan()

	//启动socket收发goroutine
	gmp.socketManagement.GoRun()
	//启动Notification分发goroutine
	gmp.GoRun()

	for {
		var resultMsg string
		var feedBack socketCommunication.CommandFeedback

		//四通道接收
		// 1. 接收command通道的新发指令
		// 2. 接收socket management从commandResult通道发来的接收报文
		// 3. 接收socket management从feedback通道发来的反馈消息
		// 4. 接收cancel通道的状态，决定是否退出运行

		select {
		case resultMsg = <-resultChan:
			resultType, err := auxiliary.UnmarshalJSON(resultMsg)
			if err != nil {
				log.Fatalf("auxiliary.UnmarshalJSON <%s> error: %v", resultMsg, err)
				continue
			}

			switch resultType.Type {
			case "notification":
				log.Printf("Receive notification from socket: %v", resultMsg)
				notificationOutputChan <- resultMsg
			case "response":
				// TODO
			case "callback":
				// TODO
			default:
				log.Fatalf("unexpected resultMsg type <%s> : %s", resultType.Type, resultMsg)
			}

		case feedBack = <-feedbackChan:
			log.Printf("Receive feedback from socket management: %v", feedBack)
		case <-gmp.cancelChan:
			return
		}
	}
}

func (gmp *GeneralMainProcessor) RegisterNotificationListener(name string, notificationChan chan typeNotificationStructure.Notification) (err error) {
	return gmp.notificationProc.RegisterListener(name, notificationChan)
}

func (gmp *GeneralMainProcessor) UnregisterNotificationListener(name string) {
	gmp.notificationProc.UnregisterListener(name)
}

func (gmp *GeneralMainProcessor) IsRunning() bool {
	select {
	case <-gmp.cancelChan:
		return false
	default:
		return true
	}
}

func (gmp *GeneralMainProcessor) Cancel() {
	close(gmp.cancelChan)
}

func (gmp *GeneralMainProcessor) GoRun() {
	gmp.runOnce.Do(func() {
		go gmp.run()
	})
}

func GeneralMainProcessorFactory(serverIPandPort string) *GeneralMainProcessor {
	commandChanSize := 1024

	ptr := &GeneralMainProcessor{
		serverIPandPort:  serverIPandPort,
		commandChan:      make(chan socketCommunication.CommandStruct, commandChanSize),
		cancelChan:       make(chan interface{}),
		notificationProc: typeNotificationProcessor.TypeNotificationProcessorFactory(),
	}
	ptr.socketManagement = socketCommunication.SocketManagementFactory(ptr.serverIPandPort, ptr.commandChan)
	return ptr
}
