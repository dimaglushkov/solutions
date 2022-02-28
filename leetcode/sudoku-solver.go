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

// This exact implementation may not be the best in terms of code quality
// because I've tried other approaches and was basically rewriting a lot of code :/

type SudokuSolver struct {
	board [9][9]int8
}

func NewSudokuSolver(board [][]byte) SudokuSolver {
	ss := SudokuSolver{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != byte('.') {
				cellVal, err := strconv.ParseInt(string(board[i][j]), 10, 8)
				if err != nil {
					panic(err)
				}
				ss.board[i][j] = int8(cellVal)
			} else {
				ss.board[i][j] = -1
			}
		}
	}
	return ss
}

func guessValue(ss SudokuSolver, i int, j int, v int8) SudokuSolver {
	ss.board[i][j] = v
	return ss
}

func backtrack(ss SudokuSolver) (SudokuSolver, bool) {
	var posVal []int8
	var nextUnknownCell [2]int

	if nextUnknownCell = ss.getNextUnknownCell(); nextUnknownCell[0] == -1 {
		return ss, true
	}
	if posVal = ss.getPossibleValues(nextUnknownCell[0], nextUnknownCell[1]); len(posVal) == 0 {
		return ss, false
	}

	for _, pv := range posVal {
		temp, status := backtrack(guessValue(ss, nextUnknownCell[0], nextUnknownCell[1], pv))
		if status {
			return temp, true
		}
	}
	return ss, false
}

func (ss *SudokuSolver) getNextUnknownCell() [2]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if ss.board[i][j] == -1 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

func (ss *SudokuSolver) countUnknownCells() int {
	var count int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if ss.board[i][j] == -1 {
				count++
			}
		}
	}
	return count
}

func (ss *SudokuSolver) getPossibleValues(i int, j int) []int8 {
	var posVal = make([]int8, 0, 10)
	var isValPossible [10]bool
	for n := 1; n < 10; n++ {
		isValPossible[n] = true
	}

	squareVerOff, squareHorOff := i/3*3, j/3*3
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			if ss.board[m+squareVerOff][n+squareHorOff] != -1 {
				isValPossible[ss.board[m+squareVerOff][n+squareHorOff]] = false
			}
		}
	}

	for m := 0; m < 9; m++ {
		if ss.board[i][m] != -1 {
			isValPossible[ss.board[i][m]] = false
		}
		if ss.board[m][j] != -1 {
			isValPossible[ss.board[m][j]] = false
		}
	}

	for m := 1; m < 10; m++ {
		if isValPossible[m] {
			posVal = append(posVal, int8(m))
		}
	}

	return posVal
}

func (ss *SudokuSolver) fillSimpleCells(i int, j int) {
	if ss.board[i][j] != -1 {
		return
	}

	posVal := ss.getPossibleValues(i, j)
	sqi := i/3*3 + j/3

	if len(posVal) == 1 {
		ss.board[i][j] = posVal[0]
		for ind := range posVal {
			ci, cj := ind/10, ind%10
			csqi := ci/3*3 + cj/3
			if csqi == sqi || i == ci || j == cj {
				ss.fillSimpleCells(ci, cj)
			}
		}
	}
}

func (ss *SudokuSolver) Solve() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if ss.board[i][j] == -1 {
				ss.fillSimpleCells(i, j)
			}
		}
	}
	*ss, _ = backtrack(*ss)
}

func (ss *SudokuSolver) Print() {
	result := make([]string, 9)
	for i := 0; i < 9; i++ {
		var values = make([]string, 9)
		for j := 0; j < 9; j++ {
			values[j] = fmt.Sprintf("\"%d\"", ss.board[i][j])
		}

		result[i] = fmt.Sprintf("[%s]", strings.Join(values, ","))
	}
	fmt.Printf("[%s]", strings.Join(result, ","))
}

func (ss *SudokuSolver) SetBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = strconv.FormatInt(int64(ss.board[i][j]), 10)[0]
		}
	}
}

func solveSudoku(board [][]byte) {
	ss := NewSudokuSolver(board)
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
	solveSudoku(board1)
	//fmt.Println("Expected: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]	Output: ", solveSudoku(board1))

}
