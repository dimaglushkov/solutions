package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/verifying-an-alien-dictionary/

func isAlienSorted(words []string, order string) bool {
	a := make(map[byte]int, 26)
	for i, b := range order {
		a[byte(b)] = i
	}

	less := func(w1, w2 string) bool {
		l1, l2 := len(w1), len(w2)
		for i := 0; i < l1 && i < l2; i++ {
			if w1[i] != w2[i] {
				if a[w1[i]] < a[w2[i]] {
					return true
				} else {
					return false
				}
			}
		}

		if l1 < l2 {
			return true
		}
		return false
	}

	return sort.SliceIsSorted(words, func(i, j int) bool {
		return less(words[i], words[j])
	})
}

func main() {
	testCases := []struct {
		words []string
		order string
		want  bool
	}{
		{
			words: []string{"ubg", "kwh"},
			order: "qcipyamwvdjtesbghlorufnkzx",
			want:  true,
		},
		{
			words: []string{"hello", "leetcode"},
			order: "hlabcdefgijkmnopqrstuvwxyz",
			want:  true,
		},
		{
			words: []string{"word", "world", "row"},
			order: "worldabcefghijkmnpqstuvxyz",
			want:  false,
		},
		{
			words: []string{"apple", "app"},
			order: "abcdefghijklmnopqrstuvwxyz",
			want:  false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isAlienSorted(tc.words, tc.order)
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
