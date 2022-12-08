package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minScore(n int, roads [][]int) int {
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

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for _, r := range roads {
		union(r[0]-1, r[1]-1)
	}

	t := find(0)
	res := math.MaxInt
	for _, r := range roads {
		if find(r[0]-1) != t || find(r[1]-1) != t {
			continue
		}
		res = min(res, r[2])
	}

	return res
}
func main() {
	testCases := []struct {
		n     int
		roads [][]int
		want  int
	}{
		{
			n:     4,
			roads: [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}},
			want:  5,
		},
		{
			n:     4,
			roads: [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}},
			want:  2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minScore(tc.n, tc.roads)
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
