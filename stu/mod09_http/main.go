package main

import (
	"fmt"
	"net"
)

// http 请求包格式

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Listen err = ", err)
		return
	}

	// 阻塞等待用户的连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err = ", err)
		return
	}

	defer conn.Close()

	// 接受客户端的数据

	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err = ", err)
		return
	}
	fmt.Printf("#%v#", string(buf[:n]))

}
