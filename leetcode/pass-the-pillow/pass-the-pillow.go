package main

import (
	"fmt"
)

func passThePillow(n int, time int) int {
	chunks := time / (n - 1)
	if chunks%2 == 0 {
		return (time % (n - 1)) + 1
	} else {
		return n - (time % (n - 1))
	}
}

func main() {
	testCases := []struct {
		n    int
		time int
		want int
	}{
		{
			n:    4,
			time: 5,
			want: 2,
		},
		{
			n:    3,
			time: 2,
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := passThePillow(tc.n, tc.time)
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
