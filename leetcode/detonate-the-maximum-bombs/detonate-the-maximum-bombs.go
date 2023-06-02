package main

import (
	"fmt"
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type queue []int

func (q *queue) push(x int) {
	*q = append(*q, x)
}
func (q *queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func bfs(start int, graph [][]bool) int {
	var q queue
	var cnt int
	visited := make([]bool, len(graph))
	q.push(start)
	for len(q) > 0 {
		x := q.pop()
		if visited[x] {
			continue
		}
		visited[x] = true
		cnt++
		for i, connected := range graph[x] {
			if connected {
				q.push(i)
			}
		}
	}
	return cnt
}

func isConnected(c, p []int) bool {
	d := math.Sqrt(float64((c[0]-p[0])*(c[0]-p[0]) + (c[1]-p[1])*(c[1]-p[1])))
	return d <= float64(c[2])
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
		for j := range bombs {
			if i == j {
				continue
			}
			graph[i][j] = isConnected(bombs[i], bombs[j])
		}
	}

	ans := 0
	for i := range bombs {
		ans = max(ans, bfs(i, graph))
	}

	return ans
}

func main() {
	testCases := []struct {
		bombs [][]int
		want  int
	}{
		{
			bombs: [][]int{{2, 1, 3}, {6, 1, 4}},
			want:  2,
		},
		{
			bombs: [][]int{{1, 1, 5}, {10, 10, 5}},
			want:  1,
		},
		{
			bombs: [][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}},
			want:  5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumDetonation(tc.bombs)
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
