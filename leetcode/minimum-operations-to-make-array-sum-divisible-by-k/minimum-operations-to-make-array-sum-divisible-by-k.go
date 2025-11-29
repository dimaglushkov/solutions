package main

// source: https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/
func minOperations(nums []int, k int) int {
	sum := 0

	for i := range nums {
		sum += nums[i]
	}

	return (sum - k) % k
}
