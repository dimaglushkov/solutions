package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/array-partition-i/

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{1, 4, 3, 2}
	fmt.Println("Expected: 4	Output: ", arrayPairSum(nums1))

	// Example 2
	var nums2 = []int{6, 2, 6, 5, 1, 2}
	fmt.Println("Expected: 9	Output: ", arrayPairSum(nums2))

}
