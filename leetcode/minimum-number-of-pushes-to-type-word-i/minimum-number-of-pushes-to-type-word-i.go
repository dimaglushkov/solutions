package main

import (
	"fmt"
)

func minimumPushes(word string) int {
	ans := 0
	for i := 0; i < len(word); i++ {
		ans += (i / 8) + 1

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
			word: "xycdefghij",
			want: 12,
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
