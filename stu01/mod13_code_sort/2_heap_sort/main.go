package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSort(arr))

}

//创建最大堆或者最小堆
//调整堆

func HeapSortMax(arr []int, length int) []int {

	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 // 二叉樹深度
	for i := depth; i >= 0; i-- {
		topMax := i //假定最大位置就在
		leftChild := 2*i + 1
		rightChild := 2*i + 2

		if leftChild <= length-1 && arr[leftChild] > arr[topMax] {
			topMax = leftChild
		} else if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
			topMax = rightChild
		}

		if topMax != i {
			arr[i], arr[topMax] = arr[topMax], arr[i]
		}
	}
	return arr
}

// 交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - i
		HeapSortMax(arr, lastLen)
		if i < length {
			arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
		}
	}
	return arr
}
