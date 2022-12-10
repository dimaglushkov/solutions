package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/frog-jump-ii/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxJump(s []int) int {
	n := len(s)
	res := s[1] - s[0]
	for i := 2; i < n; i++ {
		res = max(res, s[i]-s[i-2])
	}
	return res
}

func main() {
	testCases := []struct {
		stones []int
		want   int
	}{
		{
			stones: []int{0, 2, 5, 6, 7},
			want:   5,
		},
		{
			stones: []int{0, 3, 9},
			want:   9,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxJump(tc.stones)
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
