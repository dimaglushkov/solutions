package main

import "fmt"

// source: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

func main() {
	// Example 1
	var nums0 = []int{}
	fmt.Println("Expected: 2, nums = [1,2,_]	Output: ", removeDuplicates(nums0))

	// Example 1
	var nums1 = []int{1, 1, 2}
	fmt.Println("Expected: 2, nums = [1,2,_]	Output: ", removeDuplicates(nums1))

	// Example 2
	var nums2 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Expected: 5, nums = [0,1,2,3,4,_,_,_,_,_]	Output: ", removeDuplicates(nums2))

}
