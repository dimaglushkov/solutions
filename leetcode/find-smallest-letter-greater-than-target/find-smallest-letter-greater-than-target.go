package main

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	ans := letters[0]

	for _, l := range letters {
		if l > target {
			return l
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'a',
			want:    'c',
		},
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'c',
			want:    'f',
		},
		{
			letters: []byte{'x', 'x', 'y', 'y'},
			target:  'z',
			want:    'x',
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := nextGreatestLetter(tc.letters, tc.target)
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
