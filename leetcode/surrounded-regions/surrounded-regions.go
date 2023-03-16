package main

// source: https://leetcode.com/problems/surrounded-regions/
func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	dirs := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	region := make(map[[2]int]bool)

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i < 0 || i >= n || j < 0 || j >= m {
			return false
		}
		if board[i][j] == 'X' || region[[2]int{i, j}] {
			return true
		}
		region[[2]int{i, j}] = true
		for _, d := range dirs {
			if !dfs(i+d[0], j+d[1]) {
				return false
			}
		}
		return true
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' && dfs(i, j) {
				for c := range region {
					board[c[0]][c[1]] = 'X'
				}
			}
			if len(region) > 0 {
				region = make(map[[2]int]bool)
			}
		}
	}
}
