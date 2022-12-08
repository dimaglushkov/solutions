package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/circular-sentence/

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	n := len(words)
	for i := 0; i < n-1; i++ {
		if words[i][len(words[i])-1] != words[i+1][0] {
			return false
		}
	}
	return words[n-1][len(words[n-1])-1] == words[0][0]
}

func main() {
	testCases := []struct {
		sentence string
		want     bool
	}{
		{
			sentence: "leetcode exercises sound delightful",
			want:     true,
		},
		{
			sentence: "eetcode",
			want:     true,
		},
		{
			sentence: "Leetcode is cool",
			want:     false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isCircularSentence(tc.sentence)
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
