package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/3sum/

// I guess this solution is pretty simple and no description is needed
// We could have used the same approach we used for two-sum problem, but
// in this case we need to also dedup result, and it may take too much time to check
// whether triplet is already presented in the result or not
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	length := len(nums)
	res := make([][]int, 0, length/2)

	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		for l, r := i+1, length-1; l < r; {
			n := nums[i] + nums[l] + nums[r]
			if n == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for dup := nums[l]; l < r && nums[l] == dup; {
					l++
				}
			} else if n > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func main() {
	// Example 1
	var nums1 = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("Expected: [[-1,-1,2],[-1,0,1]]	Output: ", threeSum(nums1))

	// Example 2
	var nums2 = []int{-2, -2, 4, 4}
	fmt.Println("Expected: [[-2,-2,4],[-2,-2,4]]	Output: ", threeSum(nums2))

	// Example 3
	var nums3 = []int{0, 1, 2}
	fmt.Println("Expected: []	Output: ", threeSum(nums3))

	// Example 4
	var nums4 = []int{}
	fmt.Println("Expected: []	Output: ", threeSum(nums4))

	// Example 5
	var nums5 = []int{0}
	fmt.Println("Expected: []	Output: ", threeSum(nums5))

	// Example 6
	var nums6 = []int{0, 1, -1}
	fmt.Println("Expected: [-1, 0, 1]	Output: ", threeSum(nums6))

	// Example 7
	var nums7 = []int{0, 1, -1, -1}
	fmt.Println("Expected: [-1, 0, 1]	Output: ", threeSum(nums7))

}
