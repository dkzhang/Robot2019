package mainProcessor

import (
	"Robot2019/chassisDriverForRobot/socketCommunication"
	"sync"
	"time"
)

type MainProcessor struct {
	//robotStatusListenerMap map[string](chan robotStatus.CallbackTopic)
	//inChan                 chan string
	serverIP string

	commandChan chan socketCommunication.CommandStruct
	resultChan  chan socketCommunication.CommandResultStruct
	cancelChan  chan interface{}
}

func (mp *MainProcessor) run() {
	go socketCommunication.SocketManagementRun(mp.serverIP, mp.commandChan, mp.resultChan, mp.cancelChan)

	//go PrintResultChan(resultChan)
}

var ptr *MainProcessor = nil
var once sync.Once
var timeout = 10 * time.Second
var chanSize int = 1024

func TypeCallbackProcessorFactory() *MainProcessor {
	once.Do(func() {
		ptr = &MainProcessor{
			//robotStatusListenerMap: make(map[string](chan robotStatus.CallbackTopic)),
			//inChan:                 make(chan string, inChanSize),
			serverIP:    "192.168.10.10:31001",
			commandChan: make(chan socketCommunication.CommandStruct, chanSize),
			resultChan:  make(chan socketCommunication.CommandResultStruct, chanSize),
			cancelChan:  make(chan interface{}),
		}
		go ptr.run()
	})
	return ptr
}
