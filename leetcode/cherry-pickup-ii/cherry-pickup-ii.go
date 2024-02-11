package main

import (
	"fmt"
)

func cherryPickup(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, m)
		}
	}

	ans := 0
	dp[0][0][m-1] = grid[0][0] + grid[0][m-1]

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				if j > i || k < m-i-1 || j > k {
					continue
				}

				dp[i][j][k] = dp[i-1][j][k]
				if k > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k-1])
				}
				if k+1 < m {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k+1])
				}

				if j > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k])
					if k > 0 {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k-1])
					}
					if k+1 < m {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k+1])
					}
				}
				if j+1 < m {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j+1][k])
					if k > 0 {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j+1][k-1])
					}
					if k+1 < m {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j+1][k+1])
					}
				}

				if j != k {
					dp[i][j][k] += grid[i][j] + grid[i][k]
				} else {
					dp[i][j][k] += grid[i][j]
				}

				ans = max(ans, dp[i][j][k])
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}},
			want: 24,
		},
		{
			grid: [][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}},
			want: 28,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := cherryPickup(tc.grid)
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
