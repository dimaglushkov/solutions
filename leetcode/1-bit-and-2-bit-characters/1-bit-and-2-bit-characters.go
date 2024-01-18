package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	last, before := len(bits)-1, -1
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 0 {
			before = i
			break
		}
	}

	return (last-before+1)%2 == 0
}

func main() {
	testCases := []struct {
		bits []int
		want bool
	}{
		{
			bits: []int{0, 0},
			want: true,
		},
		{
			bits: []int{1, 0, 0},
			want: true,
		},
		{
			bits: []int{0},
			want: true,
		},
		{
			bits: []int{1, 0},
			want: false,
		},
		{
			bits: []int{1, 1, 1, 0},
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isOneBitCharacter(tc.bits)
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
