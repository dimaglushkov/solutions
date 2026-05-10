package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = -1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] != -1 && abs(nums[i]-nums[j]) <= target && dp[j] >= dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	return dp[n-1]
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{0, 2, 1, 3},
			target: 1,
			want:   -1,
		},
		{
			nums:   []int{1, 3, 6, 4, 1, 2},
			target: 2,
			want:   3,
		},
		{
			nums:   []int{1, 3, 6, 4, 1, 2},
			target: 3,
			want:   5,
		},
		{
			nums:   []int{1, 3, 6, 4, 1, 2},
			target: 0,
			want:   -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumJumps(tc.nums, tc.target)
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
