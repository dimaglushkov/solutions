package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/max-increase-to-keep-city-skyline/

func min(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] < x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	mh, mv := make([]int, n), make([]int, n)
	for i := range grid {
		for j := range grid {
			if grid[i][j] > mh[i] {
				mh[i] = grid[i][j]
			}
			if grid[i][j] > mv[j] {
				mv[j] = grid[i][j]
			}
		}
	}
	res := 0

	for i := range grid {
		for j := range grid {
			res += min(mh[i]-grid[i][j], mv[j]-grid[i][j])
		}
	}

	return res
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}},
			want: 35,
		},
		{
			grid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxIncreaseKeepingSkyline(tc.grid)
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
