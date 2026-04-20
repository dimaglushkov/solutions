package main

import (
	"fmt"
)

func maxDistance(colors []int) int {
	i, j := 0, 0

	for j < len(colors) && colors[j] == colors[i] {
		j++
	}

	ans := j - i

	for t := j; t < len(colors); t++ {
		if colors[t] == colors[i] {
			ans = max(ans, t-j)
		} else {
			ans = max(ans, t-i)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		colors []int
		want   int
	}{
		{
			colors: []int{1, 1, 1, 6, 1, 1, 1},
			want:   3,
		},
		{
			colors: []int{1, 8, 3, 8, 3},
			want:   4,
		},
		{
			colors: []int{0, 1},
			want:   1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxDistance(tc.colors)
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
