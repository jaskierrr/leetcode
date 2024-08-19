package main

import "fmt"

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}

func average(salary []int) float64 {
	max := salary[0]
	min := salary[0]
	sum := 0

	for _, v := range salary {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}

		sum += v
	}

	sum = sum - max - min

	return float64(sum) / float64(len(salary)-2)
}
