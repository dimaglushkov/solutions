package main

import "fmt"

// source: https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

func numberOfSteps(num int) int {
	var cnt int

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		cnt++
	}

	return cnt
}

func main() {
	// Example 1 
	var num1 int = 14
	fmt.Println("Expected: 6	Output: ", numberOfSteps(num1))

	// Example 2 
	var num2 int = 8
	fmt.Println("Expected: 4	Output: ", numberOfSteps(num2))

	// Example 3 
	var num3 int = 123
	fmt.Println("Expected: 12	Output: ", numberOfSteps(num3))

}
