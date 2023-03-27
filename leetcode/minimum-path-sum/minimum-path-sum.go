package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/minimum-path-sum/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = 1<<31 - 1
		}
	}
	dp[0][0] = 0
	for i := range dp {
		for j := range dp[i] {
			if ni := i - 1; ni >= 0 {
				dp[i][j] = min(dp[i][j], dp[ni][j])
			}
			if nj := j - 1; nj >= 0 {
				dp[i][j] = min(dp[i][j], dp[i][nj])
			}
			dp[i][j] += grid[i][j]
		}
	}

	return dp[n-1][m-1]
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want: 7,
		},
		{
			grid: [][]int{{1, 2, 3}, {4, 5, 6}},
			want: 12,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minPathSum(tc.grid)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
