package main

import (
	"fmt"
	"sort"
)

func numMagicSquaresInside(grid [][]int) int {
	ans := 0

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if grid[i+1][j+1] != 5 {
				continue
			}

			if check(grid[i][j], grid[i+1][j], grid[i+2][j],
				grid[i][j+1], grid[i+1][j+1], grid[i+2][j+1],
				grid[i][j+2], grid[i+1][j+2], grid[i+2][j+2]) {
				ans++
			}
		}
	}

	return ans
}

func check(a, b, c, d, e, f, g, h, i int) bool {
	nums := []int{a, b, c, d, e, f, g, h, i}

	sort.Ints(nums)

	for x, v := range nums {
		if x+1 != v {
			return false
		}
	}

	return a+b+c == 15 && d+e+f == 15 && g+h+i == 15 && a+d+g == 15 && b+e+h == 15 && c+f+i == 15 && a+e+i == 15 && c+e+g == 15
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}},
			want: 1,
		},
		{
			grid: [][]int{{8}},
			want: 0,
		},
		{
			grid: [][]int{{11, 12, 13, 14, 15}, {21, 22, 23, 24, 25}, {31, 32, 33, 34, 35}, {41, 42, 43, 44, 45}},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numMagicSquaresInside(tc.grid)
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
