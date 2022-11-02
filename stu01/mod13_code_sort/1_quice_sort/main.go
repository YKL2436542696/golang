package main

import "fmt"

//将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于
//将两边数据进行递归调用步骤1
//将所有数据合并

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	sp := arr[0]
	low := make([]int, 0, 0)
	hight := make([]int, 0, 0)
	mid := make([]int, 0, 0)

	mid = append(mid, sp)

	for i := 1; i < len(arr); i++ {
		if arr[i] > sp {
			hight = append(hight, arr[i])
		} else if arr[i] < sp {
			low = append(low, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

func main() {

	arr := []int{10, 9, 2, 1, 5, 4, 6, 7, 10, 123, 234}
	fmt.Println(QuickSort(arr))

}
