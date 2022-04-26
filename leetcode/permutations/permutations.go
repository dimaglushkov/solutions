package main

import "fmt"

// source: https://leetcode.com/problems/permutations/

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	permuteUtil(nums, []int{})
	return res
}

func permuteUtil(nums, cur []int) {
	if len(nums) == 0 {
		res = append(res, cur)
		return
	}
	for i := 0; i < len(nums); i++ {
		var newNums = make([]int, len(nums))
		copy(newNums, nums)
		permuteUtil(append(newNums[:i], newNums[i+1:]...), append(cur, nums[i]))
	}
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3}
	fmt.Println("Expected: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]	Output: ", permute(nums1))

	// Example 2
	var nums2 = []int{0, 1}
	fmt.Println("Expected: [[0,1],[1,0]]	Output: ", permute(nums2))

	// Example 3
	var nums3 = []int{1}
	fmt.Println("Expected: [[1]]	Output: ", permute(nums3))

}
