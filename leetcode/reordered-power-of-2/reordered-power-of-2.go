package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/reordered-power-of-2/

func reorderedPowerOf2(N int) bool {
	s := sign(N)

	for i := 0; i < 32; i++ {
		if sign(1<<i) == s {
			return true
		}
	}
	return false
}

func sign(n int) int {
	if n == 0 {
		return 0
	}
	l, r := n/10, n%10
	return sign(l) + int(math.Pow(10, float64(r)))
}

func main() {
	// Example 1
	var n1 int = 1
	fmt.Println("Expected: true	Output: ", reorderedPowerOf2(n1))

	// Example 2
	var n2 int = 10
	fmt.Println("Expected: false	Output: ", reorderedPowerOf2(n2))

}
