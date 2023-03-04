package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var ans int64
	lastMin, lastMax := -1, -1
	leftmost := 0
	for i, d := range nums {
		if d < minK || d > maxK {
			leftmost = i + 1
			lastMin = -1
			lastMax = -1
			continue
		}
		if d == minK {
			lastMin = i
		}
		if d == maxK {
			lastMax = i
		}
		if lastMin != -1 && lastMax != -1 {
			ans += int64(min(lastMin, lastMax) - leftmost + 1)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		minK int
		maxK int
		want int64
	}{

		{
			nums: []int{4, 4, 1, 3, 9},
			minK: 1,
			maxK: 9,
			want: 3,
		},
		{
			nums: []int{1, 4, 9, 9, 8, 9, 1, 9, 3, 9, 1, 2, 7, 5, 6},
			minK: 1,
			maxK: 9,
			want: 81,
		},
		{
			nums: []int{1, 1, 1, 1},
			minK: 1,
			maxK: 1,
			want: 10,
		},
		{
			nums: []int{1, 3, 5, 2, 7, 5},
			minK: 1,
			maxK: 5,
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countSubarrays(tc.nums, tc.minK, tc.maxK)
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
