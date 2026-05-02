package main

import (
	"fmt"
)

func countCollisions(directions string) int {
	d := []byte(directions)
	ans := 0

	rcnt := 0
	stuck := false

	for _, c := range d {
		switch c {
		case 'R':
			rcnt++
			stuck = false
		case 'L':
			if stuck {
				ans++
			}
			if rcnt > 0 {
				ans += rcnt + 1
				rcnt = 0
				stuck = true
			}

		case 'S':
			if rcnt > 0 {
				ans += rcnt
				rcnt = 0
			}
			stuck = true
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		directions string
		want       int
	}{
		{
			directions: "RRRL",
			want:       4,
		},
		{
			directions: "RLLL",
			want:       4,
		},
		{
			directions: "RLRSLL",
			want:       5,
		},
		{
			directions: "LLRR",
			want:       0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countCollisions(tc.directions)
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
