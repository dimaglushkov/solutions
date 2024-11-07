package main

import (
	"fmt"
)

func largestCombination(candidates []int) int {
	bits := make([]int, 16)
	ans := 0

	for _, c := range candidates {
		i := 0
		for c > 0 {
			if c%2 == 1 {
				bits[i]++
				ans = max(ans, bits[i])
			}
			c /= 2
			i++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		candidates []int
		want       int
	}{
		{
			candidates: []int{16, 17, 71, 62, 12, 24, 14},
			want:       4,
		},
		{
			candidates: []int{8, 8},
			want:       2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := largestCombination(tc.candidates)
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
