package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
	"sort"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getMinimumDifference(root *TreeNode) int {
	ans := 1<<31 - 1
	var values []int
	var trav func(node *TreeNode)
	trav = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		trav(node.Left)
		trav(node.Right)
	}

	trav(root)

	sort.Ints(values)
	for i := 0; i < len(values)-1; i++ {
		ans = min(ans, values[i+1]-values[i])
	}

	return ans
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{4, 2, 6, 1, 3}),
			want: 1,
		},
		{
			root: NewTreeNode([]int{1, 0, 48, -1, -1, 12, 49}),
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := getMinimumDifference(tc.root)
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
