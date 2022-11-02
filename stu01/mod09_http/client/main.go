package main

import (
	"fmt"
	"net"
)

func main() {

	// 主动连接服务器
	conn, err := net.Dial("tcp", "8000")

	if err != nil {
		fmt.Println("dial err = ", err)
		return
	}

	defer conn.Close()

	requestBuf := "GET /go HTTP/1.1\nHost: localhost:8000\nConnection: keep-alive\nsec-ch-ua: \"Google Chrome\";v=\"105\", \"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"105\"\nsec-ch-ua-mobile: ?0\nsec-ch-ua-platform: \"Windows\"\nUpgrade-Insecure-Requests: 1\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\nSec-Fetch-Site: none\nSec-Fetch-Mode: navigate\nSec-Fetch-User: ?1\nSec-Fetch-Dest: document\nAccept-Encoding: gzip, deflate, br\nAccept-Language: zh-CN,zh;q=0.9,zh-TW;q=0.8\n"

	// 先发请求包，服务器才会回相应包
	conn.Write([]byte(requestBuf))

	// 接受服务器回复的响应包
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err = ", err)
		return
	}

	// 打印响应保温
	fmt.Printf("#%v#", string(buf[:n]))

}
