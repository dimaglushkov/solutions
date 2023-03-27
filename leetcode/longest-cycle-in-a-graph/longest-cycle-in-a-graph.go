package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/longest-cycle-in-a-graph/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestCycle(edges []int) int {
	ans := -1
	n := len(edges)
	visited := make([]bool, n)
	var dfs func(i, d int, path map[int]int)
	dfs = func(i, d int, path map[int]int) {
		if v, ok := path[i]; ok {
			ans = max(ans, d-v)
			return
		}
		if edges[i] == -1 || visited[i] {
			return
		}

		path[i] = d
		visited[i] = true
		dfs(edges[i], d+1, path)
		delete(path, i)
	}
	for i := range edges {
		if !visited[i] {
			dfs(i, 0, make(map[int]int))
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		edges []int
		want  int
	}{
		{
			edges: []int{3, 3, 4, 2, 3},
			want:  3,
		},
		{
			edges: []int{2, -1, 3, 1},
			want:  -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := longestCycle(tc.edges)
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
