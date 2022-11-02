package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	// 从键盘输入。Win平台 按回车，会多两个字符
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) // 从键盘读取内容，放在str
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
			// 把输入内容给服务器发送
			conn.Write(str[:n])
		}

	}()

	// 接受服务器回复的数据
	// 切片缓存
	buf := make([]byte, 1024)
	for {
		for {
			n, err := conn.Read(buf) // 接受服务器请求
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
			fmt.Println(string(buf[:n])) // 打印接受到的内容
		}
	}
}
