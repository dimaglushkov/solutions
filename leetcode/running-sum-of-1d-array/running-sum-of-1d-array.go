package main

import "fmt"

// source: https://leetcode.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + prev
		prev = nums[i]
	}
	return nums
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3, 4}
	fmt.Println("Expected: [1,3,6,10] Output: ", runningSum(nums1))

	// Example 2
	var nums2 = []int{1, 1, 1, 1, 1}
	fmt.Println("Expected: [1,2,3,4,5] Output: ", runningSum(nums2))

	// Example 3
	var nums3 = []int{3, 1, 2, 10, 1}
	fmt.Println("Expected: [3,4,6,16,17] Output: ", runningSum(nums3))

	// Example 4
	var nums4 = []int{3}
	fmt.Println("Expected: [3] Output: ", runningSum(nums4))

}
