package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	var ans int64
	graph := make([][]int, n)
	for _, r := range roads {
		graph[r[0]] = append(graph[r[0]], r[1])
		graph[r[1]] = append(graph[r[1]], r[0])
	}

	var dfs func(node, prev int) int64
	dfs = func(node, prev int) int64 {
		var cnt int64 = 1
		for _, edge := range graph[node] {
			if edge != prev {
				cnt += dfs(edge, node)
			}
		}

		if node != 0 {
			ans += int64(math.Ceil(float64(cnt) / float64(seats)))
		}
		return cnt
	}

	dfs(0, -1)

	return ans
}

func main() {
	testCases := []struct {
		roads [][]int
		seats int
		want  int64
	}{
		{
			roads: [][]int{{0, 1}, {0, 2}, {0, 3}},
			seats: 5,
			want:  3,
		},
		{
			roads: [][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}},
			seats: 2,
			want:  7,
		},
		{
			roads: [][]int{},
			seats: 1,
			want:  0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumFuelCost(tc.roads, tc.seats)
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
