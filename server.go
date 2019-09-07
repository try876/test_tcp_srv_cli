package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		fmt.Println("listen err:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("listen ok.")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {

	defer conn.Close()

	buf := make([]byte, 1024)

	total := 0
	for {
		reqLen, err := conn.Read(buf[total:])
		if err != nil {
			fmt.Println("Error ", err.Error())
			break
		}
		total += reqLen

		if total > 1000 {
			break
		}

		fmt.Println("aa:", string(buf[:total]))
	}

}
