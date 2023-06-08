package main

import (
	"fmt"
)

func countNegatives(grid [][]int) int {
	ans := 0
	n := len(grid[0])
	i := n - 1

	for _, row := range grid {
		for i >= 0 && row[i] < 0 {
			i--
		}
		ans += n - i - 1
	}
	return ans
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			want: 8,
		},
		{
			grid: [][]int{{3, 2}, {1, 0}},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countNegatives(tc.grid)
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
