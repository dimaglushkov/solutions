package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func flatten(root *TreeNode) {
	if root == nil || root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil {
		flatten(root.Left)

		tmp := root.Right
		root.Right = root.Left
		root.Left = nil

		t := root.Right
		for t.Right != nil {
			t = t.Right
		}

		t.Right = tmp
	}

	flatten(root.Right)
}
