package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/naming-a-company/

func distinctNames(ideas []string) int64 {
	var ans int64
	var words [26]map[string]bool
	for i := range words {
		words[i] = make(map[string]bool)
	}

	for _, idea := range ideas {
		bi := []byte(idea)
		words[int(bi[0]-'a')][string(bi[1:])] = true
	}

	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			var mut int
			for suf := range words[i] {
				if words[j][suf] {
					mut++
				}
			}
			ans += 2 * int64(len(words[i])-mut) * int64(len(words[j])-mut)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		ideas []string
		want  int64
	}{
		{
			ideas: []string{"coffee", "donuts", "time", "toffee"},
			want:  6,
		},
		{
			ideas: []string{"lack", "back"},
			want:  0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := distinctNames(tc.ideas)
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
