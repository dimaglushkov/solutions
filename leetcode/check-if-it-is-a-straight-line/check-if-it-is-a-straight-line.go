package main

import (
	"fmt"
)

func checkStraightLine(coords [][]int) bool {
	dx0 := coords[1][0] - coords[0][0]
	dy0 := coords[1][1] - coords[0][1]
	for i := 1; i < len(coords)-1; i++ {
		dx := coords[i+1][0] - coords[i][0]
		dy := coords[i+1][1] - coords[i][1]
		if dy*dx0 != dy0*dx {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		coordinates [][]int
		want        bool
	}{
		{
			coordinates: [][]int{{0, 0}, {0, 1}, {0, -1}},
			want:        true,
		},
		{
			coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
			want:        true,
		},
		{
			coordinates: [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
			want:        false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := checkStraightLine(tc.coordinates)
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
