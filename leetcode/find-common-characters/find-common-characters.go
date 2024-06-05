package main

import (
	"fmt"
)

func countLetter(word string) []int {
	cnt := make([]int, 26)
	for _, c := range word {
		cnt[c-'a']++
	}
	return cnt
}

func commonChars(words []string) []string {
	common := make([]int, 26)
	for i := range common {
		common[i] = 100
	}

	for _, word := range words {
		cnt := countLetter(word)
		for i := 0; i < 26; i++ {
			if cnt[i] < common[i] {
				common[i] = cnt[i]
			}
		}
	}

	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < common[i]; j++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}

func main() {
	testCases := []struct {
		words []string
		want  []string
	}{
		{
			words: []string{"bella", "label", "roller"},
			want:  []string{"e", "l", "l"},
		},
		{
			words: []string{"cool", "lock", "cook"},
			want:  []string{"c", "o"},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := commonChars(tc.words)
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
