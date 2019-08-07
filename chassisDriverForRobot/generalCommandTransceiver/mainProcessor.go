package generalCommandTransceiver

import (
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"log"
	"sync"
)

type GeneralCommandTransceiver struct {
	socketManagement *socketCommunication.SocketManagement

	CommandChan       chan *RegisteringStructure
	commandArray      []*RegisteringStructure
	removeCommandChan chan *RegisteringStructure

	cancelChan chan interface{}

	runOnce sync.Once
}

func (gct *GeneralCommandTransceiver) IsRunning() bool {
	select {
	case <-gct.cancelChan:
		return false
	default:
		return true
	}
}

func (gct *GeneralCommandTransceiver) Cancel() {
	close(gct.cancelChan)
}

func (gct *GeneralCommandTransceiver) GoRun() {
	gct.runOnce.Do(func() {
		go gct.run()
	})
}

func (gct *GeneralCommandTransceiver) run() {

	for {
		var resultMsg string
		var feedBack socketCommunication.CommandFeedback
		var registerCommand *RegisteringStructure

		//四通道接收
		// 1. 接收command通道的新发指令
		// 2. 接收socket management从commandResult通道发来的接收报文
		// 3. 接收socket management从feedback通道发来的反馈消息
		// 4. 接收cancel通道的状态，决定是否退出运行

		select {
		case registerCommand = <-gct.CommandChan:
			//登记
			gct.commandArray = append(gct.commandArray, registerCommand)

			//发送指令
			if registerCommand.Command != nil {
				gct.socketManagement.CommandChan <- registerCommand.Command
			}

			//等待接收消息

		case resultMsg = <-gct.socketManagement.ResultChan:

		case feedBack = <-gct.socketManagement.FeedbackChan:
			log.Printf("Receive feedback from socket management: %v", feedBack)
		case <-gct.cancelChan:
			return
		}
	}
}

func GeneralMainProcessorFactory(psm *socketCommunication.SocketManagement) *GeneralCommandTransceiver {
	commandChanSize := 1024

	//创建
	gct := &GeneralCommandTransceiver{
		socketManagement: psm,
		CommandChan:      make(chan *RegisteringStructure, commandChanSize),
		commandArray:     make([]RegisteringStructure, 0, 0),
		cancelChan:       make(chan interface{}),
		runOnce:          sync.Once{},
	}

	//每个实例只运行一次
	gct.runOnce.Do(func() {
		go gct.run()
	})

	return gct
}
