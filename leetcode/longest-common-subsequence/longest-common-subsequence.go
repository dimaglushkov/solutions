package main

import (
	"fmt"
)
// source: https://leetcode.com/problems/longest-common-subsequence/

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func longestCommonSubsequence(t1 string, t2 string) int {
	n1, n2 := len(t1), len(t2)
	dp := make([][]int, n1+1)
    dp[0] = make([]int, n2+1)
	for i := 1; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
		for j := 1; j <= n2; j++ {
			if t1[i-1] == t2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]

}

func main() {
    testCases := []struct {
		text1 string
		text2 string
		want int
    }{
		{
			text1:  "abcde",
			text2:  "ace" ,
			want:  3  ,
		},
		{
			text1:  "abc",
			text2:  "abc",
			want:  3,
		},
		{
			text1:  "abc",
			text2:  "def",
			want:  0,
		},
	}


    successes := 0
    for _, tc := range testCases {
        x := longestCommonSubsequence(tc.text1, tc.text2)
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
        fmt.Printf("===\nFAIL: %d tests failed\n", l - successes)
    }

}
