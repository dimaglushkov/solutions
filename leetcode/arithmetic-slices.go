package main

import "fmt"

// source: https://leetcode.com/problems/arithmetic-slices/

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	var count, diff, id int
	var prevDiff = nums[1] - nums[0]

	for i := 1; i < len(nums)-1; i++ {
		if diff = nums[i+1] - nums[i]; diff == prevDiff {
			id++
		} else {
			prevDiff = diff
			id = 0
		}
		count += id
	}
	return count
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3, 4}
	fmt.Println("Expected: 3	Output: ", numberOfArithmeticSlices(nums1))

	// Example 2
	var nums2 = []int{1}
	fmt.Println("Expected: 0	Output: ", numberOfArithmeticSlices(nums2))

}
