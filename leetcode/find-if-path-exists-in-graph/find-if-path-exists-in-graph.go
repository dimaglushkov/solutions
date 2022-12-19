package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-if-path-exists-in-graph/

func validPath(n int, edges [][]int, src int, dst int) bool {
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
	return find(src) == find(dst)
}

func main() {
	testCases := []struct {
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{
			n:           3,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 0}},
			source:      0,
			destination: 2,
			want:        true,
		},
		{
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			source:      0,
			destination: 5,
			want:        false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := validPath(tc.n, tc.edges, tc.source, tc.destination)
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
