package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/count-the-number-of-fair-pairs/

func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	var ans int64
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	for _, x := range nums {
		li, ui := sort.SearchInts(sortedNums, lower-x), sort.SearchInts(sortedNums, upper-x)
		if x >= lower-x && x <= upper-x {
			ans--
		}
		for ui < n && sortedNums[ui] == upper-x {
			ui++
		}
		ans += int64(ui - li)
	}

	return ans / 2
}

func main() {
	testCases := []struct {
		nums  []int
		lower int
		upper int
		want  int64
	}{
		{
			nums:  []int{0, 1, 7, 4, 4, 5},
			lower: 3,
			upper: 6,
			want:  6,
		},
		{
			nums:  []int{1, 7, 9, 2, 5},
			lower: 11,
			upper: 11,
			want:  1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countFairPairs(tc.nums, tc.lower, tc.upper)
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
