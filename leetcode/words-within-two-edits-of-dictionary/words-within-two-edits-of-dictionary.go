package main

import (
	"fmt"
)

func twoEditWords(queries []string, dictionary []string) []string {
	var ans []string

	for _, q := range queries {
		for _, d := range dictionary {
			swaps := 0
			for i := range q {
				if q[i] != d[i] {
					swaps++
					if swaps > 2 {
						break
					}
				}
			}

			if swaps < 3 {
				ans = append(ans, q)
				break
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		queries    []string
		dictionary []string
		want       []string
	}{
		{
			queries:    []string{"word", "note", "ants", "wood"},
			dictionary: []string{"wood", "joke", "moat"},
			want:       []string{"word", "note", "wood"},
		},
		{
			queries:    []string{"yes"},
			dictionary: []string{"not"},
			want:       []string{},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := twoEditWords(tc.queries, tc.dictionary)
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
