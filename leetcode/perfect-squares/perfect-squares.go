package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/perfect-squares/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 100000
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j])
		}
		dp[i]++
	}
	return dp[n]
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    12,
			want: 3,
		},
		{
			n:    13,
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numSquares(tc.n)
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
