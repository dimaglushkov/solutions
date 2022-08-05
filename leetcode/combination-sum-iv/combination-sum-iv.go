package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/combination-sum-iv/

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num > i {
				break
			}
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3}
	var target1 int = 4
	fmt.Println("Expected: 7	Output: ", combinationSum4(nums1, target1))

	// Example 2
	var nums2 = []int{9}
	var target2 int = 3
	fmt.Println("Expected: 0	Output: ", combinationSum4(nums2, target2))

}
