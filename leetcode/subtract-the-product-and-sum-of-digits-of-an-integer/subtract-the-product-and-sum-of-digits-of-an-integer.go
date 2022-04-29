package main

import "fmt"

// source: https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

func subtractProductAndSum(n int) int {
	var sum, prod, x int
	prod = 1
	for n > 0 {
		x = n % 10
		n /= 10
		sum += x
		prod *= x
	}
	return prod - sum
}

func main() {
	// Example 1
	var n1 int = 234
	fmt.Println("Expected: 15 	Output: ", subtractProductAndSum(n1))

	// Example 2
	var n2 int = 4421
	fmt.Println("Expected: 21	Output: ", subtractProductAndSum(n2))

}
