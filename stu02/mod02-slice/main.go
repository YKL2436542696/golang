package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}

	ch := arr[1:]

	ch2 := make([]int, 10)
	copy(ch2, ch)

	fmt.Println("&ch2[0]：", &ch2[0])
	fmt.Println("ch2.len：", len(ch2))
	fmt.Println("ch2.cap：", cap(ch2))
	fmt.Println(ch2)
	fmt.Println()

	fmt.Println("&arr[0]：", &arr[0])
	fmt.Println("&ch[0]：", &ch[0])
	fmt.Println("ch.len：", len(ch))
	fmt.Println("ch.cap：", cap(ch))
	fmt.Println()

	ch = append(ch, 2)
	fmt.Println("&arr[0]：", &arr[0])
	fmt.Println("&ch[0]：", &ch[0])
	fmt.Println("ch.len：", len(ch))
	fmt.Println("ch.cap：", cap(ch))
	fmt.Println()

	fmt.Println(ch)

}
