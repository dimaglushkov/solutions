package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/smallest-string-with-swaps/

func smallestStringWithSwaps(s string, pairs [][]int) string {
	parent := make([]int, len(s))
	var find func(int) int
	var union func(int, int)

	union = func(x, y int) {
		parent[find(x)] = find(y)
	}
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// building union-find structure
	for i := range parent {
		parent[i] = i
	}
	for i := range pairs {
		union(pairs[i][0], pairs[i][1])
	}

	// grouping union-find structures
	sets := make(map[int][]int)
	for i := range s {
		p := find(i)
		sets[p] = append(sets[p], i)
	}

	// sorting and building result
	var chars []rune
	res := []rune(s)
	//res := make([]rune, len(s))
	for _, set := range sets {
		sort.Ints(set)
		chars = make([]rune, len(set))
		for i := range set {
			chars[i] = res[set[i]]
		}
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		for i := range set {
			res[set[i]] = chars[i]
		}
	}

	return string(res)
}

func main() {
	// Example 1
	var s1 string = "dcab"
	var pairs1 = [][]int{{0, 3}, {1, 2}}
	fmt.Println("Expected: \"bacd\"	Output: ", smallestStringWithSwaps(s1, pairs1))

	// Example 2
	var s2 string = "dcab"
	var pairs2 = [][]int{{0, 3}, {1, 2}, {0, 2}}
	fmt.Println("Expected: \"abcd\"	Output: ", smallestStringWithSwaps(s2, pairs2))

	// Example 3
	var s3 string = "cba"
	var pairs3 = [][]int{{0, 1}, {1, 2}}
	fmt.Println("Expected: \"abc\"	Output: ", smallestStringWithSwaps(s3, pairs3))

}
