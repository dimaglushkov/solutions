package main

import "fmt"

// source: https://leetcode.com/problems/unique-paths-iii/

var res = 0

func dfs(grid [][]int, i, j, z int) {
	if grid[i][j] == 2 {
		if z != -1 {
			return
		}
		res++
		return
	}

	grid[i][j] = 1

	if i+1 < len(grid) && (grid[i+1][j] == 0 || grid[i+1][j] == 2) {
		dfs(grid, i+1, j, z-1)
	}

	if j+1 < len(grid[0]) && (grid[i][j+1] == 0 || grid[i][j+1] == 2) {
		dfs(grid, i, j+1, z-1)
	}

	if i-1 > -1 && (grid[i-1][j] == 0 || grid[i-1][j] == 2) {
		dfs(grid, i-1, j, z-1)

	}

	if j-1 > -1 && (grid[i][j-1] == 0 || grid[i][j-1] == 2) {
		dfs(grid, i, j-1, z-1)
	}
	grid[i][j] = 0
}

func uniquePathsIII(grid [][]int) int {
	var iStart, jStart, zeros int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				zeros++
			} else if grid[i][j] == 1 {
				iStart, jStart = i, j
			}
		}
	}

	res = 0
	dfs(grid, iStart, jStart, zeros)
	return res
}

func main() {
	// Example 1
	var grid1 = [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}
	fmt.Println("Expected: 2	Output: ", uniquePathsIII(grid1))

	// Example 2
	var grid2 = [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}
	fmt.Println("Expected: 4	Output: ", uniquePathsIII(grid2))

	// Example 3
	var grid3 = [][]int{{0, 1}, {2, 0}}
	fmt.Println("Expected: 0	Output: ", uniquePathsIII(grid3))

}
