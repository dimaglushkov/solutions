package main

import (
	"fmt"
)

func furthestDistanceFromOrigin(moves string) int {
	var cu, s int
	for _, i := range moves {
		if i == 'R' {
			s++
		} else if i == 'L' {
			s--
		} else {
			cu++
		}
	}

	if s >= 0 {
		return s + cu
	} else {
		return -s + cu
	}
}

func main() {
	testCases := []struct {
		moves string
		want  int
	}{
		{
			moves: "L_RL__R",
			want:  3,
		},
		{
			moves: "_R__LL_",
			want:  5,
		},
		{
			moves: "_______",
			want:  7,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := furthestDistanceFromOrigin(tc.moves)
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
