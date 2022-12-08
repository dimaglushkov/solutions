package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/

func dividePlayers(skill []int) int64 {
	n := len(skill)
	s := 0

	m := make(map[int]int, n)
	for _, i := range skill {
		m[i]++
		s += i
	}
	ts := s / (n / 2)

	var res int64

	for _, x := range skill {
		if m[x] == 0 {
			continue
		}
		y := ts - x
		if m[y] != m[x] {
			return -1
		}
		c := m[x]
		if x == y {
			c /= 2
		}
		res += int64(c * (x * y))
		m[x], m[y] = 0, 0
	}

	return res
}

func main() {
	testCases := []struct {
		skill []int
		want  int64
	}{
		{
			skill: []int{3, 2, 5, 1, 3, 4},
			want:  22,
		},
		{
			skill: []int{3, 4},
			want:  12,
		},
		{
			skill: []int{1, 1, 2, 3},
			want:  -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := dividePlayers(tc.skill)
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
