package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minOperations(grid [][]int, x int) int {
	if len(grid) == 0 && len(grid[0]) == 0 {
		return 0
	}

	ans := 0
	values := make([]int, 0)
	for _, i := range grid {
		for _, j := range i {
			values = append(values, j)
		}
	}

	sort.Ints(values)
	t := values[len(values)/2]

	for _, i := range grid {
		for _, j := range i {
			if abs(j-t)%x != 0 {
				return -1
			}
			ans += abs(j-t) / x
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		grid [][]int
		x    int
		want int
	}{
		{
			grid: [][]int{{980, 476, 644, 56}, {644, 140, 812, 308}, {812, 812, 896, 560}, {728, 476, 56, 812}},
			x:    84,
			want: 45,
		},
		{
			grid: [][]int{{931, 128}, {639, 712}},
			x:    73,
			want: 12,
		},
		{
			grid: [][]int{{2, 4}, {6, 8}},
			x:    2,
			want: 4,
		},
		{
			grid: [][]int{{1, 5}, {2, 3}},
			x:    1,
			want: 5,
		},
		{
			grid: [][]int{{1, 2}, {3, 4}},
			x:    2,
			want: -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minOperations(tc.grid, tc.x)
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
