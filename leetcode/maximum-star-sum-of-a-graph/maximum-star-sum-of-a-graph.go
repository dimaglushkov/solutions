package main

import (
	"fmt"
	"math"
	"sort"
)

// source: https://leetcode.com/problems/maximum-star-sum-of-a-graph/

func max(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] > x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func sum(x ...int) int {
	res := 0
	for _, i := range x {
		res += i
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxStarSum(vals []int, edges [][]int, k int) int {
	if k == 0 {
		return max(vals...)
	}

	graph := make(map[int]map[int]int)
	for _, e := range edges {
		if graph[e[0]] == nil {
			graph[e[0]] = make(map[int]int)
		}
		if vals[e[1]] > 0 {
			graph[e[0]][e[1]] = vals[e[1]]
		}

		if graph[e[1]] == nil {
			graph[e[1]] = make(map[int]int)
		}
		if vals[e[0]] > 0 {
			graph[e[1]][e[0]] = vals[e[0]]
		}
	}

	res := math.MinInt

	for i := range vals {
		es := make([]int, 0)
		for _, w := range graph[i] {
			es = append(es, w)
		}
		sort.Slice(es, func(i, j int) bool {
			return es[i] > es[j]
		})
		size := min(k, len(es))
		res = max(res, vals[i]+sum(es[:size]...))
	}

	return res
}

func main() {
	testCases := []struct {
		vals  []int
		edges [][]int
		k     int
		want  int
	}{
		{
			vals:  []int{1, 2, 3, 4, 10, -10, -20},
			edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6}},
			k:     2,
			want:  16,
		},
		{
			vals:  []int{-5},
			edges: [][]int{},
			k:     0,
			want:  -5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxStarSum(tc.vals, tc.edges, tc.k)
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
