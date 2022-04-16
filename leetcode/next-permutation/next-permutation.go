package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/next-permutation/

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l := len(nums) - 2
	for l >= 0 && nums[l] >= nums[l+1] {
		l--
	}

	if l >= 0 {
		r := len(nums) - 1
		for r > 0 && nums[l] >= nums[r] {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}

	reverse(nums[l+1:])
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3}
	nextPermutation(nums1)
	fmt.Println("Expected: [1 3 2]	Output: ", nums1)

	// Example 2
	var nums2 = []int{3, 2, 1}
	nextPermutation(nums2)
	fmt.Println("Expected: [1 2 3]	Output: ", nums2)

	// Example 3
	var nums3 = []int{1, 1, 5}
	nextPermutation(nums3)
	fmt.Println("Expected: [1 5 1]	Output: ", nums3)

}
