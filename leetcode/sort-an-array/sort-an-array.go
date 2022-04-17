package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/sort-an-array/

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func main() {
	// Example 1
	var nums1 []int = [5,2,3,1]
	fmt.Println("Expected: [1,2,3,5]	Output: ", sortArray(nums1))

	// Example 2
	var nums2 []int = [5,1,1,2,0,0]
	fmt.Println("Expected: [0,0,1,1,2,5]	Output: ", sortArray(nums2))

}
