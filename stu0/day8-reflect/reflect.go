package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的 type（类型）,kind（类别）,值
	//1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	rkind := rVal.Kind()
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v rVal 类型type=%T\n", rVal, rVal)
	fmt.Println("类别 rkind=", rkind)

	//将rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)

}

//对结构体的反射
func reflectTest02(b interface{}) {

	//通过反射获取的传入的变量的 type,kind,值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	rkind := rVal.Kind()
	fmt.Printf("类型type=%T，类别kind=%v\n", rVal, rkind)

	//将rVal转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iV, iV)
	//将interface{}通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu01.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func testInt(b interface{}) {
	va := reflect.ValueOf(b)
	fmt.Printf("va type=%T\n", va)
	//va.Elem()用于获取指针指向变量
	va.Elem().SetInt(110)
	fmt.Printf("va=%v\n", va)
}

func main() {

	//请编写一个案例
	//演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作

	//1. 先定义一个int
	var num int = 100
	reflectTest01(num)

	fmt.Println()
	//2. 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
	fmt.Println()

	//使用SetXXX方法设置，需要通过对应的指针类型来完成
	testInt(&num)
	fmt.Println("num=", num)

}
