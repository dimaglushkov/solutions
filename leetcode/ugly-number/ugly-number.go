package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/ugly-number/

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		n    int
		want bool
	}{
		{
			n:    -2147483648,
			want: false,
		},
		{
			n:    6,
			want: true,
		},
		{
			n:    1,
			want: true,
		},
		{
			n:    14,
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isUgly(tc.n)
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
