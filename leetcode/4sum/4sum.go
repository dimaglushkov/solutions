package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/4sum/

func threeSum(nums []int, prev, target int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	length := len(nums)
	res := make([][]int, 0, length/2)

	for i := 0; i < length; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		for l, r := i+1, length-1; l < r; {
			n := nums[i] + nums[l] + nums[r]
			if n == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for dup := nums[l]; l < r && nums[l] == dup; {
					l++
				}
			} else if n > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func fourSum(nums []int, target int) [][]int {
	var threes [][]int
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		threes = threeSum(nums[i+1:], nums[i], target-nums[i])
		for _, t := range threes {
			temp := append([]int{nums[i]}, t...)
			res = append(res, temp)
		}

	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{1, 0, -1, 0, -2, 2}
	var target1 int = 0
	fmt.Println("Expected: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]	Output: ", fourSum(nums1, target1))

	// Example 2
	var nums2 = []int{2, 2, 2, 2, 2}
	var target2 int = 8
	fmt.Println("Expected: [[2,2,2,2]]	Output: ", fourSum(nums2, target2))

}
