package main

import (
	"fmt"
)

func newMatrix(row, col int) [][]bool {
	matr := make([][]bool, row)
	for i := range matr {
		matr[i] = make([]bool, col)
	}

	return matr
}

func getMaximumGold(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var ans int

	var dfs func(i, j int, visited [][]bool) int
	dfs = func(i, j int, visited [][]bool) int {
		tm := 0
		visited[i][j] = true
		dirs := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
		for _, d := range dirs {
			if di, dj := i+d[0], j+d[1]; di >= 0 && di < n && dj >= 0 && dj < m && grid[di][dj] != 0 && !visited[di][dj] {
				tm = max(tm, dfs(di, dj, visited))
			}
		}

		visited[i][j] = false
		return grid[i][j] + tm
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				v := newMatrix(n, m)
				x := dfs(i, j, v)
				ans = max(x, ans)
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}},
			want: 24,
		},
		{
			grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}},
			want: 28,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := getMaximumGold(tc.grid)
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
