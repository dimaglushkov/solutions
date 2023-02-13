package main

import "strings"

// source: https://leetcode.com/problems/number-of-segments-in-a-string/

func countSegments(s string) int {
	if len(s) == 0 || strings.Count(s, " ") == len(s) {
		return 0
	}
	ans := 0
	var in bool
	if s[0] == ' ' {
		in = false
	} else {
		in = true
	}

	for i := 1; i < len(s); i++ {
		if s[i] == ' ' && in {
			ans++
			in = false
		}
		if s[i] != ' ' && !in {
			in = true
		}
	}

	if in {
		ans++
	}

	return ans
}

func main() {
	countSegments("    foo    bar")
}
