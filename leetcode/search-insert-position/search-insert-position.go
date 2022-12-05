package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := searchInsert(tc.nums, tc.target)
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
