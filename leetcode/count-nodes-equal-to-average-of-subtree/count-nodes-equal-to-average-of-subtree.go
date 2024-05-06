package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
	"math"
)

func averageOfSubtree(root *TreeNode) int {
	var util func(node *TreeNode) (int, int, int)
	util = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return 0, 0, 0
		}
		suml, cntl, ansl := util(node.Left)
		sumr, cntr, ansr := util(node.Right)

		sum := node.Val + sumr + suml
		cnt := cntr + cntl + 1

		ans := ansr + ansl

		if int(math.Floor(float64(sum)/float64(cnt))) == node.Val {
			ans++
		}

		return sum, cnt, ans
	}

	_, _, ans := util(root)
	return ans
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{4, 8, 5, 0, 1, -1, 6}),
			want: 5,
		},
		{
			root: NewTreeNode([]int{1}),
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := averageOfSubtree(tc.root)
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
