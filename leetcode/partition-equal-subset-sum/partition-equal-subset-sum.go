package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/partition-equal-subset-sum/

func canPartition(nums []int) bool {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%2 == 1 {
		return false
	}

	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sum]
}

func _canPartition(nums []int) bool {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%2 == 1 {
		return false
	}

	sum /= 2
	n := len(nums)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][sum]
}

func main() {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			nums: []int{1, 2, 5},
			want: false,
		},
		{
			nums: []int{1, 2, 3, 5},
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canPartition(tc.nums)
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
