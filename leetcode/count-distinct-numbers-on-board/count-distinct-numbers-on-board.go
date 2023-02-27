package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/count-distinct-numbers-on-board/

func distinctIntegers(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    5,
			want: 4,
		},
		{
			n:    3,
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := distinctIntegers(tc.n)
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
