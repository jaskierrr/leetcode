package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	left := 0
	right :=len(nums)-1

	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		mid := (left+right) / 2

		if nums[mid] == target {
			return mid
		}

		if target > mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
