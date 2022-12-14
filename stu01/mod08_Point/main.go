package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var s1 []int
	s2 := make([]int, 10, 20)
	s4 := make([]int, 10)

	fmt.Printf("s1 pointer:%+v", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Println()
	fmt.Printf("s2 pointer:%+v", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Println()
	fmt.Printf("s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))

	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
