package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var board [8]string
	for i := 0; i < 8; i++ {
		fmt.Fscan(in, &board[i])
	}

	var x, y int

	for i := 1; i < 7; i++ {
		for j := 1; j < 7; j++ {
			if board[i][j] == '#' &&
				board[i-1][j-1] == '#' &&
				board[i+1][j-1] == '#' &&
				board[i-1][j+1] == '#' &&
				board[i+1][j+1] == '#' {
				x, y = i, j
			}
		}
	}

	return fmt.Sprintf("%d %d", x+1, y+1)
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}

// util functions
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func makeMatrix(n, m int) [][]int {
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, m)
	}
	return x
}
