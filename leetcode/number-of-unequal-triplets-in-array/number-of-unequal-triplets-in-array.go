package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/number-of-unequal-triplets-in-array/

func unequalTriplets(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] != nums[j] {
				for k := j; k < n; k++ {
					if nums[i] != nums[k] && nums[j] != nums[k] {
						res++
					}
				}
			}

		}
	}
	return res
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{4, 4, 2, 4, 3},
			want: 3,
		},
		{
			nums: []int{1, 1, 1, 1, 1},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := unequalTriplets(tc.nums)
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
