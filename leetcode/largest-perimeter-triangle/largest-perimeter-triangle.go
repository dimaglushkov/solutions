package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/largest-perimeter-triangle/

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i > 1; i-- {
		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i-1] + nums[i-2] + nums[i]
		}
	}

	return 0
}

func main() {
	// Example 1
	var nums1 = []int{2, 1, 2}
	fmt.Println("Expected: 5	Output: ", largestPerimeter(nums1))

	// Example 2
	var nums2 = []int{1, 2, 1}
	fmt.Println("Expected: 0	Output: ", largestPerimeter(nums2))

}
