package tree

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
	if i >= len(data) {
		return root
	}
	root = &TreeNode{Val: data[i]}
	root.Left = root.insertNode(data, 2*i+1)
	root.Right = root.insertNode(data, 2*i+2)
	return root
}
