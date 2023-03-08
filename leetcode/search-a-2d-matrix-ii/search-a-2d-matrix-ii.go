package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/search-a-2d-matrix-ii/

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	for _, m := range matrix {
		x := sort.SearchInts(m, target)
		if x < n && m[x] == target {
			return true
		}
	}
	return false
}

func main() {
	testCases := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target: 5,
			want:   true,
		},
		{
			matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target: 20,
			want:   false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := searchMatrix(tc.matrix, tc.target)
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
