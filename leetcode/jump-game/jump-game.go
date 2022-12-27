package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/jump-game/

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i, j := range nums {
		if !dp[i] {
			continue
		}
		for t := i + 1; t <= i+j && t < n; t++ {
			dp[t] = true
		}
	}
	return dp[n-1]
}

func main() {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canJump(tc.nums)
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
