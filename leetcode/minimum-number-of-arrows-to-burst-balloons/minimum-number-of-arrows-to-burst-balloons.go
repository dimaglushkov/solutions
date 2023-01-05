package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	n := len(points)
	res := n
	end := points[0][1]

	for i := 1; i < n; i++ {
		b := points[i]
		if b[0] > end {
			end = b[1]
		} else {
			res--
			if end > b[1] {
				end = b[1]
			}
		}

	}
	return res
}

func main() {
	testCases := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want:   2,
		},
		{
			points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want:   4,
		},
		{
			points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			want:   2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findMinArrowShots(tc.points)
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
