package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	var letters [27]bool
	for _, l := range allowed {
		letters[int(l-'a')] = true
	}

	ans := 0
	for _, w := range words {
		ok := true
		for _, l := range w {
			if !letters[int(l-'a')] {
				ok = false
				break
			}
		}

		if ok {
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		allowed string
		words   []string
		want    int
	}{
		{
			allowed: "ab",
			words:   []string{"ad", "bd", "aaab", "baa", "badab"},
			want:    2,
		},
		{
			allowed: "abc",
			words:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			want:    7,
		},
		{
			allowed: "cad",
			words:   []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			want:    4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countConsistentStrings(tc.allowed, tc.words)
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
