package main

import "sort"

// source: https://leetcode.com/problems/check-if-array-is-good/
func isGood(nums []int) bool {
	n := len(nums) - 1
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return false
		}
	}

	return nums[len(nums)-1] == n
}

func main() {
	isGood([]int{1, 3, 3, 2})
}
