package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// source: https://leetcode.com/problems/maximum-alternating-sum-of-squares/
func maxAlternatingSum(nums []int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})
	ans := int64(0)
	n := len(nums)

	for i := 0; i < n/2; i++ {
		ans += int64(nums[i] * nums[i])
		ans -= int64(nums[n-i-1] * nums[n-i-1])
	}

	if n%2 == 1 {
		ans += int64(nums[n/2] * nums[n/2])
	}

	return ans
}

func main() {
	fmt.Println(maxAlternatingSum([]int{19, -20}))
	fmt.Println(maxAlternatingSum([]int{1, 2, 3}))
	fmt.Println(maxAlternatingSum([]int{1, -1, 2, -2, 3, -3}))
}
