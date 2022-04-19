package main

import "fmt"

// source: https://leetcode.com/problems/fibonacci-number/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var res int
	x, y := 0, 1
	for i := 1; i < n; i++ {
		res = x + y
		x = y
		y = res
	}
	return res
}

func main() {
	// Example 1
	var n1 int = 2
	fmt.Println("Expected: 1	Output: ", fib(n1))

	// Example 2
	var n2 int = 3
	fmt.Println("Expected: 2	Output: ", fib(n2))

	// Example 3
	var n3 int = 4
	fmt.Println("Expected: 3	Output: ", fib(n3))

}
