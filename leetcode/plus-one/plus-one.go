package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {
	n := len(digits)
	x := false
	if digits[n-1] == 9 {
		digits[n-1] = 0
		x = true
	} else {
		digits[n-1]++
		return digits
	}
	for i := n - 2; i >= 0 && x; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			x = true
		} else {
			digits[i]++
			x = false
		}
	}
	if x {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	testCases := []struct {
		digits []int
		want   []int
	}{
		{
			digits: []int{9, 9},
			want:   []int{1, 0, 0},
		},
		{
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			digits: []int{9},
			want:   []int{1, 0},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := plusOne(tc.digits)
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
