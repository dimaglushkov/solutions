package main

import "fmt"

// source: https://leetcode.com/problems/binary-search/

// internationally not using sort.Search()
func search(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		h := (i + j) >> 1
		if target > nums[h] {
			i = h + 1
		} else {
			j = h
		}
	}
	if i == len(nums) || nums[i] != target {
		return -1
	}
	return i
}

func main() {
	// Example 1
	var nums1 = []int{-1, 0, 3, 5, 9, 12}
	var target1 int = 9
	fmt.Println("Expected: 4	Output: ", search(nums1, target1))

	// Example 2
	var nums2 = []int{-1, 0, 3, 5, 9, 12}
	var target2 int = 2
	fmt.Println("Expected: -1	Output: ", search(nums2, target2))

}
