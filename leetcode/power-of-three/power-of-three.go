package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/power-of-three/

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func main() {
	// Example 1
	var n1 int = 27
	fmt.Println("Expected: true	Output: ", isPowerOfThree(n1))

	// Example 2
	var n2 int = 0
	fmt.Println("Expected: false	Output: ", isPowerOfThree(n2))

	// Example 3
	var n3 int = 9
	fmt.Println("Expected: true	Output: ", isPowerOfThree(n3))

}
