package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/word-search/

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	var dfs func(i, j int, w string) bool
	dfs = func(i, j int, w string) bool {
		if len(w) == 0 {
			return true
		}
		if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != w[0] {
			return false
		}

		t := board[i][j]
		board[i][j] = '.'
		res := dfs(i-1, j, w[1:]) || dfs(i+1, j, w[1:]) || dfs(i, j-1, w[1:]) || dfs(i, j+1, w[1:])
		board[i][j] = t
		return res
	}

	for i := range board {
		for j := range board[i] {
			if dfs(i, j, word) {
				return true
			}
		}
	}
	return false
}

func main() {
	testCases := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
			want:  true,
		},
		{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "SEE",
			want:  true,
		},
		{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCB",
			want:  false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := exist(tc.board, tc.word)
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
