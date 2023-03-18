package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/min-cost-climbing-stairs/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	n := len(cost)
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return dp[n-1]
}

func main() {
	testCases := []struct {
		cost []int
		want int
	}{
		{
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minCostClimbingStairs(tc.cost)
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
