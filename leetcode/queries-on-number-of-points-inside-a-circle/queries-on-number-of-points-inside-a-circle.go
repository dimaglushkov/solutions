package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/

// (x-h)2 + (y-k)2 < r2
func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))
	for i := range queries {
		cnt := 0
		for j := range points {
			x, y := points[j][0], points[j][1]
			h, k, r := queries[i][0], queries[i][1], queries[i][2]
			if (x-h)*(x-h)+(y-k)*(y-k) <= r*r {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}

func main() {
	testCases := []struct {
		points  [][]int
		queries [][]int
		want    []int
	}{
		{
			points:  [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
			queries: [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			want:    []int{3, 2, 2},
		},
		{
			points:  [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			queries: [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}},
			want:    []int{2, 3, 2, 4},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countPoints(tc.points, tc.queries)
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
