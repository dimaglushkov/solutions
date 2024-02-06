package main

import (
	"fmt"
)

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}

func main() {
	testCases := []struct {
		matrix [][]int
		want   bool
	}{
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}},
			want:   true,
		},
		{
			matrix: [][]int{{1, 2}, {2, 2}},
			want:   false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isToeplitzMatrix(tc.matrix)
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
