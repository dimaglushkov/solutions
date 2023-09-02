package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minExtraChar(s string, dictionary []string) int {
	var dp [51]int

	n := len(s)

	for i := n - 1; i >= 0; i-- {
		dp[i] = 1 + dp[i+1]

		for _, w := range dictionary {
			if i+len(w) <= n && s[i:i+len(w)] == w {
				dp[i] = min(dp[i], dp[i+len(w)])
			}
		}
	}

	return dp[0]
}

func main() {
	testCases := []struct {
		s          string
		dictionary []string
		want       int
	}{
		{
			s:          "leetscode",
			dictionary: []string{"leet", "code", "leetcode"},
			want:       1,
		},
		{
			s:          "sayhelloworld",
			dictionary: []string{"hello", "world"},
			want:       3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minExtraChar(tc.s, tc.dictionary)
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
