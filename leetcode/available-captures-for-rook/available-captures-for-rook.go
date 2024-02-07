package main

import (
	"fmt"
)

func numRookCaptures(board [][]byte) int {
	ans := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'R' {
				for k := i; k < len(board); k++ {
					if board[k][j] == 'B' {
						break
					}
					if board[k][j] == 'p' {
						ans++
						break
					}
				}
				for k := i; k >= 0; k-- {
					if board[k][j] == 'B' {
						break
					}
					if board[k][j] == 'p' {
						ans++
						break
					}
				}

				for k := j; k < len(board[0]); k++ {
					if board[i][k] == 'B' {
						break
					}
					if board[i][k] == 'p' {
						ans++
						break
					}
				}
				for k := j; k >= 0; k-- {
					if board[i][k] == 'B' {
						break
					}
					if board[i][k] == 'p' {
						ans++
						break
					}
				}

				break
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		board [][]byte
		want  int
	}{
		{
			board: [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'R', '.', '.', '.', 'p'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}},
			want:  3,
		},
		{
			board: [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'B', 'R', 'B', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}},
			want:  0,
		},
		{
			board: [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'p', 'p', '.', 'R', '.', 'p', 'B', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}},
			want:  3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numRookCaptures(tc.board)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
