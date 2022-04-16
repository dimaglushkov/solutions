package main

import "fmt"

// source: https://leetcode.com/problems/climbing-stairs/

// Since you can only make 1 or 2 steps,
// you can climb to n[i] stair only from stairs n[i]-1 and n[i]-2
func climbStairs(n int) (res int) {
	var x, y = 1, 2

	if n == 1 {
		return x
	}
	if n == 2 {
		return y
	}

	for i := 2; i < n; i++ {
		res = x + y
		x = y
		y = res
	}
	return
}

func main() {
	// Example 1
	var n1 int = 2
	fmt.Println("Expected: 2	Output: ", climbStairs(n1))

	// Example 2
	var n2 int = 3
	fmt.Println("Expected: 3	Output: ", climbStairs(n2))

	// Example 3
	var n3 int = 1
	fmt.Println("Expected: 1	Output: ", climbStairs(n3))

	// Example 4
	var n4 int = 5
	fmt.Println("Expected: 5 	Output: ", climbStairs(n4))

}
