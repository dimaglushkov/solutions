package main

import (
	"fmt"
)

func countTriples(n int) int {
	var sq []int
	var ans int

	for i := 1; i <= n; i++ {
		sq = append(sq, i*i)
	}

	for _, c := range sq {
		for i := 1; i < c; i++ {
			for j := i + 1; j <= c; j++ {
				if i*i+j*j == c {
					ans++
				} else if i*i+j*j > c {
					break
				}
			}
		}
	}

	return ans * 2
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    5,
			want: 2,
		},
		{
			n:    10,
			want: 4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countTriples(tc.n)
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
