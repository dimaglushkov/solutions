package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/arranging-coins/

func arrangeCoins(n int) int {
	ans := 0
	for i := 1; n > 0; i++ {
		ans++
		n -= i
	}
	if n < 0 {
		ans--
	}
	return ans
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    1,
			want: 1,
		},
		{
			n:    5,
			want: 2,
		},
		{
			n:    8,
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := arrangeCoins(tc.n)
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
