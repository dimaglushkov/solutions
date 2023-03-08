package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])

	if target < matrix[0][0] || target > matrix[n-1][m-1] {
		return false
	}

	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if matrix[h][0] < target {
			i = h + 1
		} else {
			j = h
		}
	}

	if i == n || matrix[i][0] > target {
		i--
	}

	row := i
	i, j = 0, m
	for i < j {
		h := int(uint(i+j) >> 1)
		if matrix[row][h] < target {
			i = h + 1
		} else {
			j = h
		}
	}

	return i < m && matrix[row][i] == target
}

func main() {
	testCases := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{{1, 3, 5}},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{{1, 3}},
			target: 3,
			want:   true,
		},

		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
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
