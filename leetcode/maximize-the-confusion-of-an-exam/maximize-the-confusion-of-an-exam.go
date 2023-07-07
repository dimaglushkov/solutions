package main

import (
	"fmt"
)

func maxConsecutiveAnswers(answerKey string, k int) int {
	ans := 0
	m := make(map[rune]int)
	for i, r := range answerKey {
		m[r]++
		if m['T'] <= k || m['F'] <= k {
			ans++
		} else {
			m[rune(answerKey[i-ans])]--
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		answerKey string
		k         int
		want      int
	}{
		{
			answerKey: "TTFF",
			k:         2,
			want:      4,
		},
		{
			answerKey: "TFFT",
			k:         1,
			want:      3,
		},
		{
			answerKey: "TTFTTFTT",
			k:         1,
			want:      5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxConsecutiveAnswers(tc.answerKey, tc.k)
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
