package main

import "fmt"

func main() {

	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明定义一个切片
	//方式一：定义一个切片，然后让切片去引用一个创建好的数组
	fmt.Println("----方式一：定义一个切片，然后让切片去引用一个创建好的数组----")
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice)
	fmt.Println("slice 的元素个数 =", len(slice))
	fmt.Println("slice 的容量 =", cap(slice))
	fmt.Printf("\n")
	//方式二：通过make创建
	fmt.Println("----方式二：通过make创建----")
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10.3
	slice2[3] = 20.2
	fmt.Println("slice2 的元素是 =", slice2)
	fmt.Println("slice2 的元素个数 =", len(slice2))
	fmt.Println("slice2 的容量 =", cap(slice2))
	fmt.Printf("\n")
	//方式三：
	fmt.Println("----方式三：定义一个切片，直接就指定具体数组，原理类似make----")
	var slice3 []string = []string{"tom", "jack", "mary"}
	fmt.Println("slice3 的元素是 =", slice3)
	fmt.Println("slice3 的元素个数 =", len(slice3))
	fmt.Println("slice3 的容量 =", cap(slice3))

	//指向的数据空间是同一个，所以会改变
	slice4 := slice3[1:2]
	slice4[0] = "jacki"
	fmt.Println("slice4 的元素是 =", slice4)
	fmt.Println("slice3 的元素是 =", slice3)
	fmt.Println("slice4 的元素个数 =", len(slice4))
	fmt.Println("slice4 的容量 =", cap(slice4))

	slice4 = append(slice4, "张三", "李四", "王五")
	fmt.Println("扩充后")
	fmt.Println("slice4 的元素是 =", slice4)
	fmt.Println("slice4 的元素个数 =", len(slice4))
	fmt.Println("slice4 的容量 =", cap(slice4))
	fmt.Printf("\n")

	//拷贝
	var slice5 = make([]string, 10)
	copy(slice5, slice4)
	fmt.Println("slice5 的元素是 =", slice5)
	fmt.Println("slice5 的元素个数 =", len(slice5))
	fmt.Println("slice5 的容量 =", cap(slice5))
	fmt.Println("修改slice5")
	slice5[3] = "王二狗"
	fmt.Println("修改后slice5 的元素是 =", slice5)
	fmt.Println("修改后slice4 的元素是 =", slice4)
	fmt.Printf("\n")

	//修改字符传
	str := "hello world"
	arr1 := []byte(str)
	arr1[1] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)
	//若有文字，需要转为[]rune
	str2 := "你好"
	arr2 := []rune(str2)
	arr2[0] = '他'
	str2 = string(arr2)
	fmt.Println("str2=", str2)

}
