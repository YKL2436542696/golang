package main

import (
	"fmt"
	"sort"
)

// 在排序数组中查找元素的第一个和最后一个位置

func searchRange2(nums []int, target int) []int {

	rsp := []int{-1, -1}
	sum := 0

	for i, num := range nums {
		if target == num && sum == 0 {
			rsp[0] = i
			sum++
		}
		if target == num && sum > 0 {
			rsp[1] = i
		}
	}
	return rsp
}

func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 10
	n := searchRange(nums, target)
	fmt.Println(n)

}
