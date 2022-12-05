package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/add-digits/

func addDigits(x int) int {
	for x/10 > 0 {
		nx := 0
		for x > 0 {
			nx += x % 10
			x /= 10
		}
		x = nx
	}
	return x
}

func main() {
	testCases := []struct {
		num  int
		want int
	}{
		{
			num:  38,
			want: 2,
		},
		{
			num:  0,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := addDigits(tc.num)
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
