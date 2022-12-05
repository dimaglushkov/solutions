package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		res ^= nums[i]
	}
	return res
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			nums: []int{1},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := singleNumber(tc.nums)
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
