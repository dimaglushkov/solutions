package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0

	var util func(node *TreeNode)
	util = func(node *TreeNode) {
		if node == nil {
			return
		}

		util(node.Right)

		sum += node.Val
		node.Val = sum

		util(node.Left)

	}

	util(root)

	return root
}

func main() {
	testCases := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{
			root: NewTreeNode([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}),
			want: NewTreeNode([]int{30, 36, 21, 36, 35, 26, 15, -1, -1, -1, 33, -1, -1, -1, 8}),
		},
		{
			root: NewTreeNode([]int{0, -1, 1}),
			want: NewTreeNode([]int{1, -1, 1}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := bstToGst(tc.root)
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
