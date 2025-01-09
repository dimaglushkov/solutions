package main

import "strings"

// source: https://leetcode.com/problems/counting-words-with-a-given-prefix/
func prefixCount(words []string, pref string) int {
	res := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			res++
		}
	}
	return res
}
