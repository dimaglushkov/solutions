package main

import "fmt"

// source: https://leetcode.com/problems/concatenation-of-array/

func getConcatenation(nums []int) []int {
	res := make([]int, len(nums)*2)
	var i, j int
	for i = 0; i < len(nums); i++ {
		res[i] = nums[i]
	}
	for j = 0; j < len(nums); i, j = i+1, j+1 {
		res[i] = nums[j]
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 1}
	fmt.Println("Expected: [1,2,1,1,2,1]	Output: ", getConcatenation(nums1))

	// Example 2
	var nums2 = []int{1, 3, 2, 1}
	fmt.Println("Expected: [1,3,2,1,1,3,2,1]	Output: ", getConcatenation(nums2))

}
