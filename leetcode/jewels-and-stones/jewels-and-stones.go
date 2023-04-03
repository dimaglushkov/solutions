package main

import (
	"fmt"
)

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[rune]bool)
	for _, r := range jewels {
		m[r] = true
	}
	ans := 0
	for _, r := range stones {
		if m[r] {
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		jewels string
		stones string
		want   int
	}{
		{
			jewels: "aA",
			stones: "aAAbbbb",
			want:   3,
		},
		{
			jewels: "z",
			stones: "ZZ",
			want:   0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numJewelsInStones(tc.jewels, tc.stones)
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
