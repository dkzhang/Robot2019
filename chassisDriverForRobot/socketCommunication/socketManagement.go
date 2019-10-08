package socketCommunication

import (
	"log"
	"net"
	"sync"
)

// SocketManagement一般情况下会一直运行，即使网络中断也会自动重连
// 设计为插件结构被注入到各个消息处理模块中，因此可以有多个SocketManagement的goroutine存在
type SocketManagement struct {
	ServerIPandPort string
	CommandChan     chan *CommandStruct
	ResultChan      chan string
	FeedbackChan    chan *CommandFeedback

	cancelChan chan interface{}

	runOnce sync.Once
}

var commandChanSize = 1024
var resultChanSize = 1024
var feedbackChanSize = 64

func SocketManagementFactory(serverIPandPort string) *SocketManagement {
	//创建
	psm := &SocketManagement{
		ServerIPandPort: serverIPandPort,
		CommandChan:     make(chan *CommandStruct, commandChanSize),

		ResultChan:   make(chan string, resultChanSize),
		FeedbackChan: make(chan *CommandFeedback, feedbackChanSize),

		cancelChan: make(chan interface{}),
	}

	//一个实例只能运行一次
	psm.runOnce.Do(func() {
		go psm.run()
	})

	return psm
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

func (sm *SocketManagement) run() {
	//SocketManagementRun 管理着两个go routine，分别用于发送和接收
	//发送为主goroutine，接收则另起一个goroutine
	//如果发送或接收时网络出错，则使用errorChan通知另外一个go routine退出
	//当两个go routine都退出时，SocketManagement尝试重新连接
	defer func() {
		close(sm.ResultChan)
		close(sm.FeedbackChan)
	}()

	errorChan := make(chan error, 1)
	var wg sync.WaitGroup

	var pcs *CommandStruct = nil

	for {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", sm.ServerIPandPort)
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
		go SocketReceive(conn, sm.ResultChan, errorChan, &wg)

		//发送不用go routine
		for {
			pcs, err = SocketSend(conn, sm.CommandChan, errorChan, sm.cancelChan, pcs)
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
