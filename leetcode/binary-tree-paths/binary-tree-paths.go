package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
	"strconv"
	"strings"
)
// source: https://leetcode.com/problems/binary-tree-paths/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}


func buildResString(path []int) string {
	values := make([]string, 0, len(path))
	for _, x := range path {
		values = append(values, itoa(x))
	}
	return strings.Join(values, "->")
}

func binaryTreePaths(root *TreeNode) []string {
    res := make([]string, 0)
	var dfs func(n *TreeNode, path []int)
	dfs = func(n *TreeNode, path []int)  {
		if n == nil {
			return
		}
		path = append(path, n.Val)
		if n.Left == nil && n.Right == nil {
			res = append(res, buildResString(path))
		}
		dfs(n.Left, path)
		dfs(n.Right, path)
	}

	dfs(root, []int{})
	return res
}

func main() {
    testCases := []struct {
		root *TreeNode
		want []string
    }{
		{
			root:  NewTreeNode([]int{1,2,3,-1,5}),
			want: []string {"1->2->5","1->3"},
		},
		{
			root:  NewTreeNode([]int{1}),
			want: []string {"1"},
		},
	}


    successes := 0
    for _, tc := range testCases {
        x := binaryTreePaths(tc.root)
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
        fmt.Printf("===\nFAIL: %d tests failed\n", l - successes)
    }

}
