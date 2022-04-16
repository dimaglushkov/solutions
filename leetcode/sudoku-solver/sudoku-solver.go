package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/sudoku-solver/

// Sudoku is the classic example for backtracking algorithm
// I've implemented it with what I believe is a small preprocessing to significantly
// decrease execution time by manually checking for simple cells before running backtrack

type SudokuBoard [9][9]int8

func NewSudokuBoard(board [][]byte) SudokuBoard {
	sb := SudokuBoard{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != byte('.') {
				cellVal, _ := strconv.ParseInt(string(board[i][j]), 10, 8)
				sb[i][j] = int8(cellVal)
			} else {
				sb[i][j] = -1
			}
		}
	}
	return sb
}

func guessValue(sb SudokuBoard, i int, j int, v int8) SudokuBoard {
	sb[i][j] = v
	return sb
}

func backtrack(sb SudokuBoard) (SudokuBoard, bool) {
	var posVal []int8
	var nextUnknownCell [2]int

	if nextUnknownCell = sb.getNextUnknownCell(); nextUnknownCell[0] == -1 {
		return sb, true
	}
	if posVal = sb.getPossibleValues(nextUnknownCell[0], nextUnknownCell[1]); len(posVal) == 0 {
		return sb, false
	}

	for _, pv := range posVal {
		temp, status := backtrack(guessValue(sb, nextUnknownCell[0], nextUnknownCell[1], pv))
		if status {
			return temp, true
		}
	}
	return sb, false
}

func (sb *SudokuBoard) getNextUnknownCell() [2]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sb[i][j] == -1 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

func (sb *SudokuBoard) getPossibleValues(i int, j int) []int8 {
	var posVal = make([]int8, 0, 10)
	var isValPossible [10]bool
	for n := 1; n < 10; n++ {
		isValPossible[n] = true
	}

	squareVerOff, squareHorOff := i/3*3, j/3*3
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			if sb[m+squareVerOff][n+squareHorOff] != -1 {
				isValPossible[sb[m+squareVerOff][n+squareHorOff]] = false
			}
		}
	}

	for m := 0; m < 9; m++ {
		if sb[i][m] != -1 {
			isValPossible[sb[i][m]] = false
		}
		if sb[m][j] != -1 {
			isValPossible[sb[m][j]] = false
		}
	}

	for m := 1; m < 10; m++ {
		if isValPossible[m] {
			posVal = append(posVal, int8(m))
		}
	}

	return posVal
}

func (sb *SudokuBoard) fillSimpleCells(i int, j int) {
	if sb[i][j] != -1 {
		return
	}

	posVal := sb.getPossibleValues(i, j)
	sqi := i/3*3 + j/3

	if len(posVal) == 1 {
		sb[i][j] = posVal[0]
		for ind := range posVal {
			ci, cj := ind/10, ind%10
			csqi := ci/3*3 + cj/3
			if csqi == sqi || i == ci || j == cj {
				sb.fillSimpleCells(ci, cj)
			}
		}
	}
}

func (sb *SudokuBoard) Solve() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sb[i][j] == -1 {
				sb.fillSimpleCells(i, j)
			}
		}
	}
	*sb, _ = backtrack(*sb)
}

func (sb *SudokuBoard) Print() {
	result := make([]string, 9)
	for i := 0; i < 9; i++ {
		var values = make([]string, 9)
		for j := 0; j < 9; j++ {
			values[j] = fmt.Sprintf("\"%d\"", sb[i][j])
		}

		result[i] = fmt.Sprintf("[%s]", strings.Join(values, ","))
	}
	fmt.Printf("[%s]", strings.Join(result, "\n,"))
}

func (sb *SudokuBoard) SetBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = strconv.FormatInt(int64(sb[i][j]), 10)[0]
		}
	}
}

func solveSudoku(board [][]byte) {
	ss := NewSudokuBoard(board)
	ss.Solve()
	ss.SetBoard(board)
}

func main() {
	//Example 1
	board1 := [][]byte{
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
	ss := NewSudokuBoard(board1)
	ss.Solve()
	ss.SetBoard(board1)
	fmt.Println("Expected:\n[[\"5\",\"3\",\"4\",\"6\",\"7\",\"8\",\"9\",\"1\",\"2\"],\n[\"6\",\"7\",\"2\",\"1\",\"9\",\"5\",\"3\",\"4\",\"8\"],\n[\"1\",\"9\",\"8\",\"3\",\"4\",\"2\",\"5\",\"6\",\"7\"],\n[\"8\",\"5\",\"9\",\"7\",\"6\",\"1\",\"4\",\"2\",\"3\"],\n[\"4\",\"2\",\"6\",\"8\",\"5\",\"3\",\"7\",\"9\",\"1\"],\n[\"7\",\"1\",\"3\",\"9\",\"2\",\"4\",\"8\",\"5\",\"6\"],\n[\"9\",\"6\",\"1\",\"5\",\"3\",\"7\",\"2\",\"8\",\"4\"],\n[\"2\",\"8\",\"7\",\"4\",\"1\",\"9\",\"6\",\"3\",\"5\"],\n[\"3\",\"4\",\"5\",\"2\",\"8\",\"6\",\"1\",\"7\",\"9\"]]	\nOutput: \n")
	ss.Print()

}
