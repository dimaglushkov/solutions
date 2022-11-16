package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			if newEnd := intervals[i][1]; newEnd > end {
				end = newEnd
			}
		}

	}
	res = append(res, []int{start, end})

	return res
}

func main() {
	testCases := []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := merge(tc.intervals)
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
