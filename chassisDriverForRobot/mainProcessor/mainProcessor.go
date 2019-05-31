package mainProcessor

import (
	"sync"
	"time"
)

type MainProcessor struct {
	//robotStatusListenerMap map[string](chan robotStatus.CallbackTopic)
	//inChan                 chan string
}

func (mp *MainProcessor) run() {

}

var ptr *MainProcessor = nil
var once sync.Once
var inChanSize = 1024
var timeout = 10 * time.Second

func TypeCallbackProcessorFactory() *MainProcessor {
	once.Do(func() {
		ptr = &MainProcessor{
			//robotStatusListenerMap: make(map[string](chan robotStatus.CallbackTopic)),
			//inChan:                 make(chan string, inChanSize),
		}
		go ptr.run()
	})
	return ptr
}
