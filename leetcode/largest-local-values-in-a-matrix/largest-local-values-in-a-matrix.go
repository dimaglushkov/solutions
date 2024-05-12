package main

import (
	"fmt"
)

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		ans[i] = make([]int, n-2)
	}

	dirs := [][2]int{{0, -1}, {-1, -1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {-1, 1}, {-1, 0}}
	for i := 0; i < n-2; i++ {
		si := i + 1
		for j := 0; j < n-2; j++ {
			sj := j + 1
			m := grid[si][sj]
			for _, d := range dirs {
				m = max(grid[si+d[0]][sj+d[1]], m)
			}
			ans[i][j] = m
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		grid [][]int
		want [][]int
	}{
		{
			grid: [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}},
			want: [][]int{{9, 9}, {8, 6}},
		},
		{
			grid: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
			want: [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := largestLocal(tc.grid)
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
