package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cncamp/golang/stu0/day5/mode1"
	"io"
	"io/ioutil"
	"os"
)

type Person struct {
	name string
	age  int
}

type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type M1 struct {
	Monster //继承
	spring  int
}

//结构体的方法
func (a *Monster) test() {
	fmt.Println(a.Name)
}

func (a *Monster) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", a.Name, a.Age)
	return str
}

//接口

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

//让Phone实现Usb接口的方法

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
}

//让Camera 实现 USb接口的方法

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

//usb变量会更具传入的实参来判断是Phone还是Camera

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

//函数
func main() {

	fmt.Println("方式一")
	var per1 Person
	per1.name = "张三"
	per1.age = 12
	fmt.Println("per1=", per1)

	fmt.Println("方式二")
	per2 := Person{"李四", 30}
	fmt.Println("per2=", per2)

	fmt.Println("方式三")
	per3 := new(Person)
	(*per3).name = "王五"
	per3.age = 40
	fmt.Println("per3=", per3)

	fmt.Println("方式四")
	per4 := &Person{}
	(*per4).name = "麻子"
	per4.age = 50
	fmt.Println("per4=", per4)

	fmt.Println("将结构体序列化为 json格式字串")
	monster := Monster{"宋江", 500}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("josn 处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
	monster.test()
	//实现了*xx类型的string方法，就会自动调用
	fmt.Println(&monster)

	//挎包调用
	var stu = mode1.NewStudent("tom~", 98.8)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, " score=", stu.GetScore())

	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)

	//断言
	var a interface{}
	var point Monster = Monster{"哈哈", 2}
	a = point
	var b Monster
	//类型断言
	b = a.(Monster)
	fmt.Println(b)

	//文件
	file, err := os.Open("E:\\test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v", file)

	defer file.Close()

	//创建一个*Reader，是带缓存的
	const (
		defaultBufSize = 4096
	)

	reader := bufio.NewReader(file)
	for {
		//读到一个换行就结束
		str, err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF表示文件末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

	//读一个文件的内容写到另一个文件里
	file1Path := "E:\\test 1.txt"
	file2Path := "E:\\test 2.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}

	//命令行参数
	fmt.Println("命令行的参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

}
