package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/determine-if-two-strings-are-close/

func closeStrings(word1 string, word2 string) bool {
	const a = 'a'
	var wc1, wc2 = make([]int, 26), make([]int, 26)
	for _, l := range word1 {
		wc1[l-a]++
	}
	for _, l := range word2 {
		wc2[l-a]++
	}
	for i := range wc1 {
		if wc1[i] != 0 && wc2[i] == 0 || wc1[i] == 0 && wc2[i] != 0 {
			return false
		}
	}

	sort.Ints(wc1)
	sort.Ints(wc2)

	for i := range wc1 {
		if wc1[i] != wc2[i] {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		word1 string
		word2 string
		want  bool
	}{
		{
			word1: "abc",
			word2: "bca",
			want:  true,
		},
		{
			word1: "a",
			word2: "aa",
			want:  false,
		},
		{
			word1: "cabbba",
			word2: "abbccc",
			want:  true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := closeStrings(tc.word1, tc.word2)
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
