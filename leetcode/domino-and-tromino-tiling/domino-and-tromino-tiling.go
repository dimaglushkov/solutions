package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/domino-and-tromino-tiling/

func numTilings(n int) int {
	dp := [1000]int{1, 2, 5}
	var calc func(x int) int
	calc = func(x int) int {
		if dp[x-1] == 0 {
			dp[x-1] = (2*calc(x-1) + calc(x-3)) % 1000000007
		}
		return dp[x-1]
	}
	return calc(n)
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 5,
		},
		{
			n:    1,
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numTilings(tc.n)
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
