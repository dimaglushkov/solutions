package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}

	return ans
}

func main() {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxProfit(tc.prices)
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
