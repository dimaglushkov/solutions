package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/

func check(nums []int) bool {
	n := len(nums)
	i := 1
	for i < n && nums[i-1] <= nums[i] {
		i++
	}
	sorted := sort.IntsAreSorted(nums[:i]) && sort.IntsAreSorted(nums[i:])
	if i < n {
		return sorted && (nums[0] >= nums[n-1])
	}
	return sorted
}

func main() {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 1, 1},
			want: true,
		},
		{
			nums: []int{3, 4, 5, 1, 2},
			want: true,
		},
		{
			nums: []int{2, 1, 3, 4},
			want: false,
		},
		{
			nums: []int{1, 2, 3},
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := check(tc.nums)
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
