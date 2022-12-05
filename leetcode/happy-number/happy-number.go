package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/happy-number/

func isHappy(n int) bool {
	m := make(map[int]bool)
	replace := func(x int) int {
		res := 0
		for x > 0 {
			res += (x % 10) * (x % 10)
			x /= 10
		}
		return res
	}
	for !m[n] && n != 1 {
		m[n] = true
		n = replace(n)
	}
	return n == 1
}

func main() {
	testCases := []struct {
		n    int
		want bool
	}{
		{
			n:    19,
			want: true,
		},
		{
			n:    2,
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isHappy(tc.n)
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
