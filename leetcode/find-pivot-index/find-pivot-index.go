package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-pivot-index/

func pivotIndex(nums []int) int {
	n := len(nums)
	pre, suf := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] + nums[i-1]
		suf[n-i-1] = suf[n-i] + nums[n-i]
	}
	for i := range nums {
		if pre[i] == suf[i] {
			return i
		}
	}
	return -1
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			nums: []int{2, 1, -1},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := pivotIndex(tc.nums)
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
