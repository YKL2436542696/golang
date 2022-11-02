package main

import (
	"fmt"
)

func main() {

	var nameArr [5]string
	nameArr[0] = "张三"
	nameArr[1] = "李四"
	nameArr[2] = "王五"
	nameArr[3] = "二麻"
	nameArr[4] = "赵前"
	var scoreArr [5]int
	scoreArr[0] = 70
	scoreArr[1] = 60
	scoreArr[2] = 80
	scoreArr[3] = 90
	scoreArr[4] = 65

	fmt.Println("--------排序前---------")
	for i := 0; i < len(nameArr); i++ {
		fmt.Printf("%v的成绩为%v\n", nameArr[i], scoreArr[i])
	}

	for i := 0; i < len(scoreArr); i++ {
		for j := 0; j < len(scoreArr)-1-i; j++ {
			if scoreArr[j] > scoreArr[j+1] {
				switInt(&scoreArr, j, j+1)
				switString(&nameArr, j, j+1)
			}
		}
	}
	fmt.Println("--------排序后---------")
	for i := 0; i < len(nameArr); i++ {
		fmt.Printf("%v的成绩为%v\n", nameArr[i], scoreArr[i])
	}

}
func switInt(arr *[5]int, i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
func switString(arr *[5]string, i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
