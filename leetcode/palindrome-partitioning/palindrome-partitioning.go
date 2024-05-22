package main

import (
	"fmt"
)

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				if l == 2 || dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
		}
	}

	var ans [][]string

	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {
		if start == n {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for end := start; end < n; end++ {
			if dp[start][end] {
				backtrack(end+1, append(path, s[start:end+1]))
			}
		}
	}

	backtrack(0, []string{})
	return ans
}

func main() {
	testCases := []struct {
		s    string
		want [][]string
	}{
		{
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			s:    "a",
			want: [][]string{{"a"}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := partition(tc.s)
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
