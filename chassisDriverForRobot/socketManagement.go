package chassisDriverForRobot

import (
	"log"
	"net"
)

func SocketSend(conn *net.TCPConn, commandChan chan string, errorChan chan SocketErrorStuct) {

}

func SocketReceive(conn *net.TCPConn, resultChan chan CommandResultStruct, errorChan chan SocketErrorStuct) {

}

func SocketManagement(serverIP string, commandChan chan string, resultChan chan CommandResultStruct) {
	for {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", serverIP)
		if err != nil {
			log.Fatalf("net.ResolveTCPAddr error: %v", err)
			continue
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			log.Fatalf("net.DialTCP error: %v", err)
			continue
		}

		errorChan := make(chan SocketErrorStuct, 2)

		go SocketSend(conn, commandChan, errorChan)
		go SocketReceive(conn, resultChan, errorChan)

		//wait if there is a socket error
		<-errorChan
		conn.Close()
	}
}

type SocketErrorStuct struct {
	err error
	str string
}
