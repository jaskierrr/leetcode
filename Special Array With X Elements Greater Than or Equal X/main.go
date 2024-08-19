package main

import "fmt"

func main() {
	fmt.Println(specialArray([]int{0,4,3,0,4}))
}

func specialArray(nums []int) int {
	res := 0
	numMap := make(map[int]int, len(nums))

	for _, v := range nums {
		numMap[v]++
	}

	for i := 1; i <= len(nums); i++ {
		for k, v := range numMap {
			if k >= i {
				res += v
			}
		}
		if res == i {
			return res
		}
		res = 0
	}

	return -1
}
