package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	b0 := -prices[0]
	b1 := b0
	s0, s1, s2 := 0, 0, 0

	for i := 1; i < len(prices); i++ {
		b0 = b1
		if b1 < s2-prices[i] {
			b0 = s2 - prices[i]
		}
		s0 = s1
		if s1 < b1+prices[i] {
			s0 = b1 + prices[i]
		}
		b1 = b0
		s2 = s1
		s1 = s0
	}

	return s0
}

func main() {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{1, 2, 3, 0, 2},
			want:   3,
		},
		{
			prices: []int{1},
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
