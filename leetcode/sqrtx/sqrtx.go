package main

import "fmt"

// source: https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	low, mid, high := 1, 1, x
	for high-low > 1 {
		mid = low + (high-low)/2
		if mid*mid > x {
			high = mid
		} else {
			low = mid
		}
	}
	return low
}

func main() {
	// Example 1
	var x1 int = 4
	fmt.Println("Expected: 2	Output: ", mySqrt(x1))

	// Example 2
	var x2 int = 8
	fmt.Println("Expected: 2	Output: ", mySqrt(x2))

	// Example 3
	var x3 int = 10
	fmt.Println("Expected: 3	Output: ", mySqrt(x3))

	// Example 4
	var x4 int = 121
	fmt.Println("Expected: 11	Output: ", mySqrt(x4))

}
