package generalCommandTransceiver

import (
	"Robot2019/chassisDriverForRobotOldVersion/socketCommunication"
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

var CommandResultChanSize = 16

func (gct *GeneralCommandTransceiver) run() {

	for {
		var resultMsg string
		var feedBack socketCommunication.CommandFeedback
		//var registerCommand *RegisteringStructure

		//四通道接收
		// 1. 接收command通道的新发指令
		// 2. 接收socket management从commandResult通道发来的接收报文
		// 3. 接收socket management从feedback通道发来的反馈消息
		// 4. 接收cancel通道的状态，决定是否退出运行

		select {
		case registerCommand := <-gct.CommandChan:
			//登记, 构造消息接收chan
			gct.commandArray = append(gct.commandArray, registerCommand)
			registerCommand.resultChan = make(chan string, CommandResultChanSize)

			//发送指令
			if registerCommand.Command != nil {
				gct.socketManagement.CommandChan <- registerCommand.Command
			}

			//等待接收消息
			go func(removeCommandChan chan *RegisteringStructure, cmd *RegisteringStructure) {
				cmdResult, ok := <-cmd.resultChan
				if !ok {
					//如果外界主动remove注册，则会关闭resultChan，从而退出等待接收消息
					return
				}
				if cmd.Filter(cmdResult) == true {
					finished, err := cmd.Callback(cmdResult)
					if err != nil {
						log.Printf("callback function error: %v", err)
					}
					if finished == true {
						//如果Callback返回true，则表示全套消息处理流程已完成且不再处理消息，主动取消注册
						removeCommandChan <- cmd
						return
					}
				}
			}(gct.removeCommandChan, registerCommand)
			//由此可见，已注册消息处理模块退出消息处理有两种方式：
			//1. Callback中判断走完流程，返回true，系统自动remove注册
			//2. 外部不想继续处理，主动remove注册

		case removeCommand := <-gct.removeCommandChan:
			//关闭channel
			close(removeCommand.resultChan)
			//从队列中移除
			for i := 0; i < len(gct.commandArray); i++ {
				if gct.commandArray[i] == removeCommand {
					//删除该cmd
					//不维持剩余元素顺序，简单地将最后一个元素赋值给移除元素所在位置
					gct.commandArray[i] = gct.commandArray[len(gct.commandArray)]
					gct.commandArray = gct.commandArray[:len(gct.commandArray)-1]
				}
			}

		case resultMsg = <-gct.socketManagement.ResultChan:
			//将消息转发给已注册的各模块的channel
			//暂时未考虑队列满且消息处理超时
			for _, cmds := range gct.commandArray {
				cmds.resultChan <- resultMsg
			}

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
		commandArray:     make([]*RegisteringStructure, 0, 0),
		cancelChan:       make(chan interface{}),
		runOnce:          sync.Once{},
	}

	//每个实例只运行一次
	gct.runOnce.Do(func() {
		go gct.run()
	})

	return gct
}
