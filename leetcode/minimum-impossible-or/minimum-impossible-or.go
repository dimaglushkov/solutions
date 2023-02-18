package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/minimum-impossible-or/

func minImpossibleOR(nums []int) int {
	m := make(map[int]bool)
	ans := 1
	for _, i := range nums {
		m[i] = true
	}

	for m[ans] {
		ans *= 2
	}
	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 1},
			want: 4,
		},
		{
			nums: []int{5, 3, 2},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minImpossibleOR(tc.nums)
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
