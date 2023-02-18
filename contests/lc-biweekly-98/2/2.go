package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/minimum-score-by-changing-two-elements/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimizeSum(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return min(nums[n-2]-nums[1], min(nums[n-1]-nums[2], nums[n-3]-nums[0]))
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 4, 3},
			want: 0,
		},
		{
			nums: []int{1, 4, 7, 8, 5},
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimizeSum(tc.nums)
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
