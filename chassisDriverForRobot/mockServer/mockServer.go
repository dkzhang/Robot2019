package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	serverIP := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverIP)
	if err != nil {
		fmt.Errorf("net.ResolveTCPAddr error: %v", err)
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Errorf("net.ListenTCP error: %v", err)
		return
	}
	log.Printf("Begin ListenTCP... \n")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Errorf("net.ListenTCP error: %v", err)
		return
	}
	log.Printf("Accept TCP Connection... \n")
	defer conn.Close()

	for {

		buff := make([]byte, 1024*10)
		counts, err := conn.Read(buff)
		if err != nil {
			fmt.Errorf("conn.Read error: %v", err)
			return
		} else {
			log.Printf("read %d bytes: %s", counts, string(buff[:counts-1]))
		}

		strJSON := `
		{
			"type":"response",
			"command":"/api/move",
			"uuid":"12345",
			"status":"OK",
			"error_message":"",
			"task_id":"xxx"
		}
		`

		counts, err = conn.Write([]byte(strJSON))
		if err != nil {
			fmt.Errorf("conn.Write error: %v", err)
			return
		} else {
			log.Printf("Write %d bytes: %s", counts, strJSON)
		}
	}
}
