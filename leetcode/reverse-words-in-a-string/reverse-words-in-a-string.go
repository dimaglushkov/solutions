package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/reverse-words-in-a-string/

func reverse(x []string) {
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
}

func reverseWords(s string) string {
	words := make([]string, 0)
	for _, word := range strings.Split(strings.Trim(s, " "), " ") {
		if len(word) != 0 {
			words = append(words, word)
		}
	}
	reverse(words)
	return strings.Join(words, " ")
}

func main() {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			s:    "a good   example",
			want: "example good a",
		},
	}

	for _, tc := range testCases {
		x := reverseWords(tc.s)
		status := "ERROR"
		// intentionally using this way of comparison
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}

}
