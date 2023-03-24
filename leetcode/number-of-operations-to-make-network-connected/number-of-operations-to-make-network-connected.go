package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/number-of-operations-to-make-network-connected/

func makeConnected(n int, connections [][]int) int {
	ans := n - 1
	if len(connections) < ans {
		return -1
	}

	parent := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		if f1, f2 := find(x), find(y); f1 != f2 {
			parent[find(x)] = find(y)
			ans--
		}
	}

	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for _, c := range connections {
		union(c[0], c[1])
	}
	return ans
}

func main() {
	testCases := []struct {
		n           int
		connections [][]int
		want        int
	}{
		{
			n:           4,
			connections: [][]int{{0, 1}, {0, 2}, {1, 2}},
			want:        1,
		},
		{
			n:           6,
			connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}},
			want:        2,
		},
		{
			n:           6,
			connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}},
			want:        -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := makeConnected(tc.n, tc.connections)
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
