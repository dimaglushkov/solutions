package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/shortest-path-with-alternating-colors/

type node struct {
	i, d, pec int
}

type queue []node

func (q *queue) push(x node) {
	*q = append(*q, x)
}
func (q *queue) pop() node {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	const redEdge, blueEdge = 0, 1
	edges := make([][][2]int, n)
	for i := range edges {
		edges[i] = make([][2]int, n)
	}
	for _, e := range redEdges {
		edges[e[0]] = append(edges[e[0]], [2]int{e[1], redEdge})
	}
	for _, e := range blueEdges {
		edges[e[0]] = append(edges[e[0]], [2]int{e[1], blueEdge})
	}
	visited := make([][2]bool, n)

	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = -1
	}

	q := new(queue)
	q.push(node{0, 0, -1})
	for len(*q) > 0 {
		x := q.pop()
		for _, e := range edges[x.i] {
			if !visited[e[0]][e[1]] && e[1] != x.pec {
				visited[e[0]][e[1]] = true
				q.push(node{e[0], x.d + 1, e[1]})
				if ans[e[0]] == -1 {
					ans[e[0]] = x.d + 1
				}
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
		want      []int
	}{
		{
			n:         5,
			redEdges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			blueEdges: [][]int{{1, 2}, {2, 3}, {3, 1}},
			want:      []int{0, 1, 2, 3, 7},
		},
		{
			n:         3,
			redEdges:  [][]int{{0, 1}, {0, 2}},
			blueEdges: [][]int{{1, 0}},
			want:      []int{0, 1, 1},
		},
		{
			n:         3,
			redEdges:  [][]int{{0, 1}, {1, 2}},
			blueEdges: [][]int{},
			want:      []int{0, 1, -1},
		},
		{
			n:         3,
			redEdges:  [][]int{{0, 1}},
			blueEdges: [][]int{{2, 1}},
			want:      []int{0, 1, -1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := shortestAlternatingPaths(tc.n, tc.redEdges, tc.blueEdges)
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
