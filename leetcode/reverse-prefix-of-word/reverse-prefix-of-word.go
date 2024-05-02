package main

import (
	"fmt"
	"strings"
)

func reverse(y string) string {
	x := []byte(y)
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}

	return string(x)
}

func reversePrefix(word string, ch byte) string {
	i := strings.Index(word, string(ch))

	if i == -1 {
		return word
	}

	return reverse(word[:i+1]) + word[i+1:]
}

func main() {
	testCases := []struct {
		word string
		ch   byte
		want string
	}{
		{
			word: "abcdefd",
			ch:   'd',
			want: "dcbaefd",
		},
		{
			word: "xyxzxe",
			ch:   'z',
			want: "zxyxxe",
		},
		{
			word: "abcd",
			ch:   'z',
			want: "abcd",
		},
		{
			word: "z",
			ch:   'z',
			want: "z",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := reversePrefix(tc.word, tc.ch)
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
