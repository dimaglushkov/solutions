package main

import (
	"fmt"
	"sort"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	cnt := map[byte]int{}
	for _, l := range licensePlate {
		if l >= 'a' && l <= 'z' || l >= 'A' && l <= 'Z' {
			cnt[strings.ToLower(string(l))[0]]++
		}
	}

	for _, w := range words {
		tmp := map[byte]int{}
		for _, l := range w {
			tmp[byte(l)]++
		}

		ok := true
		for l, c := range cnt {
			if tmp[l] < c {
				ok = false
				break
			}
		}

		if ok {
			return w
		}
	}

	return ""
}

func main() {
	testCases := []struct {
		licensePlate string
		words        []string
		want         string
	}{
		{
			licensePlate: "1s3 PSt",
			words:        []string{"step", "steps", "stripe", "stepple"},
			want:         "steps",
		},
		{
			licensePlate: "1s3 456",
			words:        []string{"looks", "pest", "stew", "show"},
			want:         "pest",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := shortestCompletingWord(tc.licensePlate, tc.words)
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
