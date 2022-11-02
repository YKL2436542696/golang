package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	// 获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect successful,addr = ", addr)

	buf := make([]byte, 2048)

	for {
		// 读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err= ", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		fmt.Println("len = ", len(string(buf[:n])))
		fmt.Println("n = ", n)
		// n-1 : 减去上面换行符的字符
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, " exit")
			return
		}

		// 把数据转化为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {
	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	defer listener.Close()

	//接受多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		// 处理用户发送的数据,新建一个协程
		go HandleConn(conn)

	}

}
