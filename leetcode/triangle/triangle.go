package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/triangle/

// obv too slow
func minimumTotal_(triangle [][]int) int {
	var res = math.MaxInt
	var util func(i, j, sum int)
	util = func(i, j, sum int) {
		if i == len(triangle) {
			if sum < res {
				res = sum
			}
			return
		}
		sum += triangle[i][j]
		util(i+1, j, sum)
		util(i+1, j+1, sum)
	}
	util(0, 0, 0)
	return res
}

func minimumTotal(triangle [][]int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	n := len(triangle)
	memo := make([]int, n)
	copy(memo, triangle[n-1])

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			memo[j] = min(memo[j], memo[j+1]) + triangle[i][j]
		}
	}

	return memo[0]
}

func main() {
	// Example 0
	var triangle0 = [][]int{{-1}, {2, 3}, {1, -1, -3}}
	fmt.Println("Expected: -1	Output: ", minimumTotal(triangle0))

	// Example 1
	var triangle1 = [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println("Expected: 11	Output: ", minimumTotal(triangle1))

	// Example 2
	var triangle2 = [][]int{{-10}}
	fmt.Println("Expected: -10	Output: ", minimumTotal(triangle2))

}
