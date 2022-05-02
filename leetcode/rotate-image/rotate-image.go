package main

import "fmt"

// source: https://leetcode.com/problems/rotate-image/

func rotate(matrix [][]int) {
	var n = len(matrix)
	for i := 0; i < len(matrix[0]); i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
func main() {
	// Example 1
	var matrix1 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix1)
	fmt.Println("Expected: [[7,4,1],[8,5,2],[9,6,3]]	Output: ", matrix1)

	// Example 2
	var matrix2 = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix2)
	fmt.Println("Expected: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]	Output: ", matrix2)

}
