package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/maximize-greatness-of-an-array/

func maximizeGreatness(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	l, r := 0, 1
	for r < n {
		if nums[l] < nums[r] {
			l++
			r++
			ans++
		} else {
			r++
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 3, 5, 2, 1, 3, 1},
			want: 4,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximizeGreatness(tc.nums)
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
