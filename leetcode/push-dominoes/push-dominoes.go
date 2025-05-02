package main

import (
	"fmt"
)

func move(dominoes string) string {
	d := []byte(dominoes)

	for i := 0; i < len(dominoes); i += 1 {
		if d[i] != '.' || i-1 >= 0 && dominoes[i-1] == 'R' && i+1 < len(dominoes) && dominoes[i+1] == 'L' {
			continue
		}

		if i-1 >= 0 && dominoes[i-1] == 'R' {
			d[i] = 'R'
		}
		if i+1 < len(dominoes) && dominoes[i+1] == 'L' {
			d[i] = 'L'
		}
	}

	return string(d)
}

func pushDominoes(dominoes string) string {
	prev := dominoes
	dominoes = move(dominoes)

	for prev != dominoes {
		prev = dominoes
		dominoes = move(dominoes)
	}

	return dominoes
}

func main() {
	testCases := []struct {
		dominoes string
		want     string
	}{
		{
			dominoes: ".L.R...LR..L..",
			want:     "LL.RR.LLRRLL..",
		},
		{
			dominoes: "RR.L",
			want:     "RR.L",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := pushDominoes(tc.dominoes)
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
