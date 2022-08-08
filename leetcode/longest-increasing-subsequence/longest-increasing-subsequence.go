package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	var n = len(nums)
	var dp = make([]int, n)
	var res int

	for i := n - 1; i >= 0; i-- {
		var t, dt = n, 0
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] && dp[j] > dt {
				t = j
				dt = dp[j]
			}
		}
		if t == n {
			dp[i] = 1
			continue
		}
		dp[i] += dp[t] + 1
	}

	for _, x := range dp {
		if x > res {
			res = x
		}
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Expected: 4	Output: ", lengthOfLIS(nums1))

	// Example 2
	var nums2 = []int{0, 1, 0, 3, 2, 3}
	fmt.Println("Expected: 4	Output: ", lengthOfLIS(nums2))

	// Example 3
	var nums3 = []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println("Expected: 1	Output: ", lengthOfLIS(nums3))

}
