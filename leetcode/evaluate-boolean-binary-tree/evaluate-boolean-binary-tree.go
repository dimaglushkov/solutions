package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

const (
	TRUE  = 1
	FALSE = 0
	OR    = 2
	AND   = 3
)

func evaluateTree(root *TreeNode) bool {
	switch root.Val {
	case TRUE:
		return true
	case FALSE:
		return false
	case OR:
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	case AND:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}

	return false
}

func main() {
	testCases := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: NewTreeNode([]int{2, 1, 3, -1, -1, 0, 1}),
			want: true,
		},
		{
			root: NewTreeNode([]int{0}),
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := evaluateTree(tc.root)
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
