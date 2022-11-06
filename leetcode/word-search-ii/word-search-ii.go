package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/word-search-ii/

type Trie struct {
	Children [26]*Trie
	Word     string
}

func (t *Trie) insert(word string) {
	for _, ch := range word {
		x := ch - 'a'
		if t.Children[x] == nil {
			t.Children[x] = new(Trie)
		}
		t = t.Children[x]
	}
	t.Word = word
}

func findWords(board [][]byte, words []string) []string {
	n, m := len(board), len(board[0])
	q := len(words)
	res := make([]string, 0, q)
	trie := new(Trie)
	for _, w := range words {
		trie.insert(w)
	}

	var dfs func(i, j int, trie *Trie)
	dfs = func(i, j int, t *Trie) {
		if i < 0 || i >= n || j < 0 || j >= m || board[i][j] == 0 {
			return
		}
		x := board[i][j]
		t = t.Children[x-'a']
		if t == nil {
			return
		}
		if t.Word != "" {
			res = append(res, t.Word)
			t.Word = ""
		}

		board[i][j] = 0
		dfs(i+1, j, t)
		dfs(i-1, j, t)
		dfs(i, j-1, t)
		dfs(i, j+1, t)
		board[i][j] = x
	}

	for i := range board {
		for j := range board[0] {
			dfs(i, j, trie)
		}
	}

	return res
}

func main() {
	var board = [][]byte{{'a'}}
	var words = []string{"ab"}
	fmt.Println("Expected: []	Output: ", findWords(board, words))

	// Example 1
	var board1 = [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	var words1 = []string{"oath", "pea", "eat", "rain"}
	fmt.Println("Expected: [eat,oath]	Output: ", findWords(board1, words1))

	// Example 2
	var board2 = [][]byte{{'a', 'b'}, {'c', 'd'}}
	var words2 = []string{"abcb"}
	fmt.Println("Expected: []	Output: ", findWords(board2, words2))

}
