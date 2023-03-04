package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/count-total-number-of-colored-cells/

func coloredCells(n int) int64 {
	dp := make([]int64, n+1)
	dp[1] = 1
	for i := int64(2); i <= int64(n); i++ {
		dp[i] = dp[i-1] + (i-1)*4
	}
	return dp[n]
}

func main() {
	testCases := []struct {
		n    int
		want int64
	}{
		{
			n:    1,
			want: 1,
		},
		{
			n:    2,
			want: 5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := coloredCells(tc.n)
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
