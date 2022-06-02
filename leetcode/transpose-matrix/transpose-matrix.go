package main

import "fmt"

// source: https://leetcode.com/problems/transpose-matrix/

func transpose(matrix [][]int) [][]int {
	var n, m = len(matrix), len(matrix[0])
	if n == m {
		for i := 0; i < n; i++ {
			for j := i + 1; j < m; j++ {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
		return matrix
	}
	newMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		newMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}
	return newMatrix
}

func main() {
	// Example 1
	var matrix1 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Expected: [[1,4,7],[2,5,8],[3,6,9]]	Output: ", transpose(matrix1))

	// Example 2
	var matrix2 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Expected: [[1,4],[2,5],[3,6]]	Output: ", transpose(matrix2))

}
