package main

import "fmt"

// source: https://leetcode.com/problems/n-th-tribonacci-number/

func tribonacci(n int) (res int) {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}

	var n1, n2, n3 = 0, 1, 1

	for i := 2; i < n; i++ {
		res = n1 + n2 + n3
		n1 = n2
		n2 = n3
		n3 = res
	}
	return
}

func main() {
	// Example 1
	var n1 int = 4
	fmt.Println("Expected: 4	Output: ", tribonacci(n1))

	// Example 2
	var n2 int = 25
	fmt.Println("Expected: 1389537	Output: ", tribonacci(n2))

}
