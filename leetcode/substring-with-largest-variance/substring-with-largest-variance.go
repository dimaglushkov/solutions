package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largestVariance(s string) int {
	ans := 0
	chars := make(map[rune]struct{})
	for _, c := range s {
		chars[c] = struct{}{}
	}

	for x, _ := range chars {
		for y, _ := range chars {
			if x == y {
				continue
			}
			totalX := 0
			totalY := 0
			for _, c := range s {
				switch c {
				case x:
					totalX += 1
				case y:
					totalY += 1
				default:
					continue
				}

				if totalY > 0 {
					ans = max(ans, totalX-totalY)
				} else {
					ans = max(ans, totalX-1)
				}

				if totalX-totalY < 0 {
					totalX = 0
					totalY = 0
				}

			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "aababbb",
			want: 3,
		},
		{
			s:    "abcde",
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := largestVariance(tc.s)
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
