package main

import (
	"fmt"
)

type pair struct {
	x, y int
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	n, m := len(grid1), len(grid1[0])

	// union-find on both islands
	ds := make(map[pair]pair, n*m)
	var find func(pair) pair
	var union func(pair, pair)
	union = func(x, y pair) {
		ds[find(x)] = find(y)
	}
	find = func(x pair) pair {
		if ds[x] != x {
			ds[x] = find(ds[x])
		}
		return ds[x]
	}

	for i := range grid2 {
		for j := range grid2[0] {
			ds[pair{i, j}] = pair{i, j}
		}
	}

	dirs := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	// union islands
	for i := range grid2 {
		for j := range grid2[0] {
			if grid2[i][j] == 1 {
				for _, dir := range dirs {
					if x, y := i+dir[0], j+dir[1]; x >= 0 && x < n && y >= 0 && y < m && grid2[x][y] == 1 {
						union(pair{i, j}, pair{x, y})
					}
				}
			}
		}
	}

	invalidIslands := make(map[pair]bool)
	for i := range grid2 {
		for j := range grid2[0] {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				invalidIslands[find(pair{i, j})] = true
			}
		}
	}

	// count valid islands
	ans := 0
	for i := range grid2 {
		for j := range grid2[0] {
			if grid2[i][j] == 1 && !invalidIslands[find(pair{i, j})] {
				ans++
				invalidIslands[find(pair{i, j})] = true
			}
		}
	}

	return ans
}

/*func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	n, m := len(grid1), len(grid1[0])
	dirs := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	// calculate
	var ans int

	var checkIsland func(i, j int)
	checkIsland = func(i, j int) {
		valid := !(grid1[i][j] == 0 && grid2[i][j] == 1)

		if grid1[i][j] == 1 {
			grid1[i][j] = 2
		}

		for _, dir := range dirs {
			if x, y := i+dir[0], j+dir[1]; x >= 0 && x < n && y >= 0 && y < m {
				//println(grid1[x][y], grid1[x][y] == 1)
				if grid1[x][y] == 1 {
					grid1[x][y] = 2
					valid = valid && checkIsland(x, y)
				}
				if grid1[x][y] == 0 && grid2[x][y] == 1 {
					return false
				}
			}
		}

		return valid
	}

	for i := range grid1 {
		for j := range grid1[0] {
			if grid1[i][j] == 1 {
				if checkIsland(i, j) {
					ans++
				}
			}
		}
	}

	return ans
}*/

func main() {
	testCases := []struct {
		grid1 [][]int
		grid2 [][]int
		want  int
	}{
		{
			grid1: [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
			grid2: [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}},
			want:  3,
		},
		{
			grid1: [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
			grid2: [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}},
			want:  2,
		},
		{
			grid1: [][]int{{1, 1}, {0, 0}},
			grid2: [][]int{{0, 1}, {0, 1}},
			want:  0,
		},

		{
			grid1: [][]int{{1, 1, 1, 1, 0, 0}, {1, 1, 0, 1, 0, 0}, {1, 0, 0, 1, 1, 1}, {1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0}, {1, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 0, 1}, {1, 0, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 0}, {1, 1, 1, 1, 1, 0}},
			grid2: [][]int{{1, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}, {1, 1, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1}, {1, 1, 0, 1, 1, 1}, {1, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1}, {1, 0, 0, 1, 0, 0}},
			want:  0,
		},
		{
			grid1: [][]int{{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1}, {1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1}, {0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0}, {0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}, {0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1}, {1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1}, {0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1}},
			grid2: [][]int{{1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1}, {1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1}, {1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1}, {0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0}, {0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1}, {1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1}, {1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1}, {1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0}, {0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0}},
			want:  6,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countSubIslands(tc.grid1, tc.grid2)
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
