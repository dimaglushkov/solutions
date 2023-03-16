package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/group-anagrams/

func sortStr(x string) string {
	r := []rune(x)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func groupAnagrams1(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		groups[sortStr(s)] = append(groups[sortStr(s)], s)
	}
	var ans [][]string
	for _, g := range groups {
		ans = append(ans, g)
	}
	return ans
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, w := range strs {
		arr := [26]int{}
		for _, c := range w {
			arr[c-'a'] += 1
		}
		groups[arr] = append(groups[arr], w)
	}
	var ans [][]string
	for _, group := range groups {
		ans = append(ans, group)
	}
	return ans

}

func main() {
	testCases := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := groupAnagrams(tc.strs)
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
