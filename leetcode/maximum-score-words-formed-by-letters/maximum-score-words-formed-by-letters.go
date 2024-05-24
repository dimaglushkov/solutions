package main

import (
	"fmt"
)

func maxScoreWords(words []string, letters []byte, score []int) int {
	ans := 0

	cnt := make(map[byte]int)
	for _, l := range letters {
		cnt[l]++
	}

	var util func(start int, prevCnt map[byte]int, cur int)
	util = func(start int, prevCnt map[byte]int, cur int) {
		if cur > ans {
			ans = cur
		}

		if start == len(words) {
			return
		}

		for i := start; i < len(words); i++ {
			w := words[i]
			curScore := 0
			ok := true

			newCnt := make(map[byte]int)
			for k, v := range prevCnt {
				newCnt[k] = v
			}

			for _, l := range []byte(w) {
				if newCnt[l] > 0 {
					newCnt[l]--
					curScore += score[l-'a']
				} else {
					ok = false
					break
				}
			}

			if ok {
				util(i+1, newCnt, cur+curScore)
			}
		}
	}

	util(0, cnt, 0)

	return ans
}

func main() {
	testCases := []struct {
		words   []string
		letters []byte
		score   []int
		want    int
	}{
		{
			words:   []string{"dog", "cat", "dad", "good"},
			letters: []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
			score:   []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:    23,
		},
		{
			words:   []string{"xxxz", "ax", "bx", "cx"},
			letters: []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
			score:   []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10},
			want:    27,
		},
		{
			words:   []string{"leetcode"},
			letters: []byte{'l', 'e', 't', 'c', 'o', 'd'},
			score:   []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			want:    0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxScoreWords(tc.words, tc.letters, tc.score)
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
