package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func smallestFromLeaf(root *TreeNode) string {
	ans := ""

	var dfs func(node *TreeNode, text string)
	dfs = func(node *TreeNode, text string) {
		if node == nil {
			return
		}

		text = string(rune(node.Val+97)) + text
		if node.Right == nil && node.Left == nil {
			if ans == "" || ans > text {
				ans = text
			}

			return
		}

		dfs(node.Left, text)
		dfs(node.Right, text)
	}

	dfs(root, "")

	return ans
}

func main() {
	testCases := []struct {
		root *TreeNode
		want string
	}{
		{
			root: {0, 1, 2, 3, 4, 3, 4},
			want: "dba",
		},
		{
			root: {25, 1, 3, 1, 3, 0, 2},
			want: "adz",
		},
		{
			root: {2, 2, 1, null, 1, 0, null, 0},
			want: "abc",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := smallestFromLeaf(tc.root)
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
