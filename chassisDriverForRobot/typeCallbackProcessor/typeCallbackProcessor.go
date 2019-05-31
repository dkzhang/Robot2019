package typeCallbackProcessor

import (
	"Robot2019/chassisDriverForRobot/typeCallbackProcessor/robotStatus"
	"Robot2019/chassisDriverForRobot/typeCallbackProcessor/typeCallbackStructure"
	"fmt"
	"log"
	"sync"
	"time"
)

type TypeCallbackProcessor struct {
	robotStatusListenerMap map[string](chan robotStatus.CallbackTopic)
	inChan                 chan string
}

func (proc *TypeCallbackProcessor) RegisterRobotStatusListener(name string, topicChan chan robotStatus.CallbackTopic) (err error) {
	if _, ok := proc.robotStatusListenerMap[name]; !ok {
		proc.robotStatusListenerMap[name] = topicChan
		return nil
	} else {
		return fmt.Errorf("RobotStatus Listener %s already registered!", name)
	}
}

func (proc *TypeCallbackProcessor) UnregisterRobotStatusListener(name string) {
	delete(proc.robotStatusListenerMap, name)
}

func (proc *TypeCallbackProcessor) GetInChan() chan string {
	return proc.inChan
}

func (proc *TypeCallbackProcessor) run() {
	for {
		topicMessage := <-proc.inChan
		log.Printf("Recive callback topic: %s", topicMessage)

		ct, err := typeCallbackStructure.UnmarshalJSON(topicMessage)
		if err != nil {
			log.Fatalf("UnmarshalJSON Callback Topic error: %v", err)
			continue
		} else {
			log.Printf("UnmarshalJSON Callback Topic success: %v", ct)
		}

		if ct.Topic == "robot_status" {
			ctrs, err := robotStatus.UnmarshalJSON(topicMessage)
			if err != nil {
				log.Fatalf("UnmarshalJSON Callback Topic RobotStatus error: %v", err)
				continue
			} else {
				log.Printf("UnmarshalJSON Callback Topic RobotStatus success: %v", ctrs)

				//
				for name, theChan := range proc.robotStatusListenerMap {
					select {
					case theChan <- ctrs:
						//Do nothing
					case <-time.After(timeout):
						log.Fatalf("Send Callback Topic RobotStatus to Listener %s timeout: %s", name, topicMessage)
						continue
					}
				}
			}
		}
	}
}

var ptr *TypeCallbackProcessor = nil
var once sync.Once
var inChanSize = 1024
var timeout = 10 * time.Second

func TypeCallbackProcessorFactory() *TypeCallbackProcessor {
	once.Do(func() {
		ptr = &TypeCallbackProcessor{
			robotStatusListenerMap: make(map[string](chan robotStatus.CallbackTopic)),
			inChan:                 make(chan string, inChanSize),
		}
		go ptr.run()
	})
	return ptr
}
