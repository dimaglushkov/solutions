package main

import (
	"fmt"
)

func rotatedDigits(n int) int {
	ans := 0

	required := map[int]bool{2: true, 5: true, 6: true, 9: true}
	disallowed := map[int]bool{3: true, 4: true, 7: true}

	for i := range n + 1 {
		x := i
		hasDisallowed := false
		hasRequired := false
		for x > 0 {
			d := x % 10
			if ok := disallowed[d]; ok {
				hasDisallowed = true
				break
			}
			if ok := required[d]; ok {
				hasRequired = true
			}
			x /= 10
		}

		if hasRequired && !hasDisallowed {
			ans++
		}

	}

	return ans
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    857,
			want: 247,
		},
		{
			n:    2,
			want: 1,
		},
		{
			n:    10,
			want: 4,
		},
		{
			n:    1,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := rotatedDigits(tc.n)
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
