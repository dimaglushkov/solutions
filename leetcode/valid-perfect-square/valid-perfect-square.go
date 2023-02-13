package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/valid-perfect-square/

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	i, j := 0, num
	for i < j {
		h := int(uint(i+j) >> 1)
		if h*h < num {
			i = h + 1
		} else {
			j = h
		}
	}
	return i*i == num
}

func main() {
	testCases := []struct {
		num  int
		want bool
	}{
		{
			num:  16,
			want: true,
		},
		{
			num:  14,
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isPerfectSquare(tc.num)
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
