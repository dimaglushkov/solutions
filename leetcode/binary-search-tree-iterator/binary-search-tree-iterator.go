package main

import (
	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/binary-search-tree-iterator/

// Instead of going with straight-forward solution
// convert given tree to array and implement Next() and HasNext()
// using index and array representation

type BSTIterator struct {
	values []int
	ind    int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		values: slicify(root),
		ind:    0,
	}
}

func (iter *BSTIterator) Next() int {
	iter.ind++
	return iter.values[iter.ind-1]
}

func (iter *BSTIterator) HasNext() bool {
	return iter.ind < len(iter.values)
}

func slicify(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	data := make([]int, 0)

	data = append(data, slicify(node.Left)...)
	data = append(data, node.Val)
	data = append(data, slicify(node.Right)...)
	return data
}

func main() {
	tr1 := NewTreeNode([]int{7, 3, 15, -1, -1, 9, 20})
	it1 := Constructor(tr1)
	print(it1.values)
}
