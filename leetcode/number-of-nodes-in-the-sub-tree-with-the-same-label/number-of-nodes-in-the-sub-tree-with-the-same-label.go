package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/

type counter [26]int

func countSubTrees(n int, edges [][]int, labels string) []int {
	tree := make([][]int, n)
	res := make([]int, n)

	// build tree
	for _, e := range edges {
		if tree[e[0]] == nil {
			tree[e[0]] = make([]int, 0)
		}
		if tree[e[1]] == nil {
			tree[e[1]] = make([]int, 0)
		}
		tree[e[0]] = append(tree[e[0]], e[1])
		tree[e[1]] = append(tree[e[1]], e[0])
	}

	var dfs func(prev, node int) counter
	dfs = func(prev, node int) counter {
		var cnt counter
		cnt[int(labels[node]-'a')]++

		if len(tree[node]) == 1 && tree[node][0] == prev {
			res[node] = 1
			return cnt
		}

		for _, ch := range tree[node] {
			if prev != ch {
				for i, c := range dfs(node, ch) {
					cnt[i] += c
				}
			}
		}

		res[node] = cnt[labels[node]-'a']
		return cnt
	}

	dfs(-1, 0)
	return res
}

func main() {
	testCases := []struct {
		n      int
		edges  [][]int
		labels string
		want   []int
	}{
		{
			n:      4,
			edges:  [][]int{{0, 1}, {1, 2}, {0, 3}},
			labels: "bbbb",
			want:   []int{4, 2, 1, 1},
		},
		{
			n:      4,
			edges:  [][]int{{0, 2}, {0, 3}, {1, 2}},
			labels: "aeed",
			want:   []int{1, 1, 2, 1},
		},
		{
			n:      7,
			edges:  [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
			labels: "abaedcd",
			want:   []int{2, 1, 1, 1, 1, 1, 1},
		},
		{
			n:      5,
			edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}},
			labels: "aabab",
			want:   []int{3, 2, 1, 1, 1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countSubTrees(tc.n, tc.edges, tc.labels)
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
