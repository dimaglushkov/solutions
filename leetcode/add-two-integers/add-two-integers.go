package main

import "fmt"

// source: https://leetcode.com/problems/add-two-integers/

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	// Example 1
	var num11 int = 12
	var num21 int = 5
	fmt.Println("Expected: 17	Output: ", sum(num11, num21))

	// Example 2
	var num12 int = -10
	var num22 int = 4
	fmt.Println("Expected: -6	Output: ", sum(num12, num22))

}
