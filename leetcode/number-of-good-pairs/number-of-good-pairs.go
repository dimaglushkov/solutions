package main

import "fmt"

// source: https://leetcode.com/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) (res int) {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3, 1, 1, 3}
	fmt.Println("Expected: 4	Output: ", numIdenticalPairs(nums1))

	// Example 2
	var nums2 = []int{1, 1, 1, 1}
	fmt.Println("Expected: 6	Output: ", numIdenticalPairs(nums2))

	// Example 3
	var nums3 = []int{1, 2, 3}
	fmt.Println("Expected: 0	Output: ", numIdenticalPairs(nums3))

}
