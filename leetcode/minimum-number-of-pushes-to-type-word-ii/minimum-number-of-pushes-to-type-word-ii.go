package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	ans := 0
	cnt := make([]int, 27)

	for _, w := range word {
		cnt[int(w-'a')]++
	}

	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})

	b := 8
	p := 1

	for _, c := range cnt {
		if c == 0 {
			break
		}

		ans += p * c

		b--
		if b == 0 {
			b = 8
			p++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		word string
		want int
	}{
		{
			word: "abcde",
			want: 5,
		},
		{
			word: "xyzxyzxyzxyz",
			want: 12,
		},
		{
			word: "aabbccddeeffgghhiiiiii",
			want: 24,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumPushes(tc.word)
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
