package main

import "fmt"

// source: https://leetcode.com/problems/missing-number/

func missingNumber(nums []int) int {
	sum := int(float64(len(nums)+1.) / 2 * float64(len(nums)))
	for _, i := range nums {
		sum -= i
	}
	return sum
}

func main() {
	// Example 1
	var nums1 = []int{3, 0, 1}
	fmt.Println("Expected: 2	Output: ", missingNumber(nums1))

	// Example 2
	var nums2 = []int{0, 1}
	fmt.Println("Expected: 2	Output: ", missingNumber(nums2))

	// Example 3
	var nums3 = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println("Expected: 8	Output: ", missingNumber(nums3))

}
