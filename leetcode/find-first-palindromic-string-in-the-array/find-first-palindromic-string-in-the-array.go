package main

import (
	"fmt"
)

func isPalindrome(word string) bool {
	n := len(word)
	for i := 0; i < n/2; i++ {
		if word[i] != word[n-1-i] {
			return false
		}
	}
	return true
}

func firstPalindrome(words []string) string {
	for _, w := range words {
		if isPalindrome(w) {
			return w
		}
	}

	return ""
}

func main() {
	testCases := []struct {
		words []string
		want  string
	}{
		{
			words: []string{"abc", "car", "ada", "racecar", "cool"},
			want:  "ada",
		},
		{
			words: []string{"notapalindrome", "racecar"},
			want:  "racecar",
		},
		{
			words: []string{"def", "ghi"},
			want:  "",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := firstPalindrome(tc.words)
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
