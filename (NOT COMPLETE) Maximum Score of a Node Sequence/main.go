package main

import (
	"fmt"
)

func main()  {
	res := maximumScore([]int{5,2,9,8,4}, [][]int{{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}})

	fmt.Println(res)
}

// * ChatGPT: https://chatgpt.com/c/671d3151-ecd4-800c-8ba4-f55ebb5a85ab
// * Problem: https://leetcode.com/problems/maximum-score-of-a-node-sequence

func maximumScore(scores []int, edges [][]int) int{
	routes := make(map[int][]int)

	resArr := make([]int, 0, 8)
	res := 0

	for _, v := range edges {
		routes[v[0]] = append(routes[v[0]], v[1])
	}

	for k, v := range routes {
		route := make([]int, 0, 4)
		for i := 0; i < 4; i++ {



		}
		sum := 0
		for _, v := range route {
			sum += v
		}
		resArr = append(resArr, sum)
	}
	for _, v := range resArr {
		if v > res {
			res = v
		}
	}

	return res
}
