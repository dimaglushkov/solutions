package main

import "fmt"

// source: https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	var res int
	var extractIsland func(i, j int)
	extractIsland = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		extractIsland(i-1, j)
		extractIsland(i+1, j)
		extractIsland(i, j-1)
		extractIsland(i, j+1)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				extractIsland(i, j)
			}
		}
	}

	return res
}

func main() {
	// Example 1
	var grid1 = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("Expected: 1	Output: ", numIslands(grid1))

	// Example 2
	var grid2 = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Expected: 3	Output: ", numIslands(grid2))

}
