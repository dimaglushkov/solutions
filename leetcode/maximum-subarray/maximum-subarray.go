package main

import "fmt"

// source: https://leetcode.com/problems/maximum-subarray/

// Using Kadane's algorithm
func maxSubArray(nums []int) int {
	var sum int
	var maxSum = nums[0]
	var l = len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func main() {
	// Example 1
	var nums1 = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Expected: 6	Output: ", maxSubArray(nums1))

	// Example 2
	var nums2 = []int{1}
	fmt.Println("Expected: 1	Output: ", maxSubArray(nums2))

	// Example 3
	var nums3 = []int{5, 4, -1, 7, 8}
	fmt.Println("Expected: 23	Output: ", maxSubArray(nums3))

}
