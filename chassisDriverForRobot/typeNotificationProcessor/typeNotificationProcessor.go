package typeNotificationProcessor

import (
	"Robot2019/chassisDriverForRobot/typeNotificationProcessor/typeNotificationStructure"
	"fmt"
	"log"
	"sync"
	"time"
)

type TypeNotificationProcessor struct {
	ListenerMap map[string](chan typeNotificationStructure.Notification)
	inChan      chan string
	cancelChan  chan interface{}

	runOnce sync.Once
}

// 由于消息收到后都是从这里分发出去，因此这里是一个消息汇聚点
// 为避免瓶颈效应，这里不对消息进行filter，各个goroutine收到所有notification后进行filter
func (proc *TypeNotificationProcessor) RegisterRobotStatusListener(name string, notificationChan chan typeNotificationStructure.Notification) (err error) {
	if _, ok := proc.ListenerMap[name]; !ok {
		proc.ListenerMap[name] = notificationChan
		return nil
	} else {
		return fmt.Errorf("RobotStatus Listener %s already registered!", name)
	}
}

func (proc *TypeNotificationProcessor) UnregisterRobotStatusListener(name string) {
	delete(proc.ListenerMap, name)
}

func (proc *TypeNotificationProcessor) GetInChan() chan string {
	return proc.inChan
}

func (proc *TypeNotificationProcessor) run() {
	for {
		var notificationMessage string

		select {
		case notificationMessage = <-proc.inChan:
			log.Printf("Recive callback topic: %s", notificationMessage)
		case <-proc.cancelChan:
			return
		}

		nf, err := typeNotificationStructure.UnmarshalJSON(notificationMessage)
		if err != nil {
			log.Fatalf("UnmarshalJSON Notification error: %v", err)
			continue
		} else {
			log.Printf("UnmarshalJSON Notification success: %v", nf)
		}

		//把消息发个已注册监听的各个channel
		//如果某个channel超时，则跳过进行下一个
		for name, theChan := range proc.ListenerMap {
			select {
			case theChan <- nf:
				//Do nothing
			case <-time.After(timeout):
				log.Fatalf("Send Callback Topic RobotStatus to Listener %s timeout: %s", name, notificationMessage)
				continue
			}
		}
	}
}

func (proc *TypeNotificationProcessor) IsRunning() bool {
	select {
	case <-proc.cancelChan:
		return false
	default:
		return true
	}
}

func (proc *TypeNotificationProcessor) Cancel() {
	close(proc.cancelChan)
}

func (proc *TypeNotificationProcessor) GoRun() {
	proc.runOnce.Do(func() {
		go proc.run()
	})
}

var inChanSize = 1024
var timeout = 1 * time.Second

func TypeNotificationProcessorFactory() *TypeNotificationProcessor {
	ptr := &TypeNotificationProcessor{
		ListenerMap: make(map[string](chan typeNotificationStructure.Notification)),
		inChan:      make(chan string, inChanSize),
		cancelChan:  make(chan interface{}),
	}
	return ptr
}
