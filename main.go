package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	RecvBufLen = 1024
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8001")

	if err != nil {
		println("error tcp resolve failed", err.Error())
		os.Exit(1)
	}

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)

	stdin := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter your name")
	stdin.Scan()
	name := stdin.Text()
	SendEcho(tcpConn, name)

	go GetEcho(tcpConn)

	for stdin.Scan() {
		echoContents := stdin.Text()

		if echoContents == "exit" {
			fmt.Println("connection close.")
			tcpConn.Close()
			break
		}

		SendEcho(tcpConn, echoContents)
	}
}

func SendEcho(conn *net.TCPConn, msg string) {
	_, err := conn.Write([]byte(msg))

	if err != nil {
		println("error send request: ", err.Error())
	}
}

func GetEcho(conn *net.TCPConn) {
	for {
		bufRecv := make([]byte, RecvBufLen)

		_, err := conn.Read(bufRecv)

		if err != nil {
			fmt.Println("error while receive response: ", err.Error())
			return
		}

		fmt.Println(string(bufRecv))
	}
}
