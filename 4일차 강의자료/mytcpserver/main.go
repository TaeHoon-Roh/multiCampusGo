package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	myListen, err := net.Listen("tcp", ":5554")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		connect, err := myListen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ConnectHandler(connect)
	}

	defer func() {
		myListen.Close()

	}()
}

func ConnectHandler(connect net.Conn) {
	recvBuf := make([]byte, 4096) // receive buffer: 4kB
	for {
		n, err := connect.Read(recvBuf)
		if err != nil {
			if io.EOF == err {
				fmt.Println("connection is closed from client; %v", connect.RemoteAddr().String())
				return
			}
			fmt.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			fmt.Println(string(data))
		}
	}
}
