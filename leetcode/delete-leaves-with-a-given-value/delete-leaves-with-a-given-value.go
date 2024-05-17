package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}

	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

func main() {
	testCases := []struct {
		root   *TreeNode
		target int
		want   *TreeNode
	}{
		{
			root:   NewTreeNode([]int{1, 2, 3, 2, -1, 2, 4}),
			target: 2,
			want:   NewTreeNode([]int{1, -1, 3, -1, 4}),
		},
		{
			root:   NewTreeNode([]int{1, 3, 3, 3, 2}),
			target: 3,
			want:   NewTreeNode([]int{1, 3, -1, -1, 2}),
		},
		{
			root:   NewTreeNode([]int{1, 2, -1, 2, -1, 2}),
			target: 2,
			want:   NewTreeNode([]int{1}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := removeLeafNodes(tc.root, tc.target)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
