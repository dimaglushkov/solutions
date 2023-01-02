package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// source: https://leetcode.com/problems/detect-capital/

func detectCapitalUse(word string) bool {
	capitalize := func(s string) string {
		if len(s) == 0 {
			return ""
		}
		res := make([]byte, 1, len(word))
		res[0] = byte(unicode.ToUpper(rune(s[0])))
		res = append(res, bytes.ToLower([]byte(s)[1:])...)
		return string(res)
	}

	return strings.ToLower(word) == word || strings.ToUpper(word) == word || capitalize(word) == word
}

func main() {
	testCases := []struct {
		word string
		want bool
	}{
		{
			word: "USA",
			want: true,
		},
		{
			word: "FlaG",
			want: false,
		},
		{
			word: "Google",
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := detectCapitalUse(tc.word)
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
