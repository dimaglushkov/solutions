package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/

func countPairs(n int, edges [][]int) int64 {
	parent := make([]int, n)
	var find func(int) int
	var union func(int, int)
	union = func(x, y int) {
		parent[find(x)] = find(y)
	}
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	for i := range parent {
		parent[i] = i
	}
	for _, e := range edges {
		union(e[0], e[1])
	}

	groups := make(map[int]int)
	for i := range parent {
		groups[find(i)]++
	}
	if len(groups) == 1 {
		return 0
	}
	var ans int64
	for _, cnt := range groups {
		ans += int64(cnt * (n - cnt))
		n -= cnt
	}
	return ans
}

func main() {
	testCases := []struct {
		n     int
		edges [][]int
		want  int64
	}{

		{
			n:     7,
			edges: [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}},
			want:  14,
		},
		{
			n:     3,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}},
			want:  0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countPairs(tc.n, tc.edges)
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
