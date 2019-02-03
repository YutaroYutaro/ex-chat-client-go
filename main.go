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
	//if len(os.Args) == 1 {
	//	println("need request parameter")
	//	os.Exit(1)
	//}
	//
	//echoContents := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8001")

	if err != nil {
		println("error tcp resolve failed", err.Error())
		os.Exit(1)
	}

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer tcpConn.Close()

	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		echoContents := stdin.Text()

		if echoContents == "exit" {
			fmt.Println("connection close.")
			break
		}

		SendEcho(tcpConn, echoContents)

		echo := GetEcho(tcpConn)

		println("echo: ", string(echo))
		println("receive success.")
	}
}

func SendEcho(conn *net.TCPConn, msg string) {
	_, err := conn.Write([]byte(msg))

	if err != nil {
		println("error send request: ", err.Error())
	} else {
		println("request sent")
	}
}

func GetEcho(conn *net.TCPConn) string {
	bufRecv := make([]byte, RecvBufLen)

	_, err := conn.Read(bufRecv)

	if err != nil {
		println("error while receive response: ", err.Error())
		return ""
	}

	return string(bufRecv)
}
