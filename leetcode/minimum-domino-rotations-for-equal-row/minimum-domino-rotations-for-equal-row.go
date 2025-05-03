package main

import (
	"fmt"
)

func minDominoRotations(tops []int, bottoms []int) int {
	var ct, cb, cs [7]int

	for i := range tops {
		ct[tops[i]]++
		cb[bottoms[i]]++

		if tops[i] == bottoms[i] {
			cs[tops[i]]++
		}
	}

	for i := 1; i < 7; i++ {
		if ct[i]+cb[i]-cs[i] == len(tops) {
			if ct[i] < cb[i] {
				return ct[i] - cs[i]
			}

			return cb[i] - cs[i]
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		tops    []int
		bottoms []int
		want    int
	}{
		{
			tops:    []int{1, 2, 3, 4, 6},
			bottoms: []int{6, 6, 6, 6, 5},
			want:    1,
		},
		{
			tops:    []int{1, 2, 1, 1, 1, 2, 2, 2},
			bottoms: []int{2, 1, 2, 2, 2, 2, 2, 2},
			want:    1,
		},
		{
			tops:    []int{2, 1, 2, 4, 2, 2},
			bottoms: []int{5, 2, 6, 2, 3, 2},
			want:    2,
		},
		{
			tops:    []int{3, 5, 1, 2, 3},
			bottoms: []int{3, 6, 3, 3, 4},
			want:    -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minDominoRotations(tc.tops, tc.bottoms)
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
