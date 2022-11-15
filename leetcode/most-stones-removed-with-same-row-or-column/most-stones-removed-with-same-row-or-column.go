package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/

func removeStones(stones [][]int) int {
	n := len(stones)
	parent := make([]int, n)
	xs, ys := make(map[int][]int), make(map[int][]int)
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

	for i, s := range stones {
		parent[i] = i
		x, y := s[0], s[1]
		if _, ok := xs[x]; !ok {
			xs[x] = make([]int, 0)
		}
		if _, ok := ys[y]; !ok {
			ys[y] = make([]int, 0)
		}
		xs[x] = append(xs[x], i)
		ys[y] = append(ys[y], i)
	}

	for _, ps := range xs {
		s := ps[0]
		for i := 1; i < len(ps); i++ {
			union(s, ps[i])
		}
	}
	for _, ps := range ys {
		s := ps[0]
		for i := 1; i < len(ps); i++ {
			union(s, ps[i])
		}
	}

	res := make(map[int]bool, 0)
	for i := 0; i < n; i++ {
		res[find(i)] = true
	}

	return n - len(res)
}

func main() {
	testCases := []struct {
		stones [][]int
		want   int
	}{
		{
			stones: [][]int{{1, 2}, {1, 3}, {3, 3}, {3, 1}, {2, 1}, {1, 0}},
			want:   5,
		},
		{
			stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
			want:   5,
		},
		{
			stones: [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
			want:   3,
		},
		{
			stones: [][]int{{0, 0}},
			want:   0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := removeStones(tc.stones)
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
		fmt.Printf("===\nFAIL: only %d of %d tests ended successfully\n", successes, l)
	}

}
