package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string
	Addr string
}

// 保存在线用户
var onlineMap map[string]Client

// 消息管道
var message = make(chan string)

func Manager() { //收到消息，则遍历map,给每个用户发送消息
	// 给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息前，这里会阻塞
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func WriteMessageToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送消息
		conn.Write([]byte(msg + "\n"))

	}
}

func MakeMessage(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + "：" + msg
	return
}

func HandleConn(coon net.Conn) {
	defer coon.Close()
	// 获取客户端网络地址
	cliAddr := coon.RemoteAddr().String()

	// 创建一个结构体
	cli := Client{make(chan string), cliAddr, cliAddr}
	// 把用户结构体添加到map
	onlineMap[cliAddr] = cli

	// 新开一个协程，专门给当前客户端发送消息
	go WriteMessageToClient(cli, coon)

	// 广播某个用户在线
	message <- MakeMessage(cli, "login")

	// 提示，我是谁
	cli.C <- MakeMessage(cli, "I am here")

	isQuit := make(chan bool)  // 对方是否主动退出
	hasData := make(chan bool) //对方是否有数据发送

	// 新开一个协程，循环接受用户发送过来的数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := coon.Read(buf)
			if n == 0 { // 对方断开连接
				isQuit <- true
				fmt.Println("coon.Read err = ", err)
				return
			}
			msg := string(buf[:n-2])
			if len(msg) == 3 && msg == "who" {
				// 遍历 map,给当前用户发送所有成员消息 [查询用户]
				coon.Write([]byte("user List:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					coon.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//rename|mike [修改用户名]
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				coon.Write([]byte("rename ok\n"))
			} else {
				// 转发此内容
				message <- MakeMessage(cli, msg)
			}
			hasData <- true
		}
	}()

	for {
		// 通过select 检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)               //移除当前用户
			message <- MakeMessage(cli, "login out") //广播下线消息
			return

		case <-hasData:

		case <-time.After(30 * time.Second):
			delete(onlineMap, cliAddr)
			message <- MakeMessage(cli, "time out leave out")
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "8000")
	if err != nil {
		fmt.Println(" net.Listen err = ", err)
		return
	}
	defer listener.Close()

	// 新开协程 转发消息，收到消息，则遍历map,给每个用户发送消息
	go Manager()

	// 主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(" listener.Accept err = ", err)
			continue
		}
		// 处理用户连接
		go HandleConn(conn)
	}
}
