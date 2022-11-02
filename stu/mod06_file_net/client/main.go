package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	// 以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	// 读文件内容
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err = ", err)
			}
			return
		}

		// 发送内容
		conn.Write(buf[:n]) // 给服务器发送内容
	}
}

func main() {
	// 提示输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)

	// 获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(" os.Stat err= ", err)
		return
	}

	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(" net.Dial err= ", err)
		return
	}

	defer conn.Close()

	// 给接受方先发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println(" conn.Write err= ", err)
		return
	}

	// 接受对方的回复，如果回复ok，说明对方准备好，可以发送文件
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(" conn.Read err= ", err)
		return
	}

	if "ok" == string(buf[:n]) {
		// 发送文件内容
		SendFile(path, conn)
	}

}
