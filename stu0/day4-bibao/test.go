package main

import (
	"fmt"
	"strings"
)

//累加器

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func AddUpper2() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str=", str)
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//闭包测试
	f := AddUpper()
	f2 := AddUpper2()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Printf("\n")
	fmt.Println(f2(1))
	fmt.Println(f2(2))
	fmt.Println(f2(3))
	fmt.Printf("\n")
	f3 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f3("winter"))
	fmt.Println("文件名处理后=", f3("brid.jpg"))
	fmt.Printf("\n")

	//二维数组测试
	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)
	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[0][1]的地址%p\n", &arr2[0][1])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])
	fmt.Printf("\n")

	//map
	fmt.Println("第一种方式")
	//声明
	var a map[string]string
	//使用前，要先make,给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "张三"
	a["no2"] = "李四"
	a["no3"] = "王五"
	fmt.Println(a)
	fmt.Printf("\n")
	fmt.Println("第二种方式")
	b := make(map[string]string)
	b["no1"] = "甲"
	b["no2"] = "乙"
	b["no3"] = "丙"
	fmt.Println(b)
	fmt.Printf("\n")
	fmt.Println("第三种方式")
	c := map[string]string{
		"hero1": "A",
		"hero2": "B",
		"hero3": "C",
	}
	c["hero4"] = "D"
	fmt.Println(c)

}
