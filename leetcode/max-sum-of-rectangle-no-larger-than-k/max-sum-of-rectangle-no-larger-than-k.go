package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = matrix[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i][j]
		}
	}

	res := math.MinInt
	var A, B, C, D int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for x := i; x < m; x++ {
				for y := j; y < n; y++ {
					A, B, C, D = dp[x][y], 0, 0, 0
					if j-1 >= 0 {
						B = dp[x][j-1]
					}
					if i-1 >= 0 {
						C = dp[i-1][y]
					}

					if i-1 >= 0 && j-1 >= 0 {
						D = dp[i-1][j-1]
					}

					if area := A - B - C + D; area <= k && area > res {
						res = area
					}
				}
			}
		}
	}

	return res
}

func main() {
	// Example 1
	var matrix1 = [][]int{{1, 0, 1}, {0, -2, 3}}
	var k1 int = 2
	fmt.Println("Expected: 2	Output: ", maxSumSubmatrix(matrix1, k1))

	// Example 2
	var matrix2 = [][]int{{2, 2, -1}}
	var k2 int = 3
	fmt.Println("Expected: 3	Output: ", maxSumSubmatrix(matrix2, k2))

}
