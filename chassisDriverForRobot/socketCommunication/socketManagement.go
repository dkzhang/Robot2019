package socketCommunication

import (
	"log"
	"net"
	"sync"
)

// SocketManagement一般情况下会一直运行，即使网络中断也会自动重连
// 因此设计为Run Once，如果结束运行，则该实例无法再次运行
type SocketManagement struct {
	serverIP    string
	commandChan chan CommandStruct
	resultChan  chan CommandResultStruct

	feedbackChan chan CommandFeedback
	cancelChan   chan interface{}

	runOnce sync.Once
}

var resultChanSize = 1024
var feedbackChanSize = 64

func SocketManagementFactory(serverIP string, commandChan chan CommandStruct) *SocketManagement {
	ptr := &SocketManagement{
		serverIP:    serverIP,
		commandChan: commandChan,

		resultChan:   make(chan CommandResultStruct, resultChanSize),
		feedbackChan: make(chan CommandFeedback, feedbackChanSize),

		cancelChan: make(chan interface{}),
	}
	return ptr
}

func (sm *SocketManagement) IsRunning() bool {
	select {
	case <-sm.cancelChan:
		return false
	default:
		return true
	}
}

func (sm *SocketManagement) Cancel() {
	close(sm.cancelChan)
}

func (sm *SocketManagement) GetResultAndFeedbackChan() (chan CommandResultStruct, chan CommandFeedback) {
	return sm.resultChan, sm.feedbackChan
}

func (sm *SocketManagement) GoRun() {
	sm.runOnce.Do(func() {
		go sm.run()
	})
}

func (sm *SocketManagement) run() {
	//SocketManagementRun 管理着两个go routine，分别用于发送和接收
	//发送为主goroutine，接收则另起一个goroutine
	//如果发送或接收时网络出错，则使用errorChan通知另外一个go routine退出
	//当两个go routine都退出时，SocketManagement尝试重新连接
	defer func() {
		close(sm.resultChan)
		close(sm.feedbackChan)
	}()

	errorChan := make(chan error, 1)
	var wg sync.WaitGroup

	var pcs *CommandStruct = nil

	for {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", sm.serverIP)
		if err != nil {
			log.Fatalf("net.ResolveTCPAddr error: %v", err)
			continue
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			log.Fatalf("net.DialTCP error: %v", err)
			continue
		}

		wg.Add(1)
		go SocketReceive(conn, sm.resultChan, errorChan, &wg)

		//发送不用go routine
		for {
			pcs, err = SocketSend(conn, sm.commandChan, errorChan, sm.cancelChan, pcs)
			if err != nil {
				log.Fatalf("SocketSend error: %v", err)
				if errClose := conn.Close(); errClose != nil {
					log.Fatalf("conn.Close() in SocketSend error: %v", errClose)
				}

				if err.Error() == "cancel" {
					return
				} else {
					break
				}

			}
		}
		wg.Wait()
	}
}
