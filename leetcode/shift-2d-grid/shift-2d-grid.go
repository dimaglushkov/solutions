package main

import "fmt"

// source: https://leetcode.com/problems/shift-2d-grid/

func shiftGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ii := (i + (j+k)/m) % n
			jj := (j + k) % m
			res[ii][jj] = grid[i][j]
		}
	}
	return res
}

func main() {
	grid1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k1 := 1
	fmt.Println("Expected: [[9,1,2],[3,4,5],[6,7,8]]	Output: ", shiftGrid(grid1, k1))

	grid2 := [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}
	k2 := 4
	fmt.Println("Expected: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]	Output: ", shiftGrid(grid2, k2))

	grid3 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k3 := 9
	fmt.Println("Expected: [[1,2,3],[4,5,6],[7,8,9]]	Output: ", shiftGrid(grid3, k3))

}
