package main

import "fmt"

// source: https://leetcode.com/problems/pascals-triangle/

func generate(n int) [][]int {
	res := make([][]int, n)

	res[0] = []int{1}
	for i := 1; i < n; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		prevI := i - 1
		for j := 1; j < i; j++ {
			res[i][j] = res[prevI][j] + res[prevI][j-1]
		}
	}
	return res
}

func main() {
	// Example 1
	var numRows1 int = 5
	fmt.Println("Expected: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]	Output: ", generate(numRows1))

	// Example 2
	var numRows2 int = 1
	fmt.Println("Expected: [[1]]	Output: ", generate(numRows2))

}
