package main

import (
	"fmt"
	"math"

	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/validate-binary-search-tree/

func util(r *TreeNode, p *int) bool {
	if r == nil {
		return true
	}
	if !util(r.Left, p) || r.Val <= *p {
		return false
	}
	*p = r.Val
	return util(r.Right, p)
}

func isValidBST(root *TreeNode) bool {
	var p = math.MinInt
	return util(root, &p)
}

func main() {
	t1 := NewTreeNode([]int{5, 4, 6, -1, -1, 3, 7})
	fmt.Println("Expected: false	Output: ", isValidBST(t1))

	t0 := NewTreeNode([]int{2, 1, 3})
	fmt.Println("Expected: true	Output: ", isValidBST(t0))

}
