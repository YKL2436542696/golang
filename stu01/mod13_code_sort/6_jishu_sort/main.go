package main

import "fmt"

func main() {
	var arr [3][]int
	myarr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	for i := 0; i < len(myarr); i++ {
		arr[myarr[i]-1] = append(arr[myarr[i]-1], myarr[i])
	}
	fmt.Println(arr)
}
