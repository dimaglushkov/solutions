package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/

// queue is used in order to implment BFS
type queue []int

func (q *queue) push(x int) {
	*q = append(*q, x)
}
func (q *queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func minReorder(n int, connections [][]int) int {
	var ans int
	var q queue

	// graph stores all the roads of a city
	graph := make([][][]int, n)
	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c)
		graph[c[1]] = append(graph[c[1]], c)
	}

	// visited keep track of visited cities
	visited := make([]bool, n)

	// start from city 0
	q.push(0)
	for len(q) > 0 {
		x := q.pop()
		if visited[x] {
			continue
		}
		visited[x] = true

		// iterate over all the roads of a given city
		for _, node := range graph[x] {
			// check if the given roads is directed towards city x. if it is,
			// then check if it comes from a previously visited city. if both
			// conditions are satisfied, then the road is in the wrong direction
			if node[1] == x && visited[node[0]] {
				ans++
			}

			if node[0] == x {
				q.push(node[1])
			} else {
				q.push(node[0])
			}
		}
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
			n:           6,
			connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			want:        3,
		},
		{
			n:           5,
			connections: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			want:        2,
		},
		{
			n:           3,
			connections: [][]int{{1, 0}, {2, 0}},
			want:        0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minReorder(tc.n, tc.connections)
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
