package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/edit-distance/
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < m+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1])
			}
		}
	}
	return dp[n][m]
}

func main() {
	testCases := []struct {
		word1 string
		word2 string
		want  int
	}{
		{
			word1: "a",
			word2: "b",
			want:  2,
		},
		{
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			word1: "intention",
			word2: "execution",
			want:  5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minDistance(tc.word1, tc.word2)
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
