package main

import "fmt"

// source: https://leetcode.com/problems/flood-fill/

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	oldCol := image[sr][sc]
	if oldCol == color {
		return image
	}

	mr, mc := len(image), len(image[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= mr || c < 0 || c >= mc || image[r][c] != oldCol {
			return
		}
		image[r][c] = color
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	dfs(sr, sc)

	return image
}

func main() {
	// Example 1
	var image1 = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	var sr1 int = 1
	var sc1 int = 1
	var color1 int = 2
	fmt.Println("Expected: [[2,2,2],[2,2,0],[2,0,1]]	Output: ", floodFill(image1, sr1, sc1, color1))

	// Example 2
	var image2 = [][]int{{0, 0, 0}, {0, 0, 0}}
	var sr2 int = 0
	var sc2 int = 0
	var color2 int = 0
	fmt.Println("Expected: [[0,0,0],[0,0,0]]	Output: ", floodFill(image2, sr2, sc2, color2))

}
