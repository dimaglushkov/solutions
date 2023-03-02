package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/string-compression/

func compress(chars []byte) int {
	// is - index in source slice, ir - index in resulting slice
	is, ir := 0, 0
	n := len(chars)

	for is < n {
		// moving through all the repeating characters
		i := is + 1
		for i < n && chars[i] == chars[is] {
			i++
		}

		// updating array
		// setting the next letter in resulting array
		chars[ir] = chars[is]
		ir++
		if cnt := i - is; cnt > 1 {
			digits := []byte(strconv.FormatInt(int64(cnt), 10))
			for j := 0; j < len(digits); j++ {
				chars[ir+j] = digits[j]
			}
			ir += len(digits)
		}
		// moving source index to the next unique letter
		is = i
	}

	chars = chars[:ir]
	return ir
}

func main() {
	testCases := []struct {
		chars []byte
		want  int
	}{
		{
			chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			want:  6,
		},
		{
			chars: []byte{'a'},
			want:  1,
		},
		{
			chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			want:  4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := compress(tc.chars)
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
