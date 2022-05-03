package main

import "fmt"

// source: https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var l, r = len(nums) - 1, 0

	var prev = nums[r]
	for i := 1; i < len(nums); i++ {
		if nums[i] < prev {
			r = i
		} else {
			prev = nums[i]
		}
	}

	prev = nums[l]
	for i := len(nums) - 1; i >= 0; i-- {
		if prev < nums[i] {
			l = i
		} else {
			prev = nums[i]
		}
	}
	if r == 0 {
		return 0
	}
	return r - l + 1
}

func main() {
	// Example 1
	var nums1 = []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println("Expected: 5	Output: ", findUnsortedSubarray(nums1))

	// Example 2
	var nums2 = []int{1, 2, 3, 4}
	fmt.Println("Expected: 0	Output: ", findUnsortedSubarray(nums2))

	// Example 3
	var nums3 = []int{1}
	fmt.Println("Expected: 0	Output: ", findUnsortedSubarray(nums3))

}
