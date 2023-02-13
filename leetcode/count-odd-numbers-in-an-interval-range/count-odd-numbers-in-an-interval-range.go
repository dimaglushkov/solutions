package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/

func countOdds(low int, high int) int {
	if low&1 == 1 {
		low--
	}
	if high&1 == 1 {
		high++
	}
	return (high - low) / 2
}

func main() {
	testCases := []struct {
		low  int
		high int
		want int
	}{
		{
			low:  3,
			high: 7,
			want: 3,
		},
		{
			low:  8,
			high: 10,
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countOdds(tc.low, tc.high)
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
