package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/single-element-in-a-sorted-array/

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if (h%2 == 0 && h+1 < n && nums[h] == nums[h+1]) || (h%2 == 1 && h-1 >= 0 && nums[h] == nums[h-1]) { // if f(h) less
			i = h + 1
		} else {
			j = h
		}
	}
	return nums[i]
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{0},
			want: 0,
		},
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			want: 2,
		},
		{
			nums: []int{3, 3, 7, 7, 10, 11, 11},
			want: 10,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := singleNonDuplicate(tc.nums)
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
