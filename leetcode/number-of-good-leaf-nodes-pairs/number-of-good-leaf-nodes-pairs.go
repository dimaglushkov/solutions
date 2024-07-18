package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func countPairs(root *TreeNode, distance int) int {
	ans := 0

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		if node.Left == nil && node.Right == nil {
			return []int{1}
		}

		leftDistances := dfs(node.Left)
		rightDistances := dfs(node.Right)

		for _, l := range leftDistances {
			for _, r := range rightDistances {
				if l+r <= distance {
					ans++
				}
			}
		}

		distances := append(leftDistances, rightDistances...)
		for i := range distances {
			distances[i]++
		}

		return distances
	}

	dfs(root)
	return ans
}

func main() {
	testCases := []struct {
		root     *TreeNode
		distance int
		want     int
	}{
		{
			root:     NewTreeNode([]int{1, 2, 3, 4, 5, 6, 7}),
			distance: 3,
			want:     2,
		},
		{
			root:     NewTreeNode([]int{80, 62, 99, 36, 45, 39, 76, 81, 44, 23, 58, 8, 14, 1, 94, 100, 10, 8, 30, 75, 7, 33, 80, 44, 2, 67, 78, 64, 30, 98, 100, 24, 48, 42, 31, 91, 37, 81, 45, 30, 77, 28}),
			distance: 8,
			want:     122,
		},
		{
			root:     NewTreeNode([]int{78, 15, 81, 73, 98, 36, -1, 30, -1, 63, 32}),
			distance: 6,
			want:     6,
		},
		{
			root:     NewTreeNode([]int{15, 66, 55, 97, 60, 12, 56, -1, 54, -1, 49, -1, 9, -1, -1, -1, -1, -1, -1, -1, -1, -1, 90}),
			distance: 5,
			want:     3,
		},
		{
			root:     NewTreeNode([]int{7, 1, 4, 6, -1, 5, 3, -1, -1, -1, -1, -1, -1, -1, 2}),
			distance: 3,
			want:     1,
		},
		{
			root:     NewTreeNode([]int{1, 2, 3, -1, 4}),
			distance: 3,
			want:     1,
		},
		{
			root:     NewTreeNode([]int{1, 2, 3}),
			distance: 2,
			want:     1,
		},
		{
			root:     NewTreeNode([]int{100}),
			distance: 1,
			want:     0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countPairs(tc.root, tc.distance)
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
