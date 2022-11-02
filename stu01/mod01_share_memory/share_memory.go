package main

import (
	"fmt"
	"sync"
	"time"
)

var mp sync.Map

func rwGlobalMemory() {
	if value, exists := mp.Load("mykey"); exists {
		fmt.Println(value)
	} else {
		mp.Store("mykey", "myvalue")
	}
}

func main() {
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()

	time.Sleep(time.Second)
}
