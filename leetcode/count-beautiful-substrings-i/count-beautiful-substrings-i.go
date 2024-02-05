package main

import (
	"fmt"
)

func beautifulSubstrings(s string, k int) int {
	ans := 0

	v := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	isV := make([]bool, len(s))

	for i := range s {
		isV[i] = v[rune(s[i])]
	}

	for i := 0; i < len(s); i++ {
		vc := 0
		for j := i; j < len(s); j++ {
			if isV[j] {
				vc++
			}
			if vc+vc == j-i+1 && vc*vc%k == 0 {
				ans++
			}
		}

	}

	return ans
}

func main() {
	testCases := []struct {
		s    string
		k    int
		want int
	}{
		{
			s:    "baeyh",
			k:    2,
			want: 2,
		},
		{
			s:    "eeebjoxxujuaeoqibd",
			k:    8,
			want: 4,
		},
		{
			s:    "abba",
			k:    1,
			want: 3,
		},
		{
			s:    "bcdf",
			k:    1,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := beautifulSubstrings(tc.s, tc.k)
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
