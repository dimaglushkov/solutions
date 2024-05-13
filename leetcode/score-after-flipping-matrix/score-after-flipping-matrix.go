package main

import (
	"fmt"
)

func matrixScore(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	flip := func(x int) int {
		if x == 1 {
			return 0
		}
		return 1
	}

	convert := func(x []int) int {
		ans := 0
		for i := range x {
			ans <<= 1
			ans |= x[i]
		}
		return ans
	}

	for i := range grid {
		if grid[i][0] == 0 {
			for j := range grid[i] {
				grid[i][j] = flip(grid[i][j])
			}
		}
	}

	for j := 1; j < m; j++ {
		cnt := 0
		for i := 0; i < n; i++ {
			if grid[i][j] == 1 {
				cnt++
			}
		}

		if cnt > n/2 {
			continue
		}

		for i := 0; i < n; i++ {
			grid[i][j] = flip(grid[i][j])
		}
	}

	ans := 0
	for i := range grid {
		ans += convert(grid[i])
	}

	return ans
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}},
			want: 39,
		},
		{
			grid: [][]int{{0}},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := matrixScore(tc.grid)
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
