package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}

	l, r := 0, len(nums)
	for l < r {
		h := int(uint(l+r) >> 1)
		if nums[h] < target {
			l = h + 1
		} else {
			r = h
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	var id int
	if id = binarySearch(nums, target); id == -1 || nums[id] != target {
		return []int{-1, -1}
	}
	l, r := id, id
	for l >= 0 && nums[l] == target {
		l--
	}
	for r < len(nums) && nums[r] == target {
		r++
	}
	return []int{l + 1, r - 1}
}

func main() {
	// Example 4
	var nums4 = []int{1, 2, 3}
	var target4 int = 2
	fmt.Println("Expected: [1,1]	Output: ", searchRange(nums4, target4))

	// Example 1
	var nums1 = []int{5, 7, 7, 8, 8, 10}
	var target1 int = 8
	fmt.Println("Expected: [3,4]	Output: ", searchRange(nums1, target1))

	// Example 2
	var nums2 = []int{5, 7, 7, 8, 8, 10}
	var target2 int = 6
	fmt.Println("Expected: [-1,-1]	Output: ", searchRange(nums2, target2))

	// Example 3
	var nums3 = []int{}
	var target3 int = 0
	fmt.Println("Expected: [-1,-1]	Output: ", searchRange(nums3, target3))

}
