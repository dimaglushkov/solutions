package main

import (
	"fmt"
	"strings"
)
// source: https://leetcode.com/problems/word-pattern/

func wordPattern(pattern string, s string) bool {
    dict := make(map[byte]string)
	cnt := make(map[string]int)
	words := strings.Split(s, " ")
	for _, w := range words{
		cnt[w] ++
	}
	n := len(words)
	if n < len(pattern) {
		return false
	}
	for i := range pattern {
		if w, ok := dict[pattern[i]]; !ok {
			dict[pattern[i]] = words[i]
			cnt[words[i]]--
		} else {
			if w != words[i] {
				return false
			} else {
				cnt[words[i]]--
			}
		}
	}

	dw := make(map[string]bool)
	for _, w := range dict {
		if dw[w] {
			return false
		}
		dw[w] = true
	}

	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}

	return true
}

func main() {
    testCases := []struct {
		pattern string
		s string
		want bool
    }{
		{
			pattern:  "syys",
			s:  "a abc abc a",
			want:  true,
		},
		{
			pattern:  "abba",
			s:  "dog cat cat dog",
			want:  true,
		},
		{
			pattern:  "abba",
			s:  "dog cat cat fish",
			want:  false,
		},
		{
			pattern:  "aaaa",
			s:  "dog cat cat dog",
			want:  false,
		},
	}


    successes := 0
    for _, tc := range testCases {
        x := wordPattern(tc.pattern, tc.s)
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
        fmt.Printf("===\nFAIL: %d tests failed\n", l - successes)
    }

}
