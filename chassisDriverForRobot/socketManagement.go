package chassisDriverForRobot

import (
	"fmt"
	"log"
	"net"
)

func SocketReceive(conn *net.TCPConn, resultChan chan CommandResultStruct, errorChan chan error) {
	buff := make([]byte, 1024*10)

	for {
		counts, err := conn.Read(buff)
		if err != nil {
			errorChan <- fmt.Errorf("SocketReceive write wrong bytes: %v", err)
			if err = conn.Close(); err != nil {
				log.Fatalf("conn.Close() in SocketSend error: %v", err)
			}
			return
		}
		cs := CommandResultStruct{}
		cs.strJSON = string(buff[:(counts - 1)])
		resultChan <- cs
	}
}

func SocketSend(conn *net.TCPConn, commandChan chan string, errorChan chan error, rCommand *string) {
	var command string
	for {
		if *rCommand != "" {
			//*rCommand 保存着最近一次未发送成功的命令
			//因此如果*rCommand不为空，则优先发送该命令
			command = *rCommand
		} else {
			command = <-commandChan
		}
	}
	counts, err := conn.Write([]byte(command))
	if err != nil {
		//网络发送错误
		*rCommand = command
		errorChan <- fmt.Errorf("Socket write wrong bytes: %v", err)
		if err = conn.Close(); err != nil {
			log.Fatalf("conn.Close() in SocketSend error: %v", err)
		}
		return
	} else if counts != len([]byte(command)) {
		//发送的字节数错误
		*rCommand = command
		errorChan <- fmt.Errorf("Socket write wrong bytes: %v", err)
		if err = conn.Close(); err != nil {
			log.Fatalf("conn.Close() in SocketSend error: %v", err)
		}
		return
	} else {
		//发送成功
		*rCommand = ""
	}
}

func SocketManagement(serverIP string, commandChan chan string, resultChan chan CommandResultStruct) {
	errorChan := make(chan error, 2)
	command := ""

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

		go SocketSend(conn, commandChan, errorChan, &command)
		go SocketReceive(conn, resultChan, errorChan)

		//wait if there is a socket error
		err1 := <-errorChan
		log.Fatalf("SocketManagement get error1: %v", err1)
		err2 := <-errorChan
		log.Fatalf("SocketManagement get error1: %v", err2)
	}
}
