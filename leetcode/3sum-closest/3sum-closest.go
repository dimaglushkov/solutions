package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/3sum-closest/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	var res = 1001

	sort.Ints(nums)
	for i := 0; i < length; i++ {
		for l, r := i+1, length-1; l < r; {
			n := nums[i] + nums[l] + nums[r]
			if abs(target-n) < abs(target-res) {
				res = n
			}
			if n-target > 0 {
				r--
			} else {
				l++
			}
		}
	}
	if res == 1001 {
		return 0
	}
	return res
}

func main() {
	// Example 0
	var nums0 = []int{0, 0, 0}
	var target0 int = 10000
	fmt.Println("Expected: 0	Output: ", threeSumClosest(nums0, target0))

	// Example 3
	var nums3 = []int{0, 2, 1, -3}
	var target3 int = 1
	fmt.Println("Expected: 0	Output: ", threeSumClosest(nums3, target3))

	// Example 1
	var nums1 = []int{-1, 2, 1, -4}
	var target1 int = 1
	fmt.Println("Expected: 2	Output: ", threeSumClosest(nums1, target1))

	// Example 2
	var nums2 = []int{0, 0, 0}
	var target2 int = 1
	fmt.Println("Expected: 0	Output: ", threeSumClosest(nums2, target2))

}
