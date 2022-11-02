package main

import "fmt"

func isIn(po []int, poy []int, r int) bool {
	if (po[0]-poy[0])*(po[0]-poy[0])+(po[1]-poy[1])*(po[1]-poy[1]) <= r*r {
		return true
	} else {
		return false
	}

}

func countPoints(points [][]int, queries [][]int) []int {

	resp := make([]int, 0)

	for _, query := range queries {
		j := 0
		for i2 := range points {
			if isIn(points[i2], []int{query[0], query[1]}, query[2]) {
				j++
			}
		}
		resp = append(resp, j)
	}
	return resp
}

func main() {
	points := [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}

	queries := [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}

	n := countPoints(points, queries)

	fmt.Println(n)

}
