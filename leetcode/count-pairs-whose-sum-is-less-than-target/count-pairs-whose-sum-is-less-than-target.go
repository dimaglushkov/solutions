package main

import (
	"fmt"
)

func countPairs(nums []int, target int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] < target {
				res++
			}
		}
	}

	return res
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 1, 2, 3, 1},
			target: 2,
			want:   3,
		},
		{
			nums:   []int{-6, 2, 5, -2, -7, -1, 3},
			target: -2,
			want:   10,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countPairs(tc.nums, tc.target)
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
