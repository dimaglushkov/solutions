package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	t, l, b, r := 0, 0, 0, 0
	ans := make([]int, 0, m*n)
	for len(ans) < m*n {
		// left -> right
		for i := l; t+b < n && i < m-r; i++ {
			ans = append(ans, matrix[t][i])
		}
		t++

		// up -> down
		for i := t; l+r < m && i < n-l; i++ {
			ans = append(ans, matrix[i][m-r-1])
		}
		r++

		// right -> left
		for i := m - r - 1; t+b < n && i >= l; i-- {
			ans = append(ans, matrix[n-b-1][i])
		}
		b++

		// down -> up
		for i := n - b - 1; l+r < m && i >= t; i-- {
			ans = append(ans, matrix[i][l])
		}
		l++
	}
	return ans
}

func main() {
	testCases := []struct {
		matrix [][]int
		want   []int
	}{

		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := spiralOrder(tc.matrix)
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
