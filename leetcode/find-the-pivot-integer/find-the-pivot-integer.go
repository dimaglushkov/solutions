package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-the-pivot-integer/

func pivotInteger(n int) int {
	ps := make([]int, n+1)
	ps[0] = 0
	for i := 1; i <= n; i++ {
		ps[i] = ps[i-1] + i
	}

	for i := 1; i <= n; i++ {
		if ps[n]-ps[i]+i == ps[i] {
			return i
		}
	}
	return -1
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    8,
			want: 6,
		},
		{
			n:    1,
			want: 1,
		},
		{
			n:    4,
			want: -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := pivotInteger(tc.n)
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
