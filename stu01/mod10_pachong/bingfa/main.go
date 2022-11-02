package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// 爬取网页内容

func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer resp.Body.Close()

	// 读取网页body内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { // 读取结束，或者，出问题
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

// 爬取一个网页

func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=apex%E8%8B%B1%E9%9B%84&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Println("url = ", url)

	// 2）爬
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	// 写入文件
	fileName := strconv.Itoa(i) + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	f.WriteString(result) // 写入
	f.Close()             //关闭文件

	page <- i
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	page := make(chan int)

	// 明确目标 （要知道你准被在哪个范围或者网站去搜索）
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页（ >=1 ） : ")
	fmt.Scan(&start)
	fmt.Println("请输入终止页（ >= 起始页） : ")
	fmt.Scan(&end)

	DoWork(start, end)
}
