package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/word-ladder-ii/

func findLaddersTLE(start string, end string, wordList []string) [][]string {
	m := make(map[string][]string)
	for _, w := range wordList {
		m[w] = make([]string, 0)
	}

	if _, ok := m[end]; !ok {
		return [][]string{}
	}

	// building graph
	m[start] = make([]string, 0)
	for w := range m {
		for k := range m {
			if w != k && k != start && oneCharDiff(w, k) {
				m[w] = append(m[w], k)
			}
		}
	}

	resPathLen := 501
	res := make([][]string, 0)
	path := make([]string, 0)

	var dfs func(s string)
	dfs = func(s string) {
		path = append(path, s)

		if len(path) > resPathLen {
			path = path[:len(path)-1]
			return
		}

		if s == end {
			if len(path) < resPathLen {
				resPathLen = len(path)
				res = [][]string{}
			}

			resPath := make([]string, len(path))
			copy(resPath, path)
			res = append(res, resPath)

			path = path[:len(path)-1]
			return
		}

		for _, w := range m[s] {
			if !inPath(w, path) {
				dfs(w)
			}
		}

		path = path[:len(path)-1]
	}

	dfs(start)

	return res
}

func inPath(w string, path []string) bool {
	for _, ws := range path {
		if ws == w {
			return true
		}
	}
	return false
}

func oneCharDiff(x, y string) bool {
	var diff bool
	for i := range x {
		if x[i] != y[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return true
}

func main() {
	// Example 1
	var beginWord1 string = "hit"
	var endWord1 string = "cog"
	var wordList1 = []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println("Expected: [[\"hit\",\"hot\",\"dot\",\"dog\",\"cog\"],[\"hit\",\"hot\",\"lot\",\"log\",\"cog\"]]	Output: ", findLadders(beginWord1, endWord1, wordList1))

	// Example 2
	var beginWord2 string = "hit"
	var endWord2 string = "cog"
	var wordList2 = []string{"hot", "dot", "dog", "lot", "log"}
	fmt.Println("Expected: []	Output: ", findLadders(beginWord2, endWord2, wordList2))

}
