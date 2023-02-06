package main

import "fmt"

// source: https://leetcode.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}
	return res
}

// unnecessary complicated old version of the same solution
func shuffle_(nums []int, n int) []int {
	res := make([]int, n*2)
	i, j := 0, n
	for k := 0; k < n*2; k += 2 {
		res[k] = nums[i]
		res[k+1] = nums[j]
		i++
		j++
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{2, 5, 1, 3, 4, 7}
	var n1 int = 3
	fmt.Println("Expected: [2,3,5,4,1,7]	Output: ", shuffle(nums1, n1))

	// Example 2
	var nums2 = []int{1, 2, 3, 4, 4, 3, 2, 1}
	var n2 int = 4
	fmt.Println("Expected: [1,4,2,3,3,2,4,1]	Output: ", shuffle(nums2, n2))

	// Example 3
	var nums3 = []int{1, 1, 2, 2}
	var n3 int = 2
	fmt.Println("Expected: [1,2,1,2]	Output: ", shuffle(nums3, n3))

}
