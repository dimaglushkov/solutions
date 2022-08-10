package main

import "fmt"

// source: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	i := n / 2
	root := &TreeNode{Val: nums[i]}
	if i != 0 {
		root.Left = sortedArrayToBST(nums[:i])
	}
	if i != n-1 {
		root.Right = sortedArrayToBST(nums[i+1:])
	}

	return root
}

func main() {
	// Example 1
	var nums1 = []int{-10, -3, 0, 5, 9}
	fmt.Println("Expected: [0,-3,9,-10,null,5]	Output: ", sortedArrayToBST(nums1))

	// Example 2
	var nums2 = []int{1, 3}
	fmt.Println("Expected: [3,1]	Output: ", sortedArrayToBST(nums2))

}
