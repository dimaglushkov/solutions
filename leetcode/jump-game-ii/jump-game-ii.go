package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/jump-game-ii/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = 1 << 32
	}

	for i := range nums {
		for j := i; j < n && j-i <= nums[i]; j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}

	return dp[n-1]
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := jump(tc.nums)
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
