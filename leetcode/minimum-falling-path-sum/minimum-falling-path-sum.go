package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/minimum-falling-path-sum/

func min(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] < x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func minTwo(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minFallingPathSum_(matrix [][]int) int {
	n := len(matrix)
	pathSum := make([][]int, n)
	for i := range pathSum {
		pathSum[i] = make([]int, n)
		for j := range pathSum[i] {
			pathSum[i][j] = math.MaxInt
		}
	}
	var dfs func(i, j, s int)
	dfs = func(i, j, s int) {
		if i >= n || s+matrix[i][j] >= pathSum[i][j] {
			return
		}
		s += matrix[i][j]
		pathSum[i][j] = min(pathSum[i][j], s)

		if nj := j - 1; nj >= 0 {
			dfs(i+1, nj, s)
		}
		dfs(i+1, j, s)
		if nj := j + 1; nj < n {
			dfs(i+1, nj, s)
		}
	}
	for j := range matrix[0] {
		dfs(0, j, 0)
	}

	return min(pathSum[n-1]...)
}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}

	for i := 1; i < n; i++ {
		matrix[i][0] = minTwo(matrix[i-1][0], matrix[i-1][1]) + matrix[i][0]
		for j := 1; j < n-1; j++ {
			matrix[i][j] = minTwo(minTwo(matrix[i-1][j-1], matrix[i-1][j]), matrix[i-1][j+1]) + matrix[i][j]
		}
		matrix[i][n-1] = minTwo(matrix[i-1][n-1], matrix[i-1][n-2]) + matrix[i][n-1]
	}
	return min(matrix[n-1]...)
}

func main() {
	testCases := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			want:   13,
		},
		{
			matrix: [][]int{{-19, 57}, {-40, -5}},
			want:   -59,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minFallingPathSum(tc.matrix)
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
