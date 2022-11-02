package main

import (
	"flag"
	"fmt"
	"sync"
)

var (
	mutex      sync.Mutex
	balance    int
	protecting uint                      // 是否加锁
	sign       = make(chan struct{}, 10) //通道，用于等待所有goroutine
)

// 存钱
func deposit(value int) {
	defer func() {
		sign <- struct{}{}
	}()

	if protecting == 1 {
		mutex.Lock()
		defer mutex.Unlock()
	}

	fmt.Printf("余额: %d\n", balance)
	balance += value
	fmt.Printf("存 %d 后的余额: %d\n", value, balance)
	fmt.Println()

}

// 取钱
func withdraw(value int) {
	defer func() {
		sign <- struct{}{}
	}()

	if protecting == 1 {
		mutex.Lock()
		defer mutex.Unlock()
	}

	fmt.Printf("余额: %d\n", balance)
	balance -= value
	fmt.Printf("取 %d 后的余额: %d\n", value, balance)
	fmt.Println()

}

func main() {
	protecting = 1
	for i := 0; i < 5; i++ {
		go withdraw(500) // 取500
		go deposit(500)  // 存500
	}

	for i := 0; i < 10; i++ {
		<-sign
	}
	fmt.Printf("当前余额: %d\n", balance)
}

func init() {
	balance = 1000 // 初始账户余额为1000
	flag.UintVar(&protecting, "protecting", 0, "是否加锁，0表示不加锁，1表示加锁")
}
