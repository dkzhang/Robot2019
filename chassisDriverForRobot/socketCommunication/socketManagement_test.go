package socketCommunication

import (
	"log"
	"testing"
	"time"
)

func TestSocketManagement(t *testing.T) {
	serverIPandPort := "192.168.10.10:31001"

	chanSize := 1024
	var commandChan chan CommandStruct = make(chan CommandStruct, chanSize)

	psm := SocketManagementFactory(serverIPandPort, commandChan)
	resultChan, _ := psm.GetResultAndFeedbackChan()

	psm.GoRun()

	go PrintResultChan(resultChan)

	sleepTime := time.Second * 180

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

	//commandChan <- CommandStruct{
	//	Command: "/api/move?marker=3",
	//}

	commandChan <- CommandStruct{
		Command: "/api/move?markers=1,2,3,4,1,Charger",
	}

	time.Sleep(sleepTime)
	psm.Cancel()
}

func PrintResultChan(resultChan chan CommandResultStruct) {
	for {
		result := <-resultChan
		log.Printf("CommandResultStruct.strJSON = %s", result.strJSON)
	}

}
