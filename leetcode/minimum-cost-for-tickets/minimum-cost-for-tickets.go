package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/minimum-cost-for-tickets/

func min(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] < x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)

	travelDays := make(map[int]bool)
	for _, d := range days {
		travelDays[d] = true
	}

	for i := 1; i <= lastDay; i++ {
		if !travelDays[i] {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = min(
			dp[i-1]+costs[0],
			dp[max(i-7, 0)]+costs[1],
			dp[max(i-30, 0)]+costs[2],
		)
	}

	return dp[lastDay]
}

func main() {
	testCases := []struct {
		days  []int
		costs []int
		want  int
	}{
		{
			days:  []int{1, 4, 6, 7, 8, 20},
			costs: []int{2, 7, 15},
			want:  11,
		},
		{
			days:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			costs: []int{2, 7, 15},
			want:  17,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := mincostTickets(tc.days, tc.costs)
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
