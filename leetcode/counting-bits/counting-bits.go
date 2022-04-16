package main

import (
	"fmt"
	"math/bits"
)

// source: https://leetcode.com/problems/counting-bits/

func countBits1(n int) []int {
	res := make([]int, 0, n+1)

	for i := 0; i <= n; i++ {
		res = append(res, bits.OnesCount(uint(i)))
	}

	return res
}

// Instead of calculating a new value on every step, using previous values
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}

func main() {
	// Example 1
	var n1 int = 2
	fmt.Println("Expected: [0,1,1]	Output: ", countBits(n1))

	// Example 2
	var n2 int = 5
	fmt.Println("Expected: [0,1,1,2,1,2]	Output: ", countBits(n2))

}
