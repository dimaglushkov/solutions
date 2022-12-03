package main

import (
	"fmt"
	"sort"
	"strings"
)

// source: https://leetcode.com/problems/sort-characters-by-frequency/

func frequencySort(s string) string {
	cnt := make([]int, 123)
	let := make([]byte, 123)
	for i := range let {
		let[i] = byte(i)
	}
	for _, i := range s {
		cnt[i]++
	}
	sort.Slice(let, func(i, j int) bool {
		return cnt[let[i]] > cnt[let[j]]
	})
	sb := strings.Builder{}
	for i := 0; i < 123; i++ {
		for j := 0; j < cnt[let[i]]; j++ {
			sb.WriteByte(let[i])
		}
	}

	return sb.String()
}

func main() {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "tree",
			want: "eert",
		},
		{
			s:    "cccaaa",
			want: "aaaccc",
		},
		{
			s:    "Aabb",
			want: "bbAa",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := frequencySort(tc.s)
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
