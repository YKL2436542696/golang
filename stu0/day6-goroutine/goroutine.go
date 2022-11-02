package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test () hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁，
	//sync 是包：synchornized 同步
	//Mutex：是互斥
	lock sync.Mutex
)

func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//将res放入到myMap
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	//开启一个协程
	//go test()
	//for i := 0; i <= 10; i++ {
	//	fmt.Println("main() hello,golang " + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}

	//获取当前系统cpu的数量
	//num := runtime.NumCPU()
	//runtime.GOMAXPROCS(num)
	//fmt.Println("num=", num)

	//通过加锁
	//for i := 1; i <= 200; i++ {
	//	go test2(i)
	//}
	//time.Sleep(time.Second * 10)
	//lock.Lock()
	//for i, v := range myMap {
	//	fmt.Printf("map[%d]=%d\n", i, v)
	//}
	//lock.Unlock()

	//管道使用
	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	//2. 看看intChan的值是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	//4.再来看看管道长度、容量
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))
	fmt.Printf("\n")

	//5.在管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//6.在没有使用协程的情况下，如果我们的管道数据以及全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	//	num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4, "num5=", num4)
	//关闭后管道无法写入，但是可以读取
	close(intChan)
	fmt.Println()
	//创建两个管道
	intChan1 := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan1)
	go readData(intChan1, exitChan)

	//time.Sleep(time.Second)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}

	}
}

//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeData", i)
		//time.Sleep(time.Second)
	}
	close(intChan) //关闭管道
}

//read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}
