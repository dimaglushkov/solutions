package main

import "fmt"

// source: https://leetcode.com/problems/max-area-of-island/

func maxAreaOfIsland(grid [][]int) (res int) {
	var tmp int
	var findIslandSize func(i, j int) int
	findIslandSize = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = -1
		return 1 + findIslandSize(i-1, j) + findIslandSize(i+1, j) + findIslandSize(i, j-1) + findIslandSize(i, j+1)
	}
	for i := range grid {
		for j := range grid[i] {
			if tmp = findIslandSize(i, j); tmp > res {
				res = tmp
			}
		}
	}
	return res
}

func main() {
	// Example 1
	var grid1 = [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println("Expected: 6	Output: ", maxAreaOfIsland(grid1))

	// Example 2
	var grid2 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println("Expected: 0	Output: ", maxAreaOfIsland(grid2))

}
