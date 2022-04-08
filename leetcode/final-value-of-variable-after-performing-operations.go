package main

import "fmt"

// source: https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

func finalValueAfterOperations(operations []string) int {
	var x int
	for _, op := range operations {
		if op[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}

func main() {
	// Example 1
	var operations1 = []string{"--X", "X++", "X++"}
	fmt.Println("Expected: 1	Output: ", finalValueAfterOperations(operations1))

	// Example 2
	var operations2 = []string{"++X", "++X", "X++"}
	fmt.Println("Expected: 3	Output: ", finalValueAfterOperations(operations2))

	// Example 3
	var operations3 = []string{"X++", "++X", "--X", "X--"}
	fmt.Println("Expected: 0	Output: ", finalValueAfterOperations(operations3))

}
