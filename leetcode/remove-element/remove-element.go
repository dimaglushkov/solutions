package main

import "fmt"

// source: https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	var i, j int
	for j = 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	// Example 2
	var nums2 = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var val2 int = 2
	k2 := removeElement(nums2, val2)
	fmt.Println("Expected: 5, nums = [0,1,4,0,3,_,_,_]	Output: ", k2, nums2)

	// Example 1
	var nums1 = []int{3, 2, 2, 3}
	var val1 int = 3
	k1 := removeElement(nums1, val1)
	fmt.Println("Expected: 2, nums = [2,2,_,_]	Output: ", k1, nums1)

}
