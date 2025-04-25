package main

import "fmt"

func dfs(grid [][]int, x, y int, seenBorder bool) (int, bool) {
	var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m := len(grid)
	n := len(grid[0])

	if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != 1 {
		return 0, false
	}

	grid[x][y] = 2
	val := 1
	isBorder := x == 0 || y == 0 || x == m-1 || y == n-1 || seenBorder

	for i := 0; i < 4; i++ {
		subval, subIsBorder := dfs(grid, x+dir[i][0], y+dir[i][1], isBorder)
		isBorder = isBorder || subIsBorder
		val += subval
	}

	return val, isBorder
}

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if t, v := dfs(grid, i, j, false); !v {
					ans += t
				}
			}
		}
	}

	return ans
}

func main() {
	//fmt.Println(numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))
	//fmt.Println(numEnclaves([][]int{{0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {1, 1, 0, 0, 0, 1, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {0, 1, 1, 0, 0, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1, 0, 0, 1, 0}, {0, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, {0, 0, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}))
	//fmt.Println(numEnclaves([][]int{
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 1, 1, 1, 1, 1, 0, 0, 0, 0},
	//	{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
	//	{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
	//	{0, 0, 1, 0, 0, 0, 0, 1, 0, 1},
	//	{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//}))
	fmt.Println(numEnclaves([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}))
}
