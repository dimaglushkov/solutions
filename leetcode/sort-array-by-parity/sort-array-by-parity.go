package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/sort-array-by-parity/

func sortArrayByParity(nums []int) []int {
	index := 0
	for i, v := range nums {
		if v%2 == 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	return nums
}

func sortArrayByParity_(nums []int) []int {
	var odd, even = make([]int, 0, len(nums)), make([]int, 0, len(nums))
	for i := range nums {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	return append(even, odd...)
}

func sortArrayByParity__(nums []int) []int {
	var res = make([]int, 0, len(nums))
	for i := range nums {
		if nums[i]%2 == 0 {
			res = append([]int{nums[i]}, res...)
		} else {
			res = append(res, nums[i])
		}
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{3, 1, 2, 4}
	fmt.Println("Expected: [2,4,3,1]	Output: ", sortArrayByParity(nums1))

	// Example 2
	var nums2 = []int{0}
	fmt.Println("Expected: [0]	Output: ", sortArrayByParity(nums2))

}
