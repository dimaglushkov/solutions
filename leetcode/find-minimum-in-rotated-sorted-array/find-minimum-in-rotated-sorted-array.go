package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

// obvious straight-forward solution :)
func findMin_(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func findMin(nums []int) (res int) {
	if nums[0] < nums[len(nums)-1] || len(nums) == 1 {
		return nums[0]
	}

	res = nums[0]
	l, m, r := 1, 0, len(nums)-1

	for l <= r {
		m = (l + r) >> 1
		if nums[m] < res {
			res = nums[m]
		}
		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	return res
}

func main() {
	// Example 4
	var nums4 = []int{5, 1, 2, 3, 4}
	fmt.Println("Expected: 1	Output: ", findMin(nums4))

	// Example 1
	var nums1 = []int{3, 4, 5, 1, 2}
	fmt.Println("Expected: 1	Output: ", findMin(nums1))

	// Example 2
	var nums2 = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("Expected: 0	Output: ", findMin(nums2))

	// Example 3
	var nums3 = []int{11, 13, 15, 17}
	fmt.Println("Expected: 11	Output: ", findMin(nums3))

}
