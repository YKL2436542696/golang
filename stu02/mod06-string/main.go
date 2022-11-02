package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 2
	}()

	defer close(ch)

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	//
	time.Sleep(time.Second)
}
