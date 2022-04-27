package tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(data []int) *TreeNode {
	return new(TreeNode).insertNode(data, 0)
}

func (r *TreeNode) insertNode(data []int, i int) *TreeNode {
	var root *TreeNode
	if i >= len(data) || data[i] == -1 {
		return root
	}
	root = &TreeNode{Val: data[i]}
	root.Left = root.insertNode(data, 2*i+1)
	root.Right = root.insertNode(data, 2*i+2)
	return root
}

func (r *TreeNode) getNode(data []int, i int) int {
	if r == nil {
		return i
	}
	data[i] = r.Val
	i++
	if r.Left != nil {
		i = r.Left.getNode(data, i)
	}
	if r.Right != nil {
		i = r.Right.getNode(data, i)
	}
	return i
}

func (r *TreeNode) Height(h int) (maxH int) {
	if r == nil {
		return h
	}
	var lHeight, rHeight = h, h
	if r.Left != nil {
		lHeight = r.Left.Height(h + 1)
	}
	if r.Right != nil {
		rHeight = r.Right.Height(h + 1)
	}
	if lHeight > rHeight {
		return lHeight
	}
	return rHeight

}

func (r *TreeNode) Slice() []int {
	var data = make([]int, 1<<(r.Height(1)))
	r.getNode(data, 0)
	return data
}

func (r *TreeNode) String() string {
	data := r.Slice()
	str := make([]string, len(data))
	for i, val := range data {
		str[i] = strconv.FormatInt(int64(val), 10)
	}
	return strings.Join(str, ", ")
}
