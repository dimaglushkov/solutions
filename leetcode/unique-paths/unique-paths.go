package main

import "fmt"

// source: https://leetcode.com/problems/unique-paths/
var mem [100][100]int

func init() {
	mem[99][99] = 1
	for i := 99; i >= 0; i-- {
		for j := 99; j >= 0; j-- {
			if i < 99 {
				mem[i][j] += mem[i+1][j]
			}
			if j < 99 {
				mem[i][j] += mem[i][j+1]
			}
		}
	}
}

func uniquePaths(m int, n int) int {
	return mem[100-m][100-n]
}

func main() {
	// Example 1
	var m1 int = 3
	var n1 int = 7
	fmt.Println("Expected: 28	Output: ", uniquePaths(m1, n1))

	// Example 2
	var m2 int = 3
	var n2 int = 2
	fmt.Println("Expected: 3	Output: ", uniquePaths(m2, n2))

}
