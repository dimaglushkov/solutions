package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/lexicographically-smallest-equivalent-string/
func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]int, 26)
	var find func(int) int
	var union func(int, int)
	union = func(x, y int) {
		parent[find(x)] = find(y)
	}
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	for i := range s1 {
		i1, i2 := int(s1[i]-'a'), int(s2[i]-'a')
		if i1 < i2 {
			union(i2, i1)
		} else {
			union(i1, i2)
		}
	}

	// minimize
	m := make(map[int]int)
	for i := range parent {
		fi := find(i)
		if _, ok := m[fi]; !ok {
			m[find(i)] = i
		}
	}

	res := make([]byte, len(baseStr))
	for i := range baseStr {
		res[i] = byte(m[find(int(baseStr[i]-'a'))]) + 'a'
	}

	return string(res)
}

/*
97  - a
98  - b
99  - c
100 - d
101 - e
...
113 - q
114 - r
115 - s
116 - t
117 - u
118 - v
119 - w
120 - x
121 - y
122 - z
*/

func main() {
	testCases := []struct {
		s1      string
		s2      string
		baseStr string
		want    string
	}{
		{
			s1:      "gmerjboftfnqseogigpdnlocmmhskigdtednfnjtlcrdpcjkbj",
			s2:      "fnnafafhqkitbcdlkpiloiienikjiikdfcafisejgeldprcmhd",
			baseStr: "ezrqfyguivmybqcsvibuvtajdvamcdjpmgcbvieegpyzdcypcx",
			want:    "azaaayauavayaaaavaauvaaaavaaaaaaaaaavaaaaayzaayaax",
		},
		{
			s1:      "abc",
			s2:      "cde",
			baseStr: "eed",
			want:    "aab",
		},
		{
			s1:      "parker",
			s2:      "morris",
			baseStr: "parser",
			want:    "makkek",
		},
		{
			s1:      "hello",
			s2:      "world",
			baseStr: "hold",
			want:    "hdld",
		},
		{
			s1:      "leetcode",
			s2:      "programs",
			baseStr: "sourcecode",
			want:    "aauaaaaada",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := smallestEquivalentString(tc.s1, tc.s2, tc.baseStr)
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
