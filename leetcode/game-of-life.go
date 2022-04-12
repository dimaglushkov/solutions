package main

import "fmt"

// source: https://leetcode.com/problems/game-of-life/
func newCellValue(board [][]int, i, j, m, n int) int {
	neighborsCnt := 0
	if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1] > 0 {
		neighborsCnt++
	}
	if i-1 >= 0 && board[i-1][j] > 0 {
		neighborsCnt++
	}
	if i-1 >= 0 && j+1 < n && board[i-1][j+1] > 0 {
		neighborsCnt++
	}

	if j-1 >= 0 && board[i][j-1] > 0 {
		neighborsCnt++
	}
	if j+1 < n && board[i][j+1] > 0 {
		neighborsCnt++
	}

	if i+1 < m && j-1 >= 0 && board[i+1][j-1] > 0 {
		neighborsCnt++
	}
	if i+1 < m && board[i+1][j] > 0 {
		neighborsCnt++
	}
	if i+1 < m && j+1 < n && board[i+1][j+1] > 0 {
		neighborsCnt++
	}

	if board[i][j] == 1 && (neighborsCnt != 2 && neighborsCnt != 3) {
		// returning 2 if prev value was 1 but the new one is 0
		return 2
	}
	if board[i][j] == 0 && neighborsCnt == 3 {
		// returning -1 if prev value was 0 but the new one is 1
		return -1
	}
	return board[i][j]
}

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = newCellValue(board, i, j, m, n)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}

func main() {
	// Example 1
	var board1 = [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board1)
	fmt.Println("Expected: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]	Output: ", board1)

	// Example 2
	var board2 = [][]int{{1, 1}, {1, 0}}
	gameOfLife(board2)
	fmt.Println("Expected: [[1,1],[1,1]]	Output: ", board2)

}
