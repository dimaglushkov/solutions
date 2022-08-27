package main

import "fmt"

// source: https://leetcode.com/problems/power-of-four/

func isPowerOfFour(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n&0x55555555) != 0
}

func main() {
	// Example 1
	var n1 int = 16
	fmt.Println("Expected: true	Output: ", isPowerOfFour(n1))

	// Example 2
	var n2 int = 5
	fmt.Println("Expected: false	Output: ", isPowerOfFour(n2))

	// Example 3
	var n3 int = 1
	fmt.Println("Expected: true	Output: ", isPowerOfFour(n3))

}
