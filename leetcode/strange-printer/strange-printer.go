package main

import (
	"fmt"
)

func strangePrinter(s string) int {
	var minB []byte

	for i := 0; i < len(s); i++ {
		if len(minB) == 0 || minB[len(minB)-1] != s[i] {
			minB = append(minB, s[i])
		}
	}

	s = string(minB)
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for l := 2; l < n+1; l++ {
		for b := 0; b < n-l+1; b++ {
			e := b + l - 1
			dp[b][e] = l

			for off := 0; off < l-1; off++ {
				t := dp[b][b+off] + dp[b+off+1][e]

				if s[b+off] == s[e] {
					t--
				}

				dp[b][e] = min(dp[b][e], t)
			}
		}
	}

	return dp[0][n-1]

}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "aaaaaaaaaaaaaaaa",
			want: 1,
		},
		{
			s:    "a",
			want: 1,
		},
		{
			s:    "leetcode",
			want: 6,
		},
		{
			s:    "aaabbb",
			want: 2,
		},
		{
			s:    "aba",
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := strangePrinter(tc.s)
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
