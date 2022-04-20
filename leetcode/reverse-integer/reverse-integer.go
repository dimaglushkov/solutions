package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/reverse-integer/

func reverse(x int) (res int) {
	var zeros int
	for d := 0; x != 0; x = x / 10 {
		d = x % 10
		if d == 0 {
			if res != 0 {
				zeros++
			}
			continue
		}
		res *= int(math.Pow10(1 + zeros))
		res += d
		zeros = 0
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return
}

func main() {
	// Example 4
	var x5 = 24077
	fmt.Println("Expected: 77042 Output: ", reverse(x5))

	// Example 3
	var x3 int = 120
	fmt.Println("Expected: 21	Output: ", reverse(x3))

	// Example 1
	var x1 int = 123
	fmt.Println("Expected: 321	Output: ", reverse(x1))

	// Example 2
	var x2 int = -123
	fmt.Println("Expected: -321	Output: ", reverse(x2))

	// Example 4
	var x4 int = -12
	fmt.Println("Expected: -21	Output: ", reverse(x4))

}
