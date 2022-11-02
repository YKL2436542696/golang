package main

import (
	"fmt"
	"time"
)

type Task struct {
	A int
	B int
}

var ch = make(chan Task, 100)

// 生产者
func producer() {
	for i := 0; i < 10; i++ {
		ch <- Task{A: i + 3, B: i - 8}
	}
}

// 消费者
func consumer() {
	for i := 0; i < 10; i++ {
		task := <-ch
		sum := task.A + task.B
		fmt.Println(sum)
	}
}

func main() {
	go producer()
	go consumer()
	time.Sleep(time.Second)
}
