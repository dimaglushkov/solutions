package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/count-ways-to-build-good-strings/

func countGoodStrings(low int, high int, zero int, one int) int {
	res, mod := 0, 1000000007
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
		if i >= low {
			res = (res + dp[i]) % mod
		}
	}
	return res
}

func main() {
	testCases := []struct {
		low  int
		high int
		zero int
		one  int
		want int
	}{
		{
			low:  3,
			high: 3,
			zero: 1,
			one:  1,
			want: 8,
		},
		{
			low:  2,
			high: 3,
			zero: 1,
			one:  2,
			want: 5,
		},
	}

	for _, tc := range testCases {
		x := countGoodStrings(tc.low, tc.high, tc.zero, tc.one)
		status := "ERROR"
		// intentionally using this way of comparison
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}

}
