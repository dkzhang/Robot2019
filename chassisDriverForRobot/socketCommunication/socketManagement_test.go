package socketCommunication

import (
	"log"
	"testing"
	"time"
)

func TestSocketManagement(t *testing.T) {
	//serverIP := ":7777"
	serverIP := "192.168.10.10:31001"

	chanSize := 1024
	var commandChan chan CommandStruct = make(chan CommandStruct, chanSize)
	var resultChan chan CommandResultStruct = make(chan CommandResultStruct, chanSize)
	var cancelChan chan interface{} = make(chan interface{})

	go SocketManagement(serverIP, commandChan, resultChan, cancelChan)

	go PrintResultChan(resultChan)

	sleepTime := time.Second * 120

	commandChan <- CommandStruct{
		Command: "/api/robot_info",
	}

	commandChan <- CommandStruct{
		Command: "/api/markers/query_list",
	}

	//commandChan <- CommandStruct{
	//	Command: "/api/move?marker=1",
	//}
	//
	//commandChan <- CommandStruct{
	//	Command: "/api/move?marker=2",
	//}

	commandChan <- CommandStruct{
		Command: "/api/move?marker=3",
	}

	commandChan <- CommandStruct{
		Command: "/api/move?marker=Charger",
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
