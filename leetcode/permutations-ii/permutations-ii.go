package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	var res = make([][]int, 0)
	var used = make([]bool, len(nums))
	var util func(nums, cur []int)
	util = func(nums, cur []int) {
		if len(cur) == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := range nums {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			used[i] = true
			cur = append(cur, nums[i])
			util(nums, cur)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}

	sort.Ints(nums)
	util(nums, []int{})

	return res
}

func main() {
	// Example 3
	var nums3 = []int{1, 1, 2, 2}
	fmt.Println("Expected: [[1,1,2,2],[1,2,1,2],[1,2,2,1],[2,1,1,2],[2,1,2,1],[2,2,1,1]]	Output: ", permuteUnique(nums3))

	// Example 1
	var nums1 = []int{1, 1, 2}
	fmt.Println("Expected: [[1,1,2], [1,2,1], [2,1,1]]	Output: ", permuteUnique(nums1))

	// Example 2
	var nums2 = []int{1, 2, 3}
	fmt.Println("Expected: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]	Output: ", permuteUnique(nums2))

}
