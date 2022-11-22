package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-subarrays-with-equal-sum/

func findSubarrays(nums []int) bool {
	n := len(nums)
	sum := make(map[int]int, n)
	for i := 0; i < n-1; i++ {
		x := nums[i] + nums[i+1]
		if _, ok := sum[x]; ok {
			return true
		}
		sum[x] = i
	}
	return false
}

func main() {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{4, 2, 4},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			nums: []int{0, 0, 0},
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findSubarrays(tc.nums)
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
