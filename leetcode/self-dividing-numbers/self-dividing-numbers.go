package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) []int {
	var ans []int

	check := func(x int) bool {
		xx := x
		for xx > 0 {
			dx := xx % 10
			if dx == 0 || x%dx != 0 {
				return false
			}
			xx /= 10
		}

		return true
	}

	for i := left; i <= right; i++ {
		if check(i) {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		left  int
		right int
		want  []int
	}{
		{
			left:  1,
			right: 22,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			left:  47,
			right: 85,
			want:  []int{48, 55, 66, 77},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := selfDividingNumbers(tc.left, tc.right)
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
