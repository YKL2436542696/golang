package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello,World!")

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("winter"))
	fmt.Println(f2("1.jpg"))

	//字符串遍历 r:=[]rune(str)
	str2 := "hello成都"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}
	fmt.Println("")

	//字符串转整数 n,err := strconv.Atoi("12")
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", n)
	}
	fmt.Println("")
	//整数转字符串
	str := strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T", str, str)
	fmt.Println("")
	//字符串转 []byte []byte("hello go")
	bytes := []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)
	fmt.Println("")
	//[]byte转 字符串 string([]byte{97,98,99})
	str3 := string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str3)
	fmt.Println("")
	//10进制转 2,8,16进制 strconv.FormatInt(123,2)
	str4 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str4)
	str5 := strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是=%v\n", str5)
	fmt.Println("")
	//查看子串是否在指定的字符串中 strings.Contains("abcde","abc")
	b := strings.Contains("abcde", "abc")
	fmt.Printf("b=%v\n", b)
	fmt.Println("")
	//统计一个字符串有几个指定的子串 strings.Count("ceheese","e")
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)
	fmt.Println("")
	//不区分大小写的字符串比较 strings.EqualFold
	b1 := strings.EqualFold("abc", "Abc")
	fmt.Printf("结果b=%v\n", b1)
	fmt.Println("")
	//返回子串在字符串第一次出现的index值  strings.Index /LastIndex
	index := strings.Index("asdqwqwe", "qwq")
	fmt.Printf("index=%v\n", index)
	fmt.Println("")
	//将指定的子串替换成 另外一个子串 strings.Replace
	str6 := "go go hello"
	str7 := strings.Replace(str6, "go", "北京", -1)
	fmt.Printf("str=%v str2=%v\n", str6, str7)
	fmt.Println("")
	//按照指定的字符为分割标识，来分割字符串 strings.Split
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)
	fmt.Println("")
	//将字符串的字母进行大小写的转换 strings.ToLower / ToUpper
	str8 := "goLang Hello"
	str9 := strings.ToLower(str8)
	str10 := strings.ToUpper(str8)
	fmt.Printf("str9=%v\n str10=%v\n", str9, str10)
	fmt.Println("")
	//将字符串左右两边的空格去掉 strings.TrimSpace
	str11 := strings.TrimSpace(" tb a lone abc  ")
	fmt.Printf("str11=%v\n", str11)
	fmt.Println("")
	//将字符串左右两边指定的字符去掉 strings.Trim /TrimLeft/TrimRight
	str12 := strings.Trim("! hello!", "!")
	fmt.Printf("str12=%q\n", str12)
	fmt.Println("")
	//判断字符串是否以指定的字符传开头 strings.HasPrefix /结束HasSuffix
	b3 := strings.HasPrefix("www.baidu.com", "www")
	fmt.Printf("b3=%v\n", b3)

}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}
