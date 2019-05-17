package socketCommunication

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func SocketReceive(conn *net.TCPConn, resultChan chan CommandResultStruct, errorChan chan error, pwg *sync.WaitGroup) {
	defer pwg.Done()

	buff := make([]byte, 1024*10)

	for {
		counts, err := conn.Read(buff)
		if err != nil {
			errorChan <- fmt.Errorf("SocketReceive conn.Read error: %v", err)
			return
		}
		cs := CommandResultStruct{}
		cs.strJSON = string(buff[:(counts - 1)])
		resultChan <- cs

		select {
		case errMsg := <-errorChan:
			log.Fatalf("SocketReceive will return because of revicing message from errorChan: %v .", errMsg)
			return
		default:
			continue
		}
	}
}
func SocketSend(conn *net.TCPConn, commandChan chan CommandStruct,
	errorChan chan error, cancelChan chan interface{},
	pCommand *CommandStruct) (*CommandStruct, error) {

	if pCommand == nil {
		//没有上次未发送成功的命令
		select {
		case errMsg := <-errorChan:
			return nil, fmt.Errorf("SocketSend will return because of revicing message from errorChan: %v .", errMsg)
		case <-cancelChan:
			return nil, fmt.Errorf("cancel")
		case commandStruct := <-commandChan:
			counts, err := conn.Write([]byte(commandStruct.Command))
			if err != nil {
				//网络发送错误
				//取出命令后发送错误，需要返回当前发送出错的命令以备重发
				return &commandStruct, fmt.Errorf("SocketSend will return because of Socket write error: %v.", err)
			} else if counts != len([]byte(commandStruct.Command)) {
				//发送的字节数错误
				//取出命令后发送错误，需要返回当前发送出错的命令以备重发
				return &commandStruct, fmt.Errorf("SocketSend will return because of Socket wrong bytes.")
			} else {
				//发送成功
				return nil, nil
			}
		}
	} else {
		//上次的命令没有发送成功，先不从commandChan中取命令，继续尝试发送上次未发送成功的命令
		select {
		case errMsg := <-errorChan:
			return nil, fmt.Errorf("SocketSend will return because of revicing message from errorChan: %v .", errMsg)
		case <-cancelChan:
			return nil, fmt.Errorf("cancel")
		default:
			counts, err := conn.Write([]byte(pCommand.Command))
			if err != nil {
				//网络发送错误
				//取出命令后发送错误，需要返回当前发送出错的命令以备重发
				return pCommand, fmt.Errorf("SocketSend will return because of Socket write error: %v.", err)
			} else if counts != len([]byte(pCommand.Command)) {
				//发送的字节数错误
				//取出命令后发送错误，需要返回当前发送出错的命令以备重发
				return pCommand, fmt.Errorf("SocketSend will return because of Socket wrong bytes.")
			} else {
				//发送成功
				return nil, nil
			}
		}
	}
}

func SocketManagement(serverIP string, commandChan chan CommandStruct, resultChan chan CommandResultStruct, cancelChan chan interface{}) {
	//SocketManagement 管理着两个go routine，分别用于发送和接收
	//如果发送或接收时网络出错，则使用errorChan通知另外一个go routine退出
	//当两个go routine都退出时，SocketManagement尝试重新连接
	errorChan := make(chan error, 1)
	var wg sync.WaitGroup

	var pcs *CommandStruct = nil

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

		wg.Add(1)
		go SocketReceive(conn, resultChan, errorChan, &wg)

		//发送不用go routine
		for {
			pcs, err = SocketSend(conn, commandChan, errorChan, cancelChan, pcs)
			if err != nil {
				log.Fatalf("SocketSend error: %v", err)
				if errClose := conn.Close(); err != nil {
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
