package main

import "fmt"

// source: https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, i := range nums {
		if m[i] {
			return true
		}
		m[i] = true
	}
	return false
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3, 1}
	fmt.Println("Expected: true	Output: ", containsDuplicate(nums1))

	// Example 2
	var nums2 = []int{1, 2, 3, 4}
	fmt.Println("Expected: false	Output: ", containsDuplicate(nums2))

	// Example 3
	var nums3 = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println("Expected: true	Output: ", containsDuplicate(nums3))

}
