package main

import "fmt"

// source: https://leetcode.com/problems/valid-sudoku/

const arrSize = '9' + 1

// Going through every row, column, and square and counting every digit
// if any appears more than once, then sudoku is invalid
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var values [arrSize]int8
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[i*3+k][j*3+l] != '.' {
						if values[board[i*3+k][j*3+l]] > 0 {
							return false
						}
						values[board[i*3+k][j*3+l]]++
					}
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		var rowValues [arrSize]int8
		var colValues [arrSize]int8
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if rowValues[board[i][j]] > 0 {
					return false
				}
				rowValues[board[i][j]]++
			}
			if board[j][i] != '.' {
				if colValues[board[j][i]] > 0 {
					return false
				}
				colValues[board[j][i]]++
			}
		}
	}

	return true
}

func main() {
	// Example 1
	var board1 = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("Expected: true	Output: ", isValidSudoku(board1))

	// Example 2
	var board2 = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("Expected: false	Output: ", isValidSudoku(board2))

	// Example 3
	var board3 = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("Expected: true	Output: ", isValidSudoku(board3))

}
