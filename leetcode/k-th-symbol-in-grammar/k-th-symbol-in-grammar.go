package main

import (
	"fmt"
)

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	if k%2 == 0 {
		return 1 - kthGrammar(n-1, (k+1)>>1)
	} else {
		return kthGrammar(n-1, (k+1)>>1)
	}
}

func main() {
	testCases := []struct {
		n    int
		k    int
		want int
	}{
		{
			n:    1,
			k:    1,
			want: 0,
		},
		{
			n:    2,
			k:    1,
			want: 0,
		},
		{
			n:    2,
			k:    2,
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := kthGrammar(tc.n, tc.k)
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
