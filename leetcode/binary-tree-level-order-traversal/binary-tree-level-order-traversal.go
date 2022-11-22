package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/binary-tree-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type node struct {
	v *TreeNode
	l int
}

func newNode(treeNode *TreeNode, level int) node {
	return node{treeNode, level}
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	q := make([]node, 0)
	push := func(treeNode *TreeNode, l int) {
		q = append(q, newNode(treeNode, l))
	}
	pop := func() (*TreeNode, int) {
		x := q[0]
		q = q[1:]
		return x.v, x.l
	}

	push(root, 0)
	for len(q) > 0 {
		nv, nl := pop()
		if nv == nil {
			continue
		}
		if nl == len(res) {
			res = append(res, make([]int, 0))
		}
		res[nl] = append(res[nl], nv.Val)
		push(nv.Left, nl+1)
		push(nv.Right, nl+1)
	}

	return res
}

func main() {
	// Example 1
	var root1 *TreeNode = NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println("Expected: [[3],[9,20],[15,7]]	Output: ", levelOrder(root1))

	// Example 2
	var root2 *TreeNode = NewTreeNode([]int{1})
	fmt.Println("Expected: [[1]]	Output: ", levelOrder(root2))

	// Example 3
	var root3 *TreeNode = NewTreeNode([]int{})
	fmt.Println("Expected: []	Output: ", levelOrder(root3))

}
