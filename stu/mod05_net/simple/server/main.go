package main

import (
	"fmt"
	"net"
)

func main() {

	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	// 阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 接受用户的请求
	buf := make([]byte, 1024)
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("err1 = ", err)
		return
	}

	fmt.Println("buf = ", string(buf[:n]))

	defer conn.Close()

}