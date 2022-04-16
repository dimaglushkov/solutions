package main

import "fmt"

// source: https://leetcode.com/problems/build-array-from-permutation/

func buildArray(nums []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = nums[v]
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{0, 2, 1, 5, 3, 4}
	fmt.Println("Expected: [0,1,2,4,5,3]	Output: ", buildArray(nums1))

	// Example 2
	var nums2 = []int{5, 0, 1, 2, 3, 4}
	fmt.Println("Expected: [4,5,0,1,2,3]	Output: ", buildArray(nums2))

}
