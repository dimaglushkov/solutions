package main

import (
	"fmt"
)

type State struct {
	mask, node, dist int
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	allVisited := (1 << n) - 1
	queue := []State{}
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		queue = append(queue, State{1 << i, i, 0})
		visited[(1<<i)*16+i] = true
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.mask == allVisited {
			return cur.dist
		}

		for _, neighbor := range graph[cur.node] {
			newMask := cur.mask | (1 << neighbor)
			hashValue := newMask*16 + neighbor

			if !visited[hashValue] {
				visited[hashValue] = true
				queue = append(queue, State{newMask, neighbor, cur.dist + 1})
			}
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		graph [][]int
		want  int
	}{
		{
			graph: [][]int{{1, 2, 3}, {0}, {0}, {0}},
			want:  4,
		},
		{
			graph: [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}},
			want:  4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := shortestPathLength(tc.graph)
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
