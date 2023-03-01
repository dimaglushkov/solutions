package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/maximum-length-of-repeated-subarray/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{0, 1, 1, 1, 1},
			nums2: []int{1, 0, 1, 0, 1},
			want:  2,
		},
		{
			nums1: []int{0, 0, 0, 0, 0},
			nums2: []int{0, 0, 0, 0, 0},
			want:  5,
		},
		{
			nums1: []int{1, 2, 3, 2, 1},
			nums2: []int{3, 2, 1, 4, 7},
			want:  3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findLength(tc.nums1, tc.nums2)
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
