package socketCommunication

import (
	"log"
	"testing"
	"time"
)

func TestSocketManagement(t *testing.T) {
	serverIP := ":7777"

	chanSize := 1024
	var commandChan chan CommandStruct = make(chan CommandStruct, chanSize)
	var resultChan chan CommandResultStruct = make(chan CommandResultStruct, chanSize)
	var cancelChan chan interface{} = make(chan interface{})

	go SocketManagement(serverIP, commandChan, resultChan, cancelChan)

	go PrintResultChan(resultChan)

	sleepTime := time.Second * 10

	commandChan <- CommandStruct{
		Command: time.Now().String(),
	}

	time.Sleep(sleepTime)
	close(cancelChan)
}

func PrintResultChan(resultChan chan CommandResultStruct) {
	for {
		result := <-resultChan
		log.Printf("CommandResultStruct.strJSON = %s", result.strJSON)
	}

}
