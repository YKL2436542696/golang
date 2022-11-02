package main

import "fmt"

func traveseChannel() {
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) //删除本行会死锁
	}()
	for ele := range ch { // 遍历并取走管道中的元素
		fmt.Println(ele)
	}
	fmt.Println("bye bye")
}

func main() {
	traveseChannel()
}
