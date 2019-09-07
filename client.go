package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		os.Exit(1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		fmt.Printf("输入发送的内容：")
		fmt.Scan(&buf)

		n, err := conn.Write(buf)
		if err != nil {
			break
		}
		fmt.Println("已发送", n, string(buf))
	}

}
