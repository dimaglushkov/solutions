package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/max-number-of-k-sum-pairs/

func maxOperations(nums []int, k int) int {
	var l, r = 0, len(nums) - 1
	var res, val int
	sort.Ints(nums)
	for l < r {
		val = nums[l] + nums[r]
		switch {
		case val == k:
			res++
			l++
			r--
		case val < k:
			l++
		case val > k:
			r--
		}
	}
	return res
}

func main() {
	// Example 3
	var nums3 = []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}
	var k3 int = 3
	fmt.Println("Expected: 4	Output: ", maxOperations(nums3, k3))

	// Example 1
	var nums1 = []int{1, 2, 3, 4}
	var k1 int = 5
	fmt.Println("Expected: 2	Output: ", maxOperations(nums1, k1))

	// Example 2
	var nums2 = []int{3, 1, 3, 4, 3}
	var k2 int = 6
	fmt.Println("Expected: 1	Output: ", maxOperations(nums2, k2))

}
