package main

import (
	"fmt"
	"math"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Left != nil && (!(node.Left.Val > min) || node.Left.Val >= node.Val) {
			return false
		}
		if node.Right != nil && (!(node.Right.Val < max) || node.Right.Val <= node.Val) {
			return false
		}
		return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}

func main() {
	t2 := NewTreeNode([]int{5, 1, 4, -1, -1, 3, 6})
	fmt.Println("Expected: false	Output: ", isValidBST(t2))

	t1 := NewTreeNode([]int{5, 4, 6, -1, -1, 3, 7})
	fmt.Println("Expected: false	Output: ", isValidBST(t1))

	t0 := NewTreeNode([]int{2, 1, 3})
	fmt.Println("Expected: true	Output: ", isValidBST(t0))

}
