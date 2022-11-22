package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/basic-calculator/

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	stack := make([]bool, 0)
	push := func(x bool) {
		stack = append(stack, x)
	}
	pop := func() bool {
		n := len(stack)
		x := stack[n-1]
		stack = stack[:n-1]
		return x
	}
	top := func() bool {
		return stack[len(stack)-1]
	}
	push(true)

	sb := strings.Builder{}
	for i, r := range s {
		switch r {
		case '(':
			if i > 0 && s[i-1] == '-' {
				push(!top())
			} else {
				push(top())
			}
		case ')':
			pop()
		case '-':
			if top() {
				sb.WriteString(" -")
			} else {
				sb.WriteByte(' ')
			}
		case '+':
			if !top() {
				sb.WriteString(" -")
			} else {
				sb.WriteByte(' ')
			}
		default:
			sb.WriteRune(r)
		}
	}
	res := 0

	for _, x := range strings.Split(sb.String(), " ") {
		res += atoi(x)
	}
	return res
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "1 + 1",
			want: 2,
		},
		{
			s:    " 2-1 + 2 ",
			want: 3,
		},
		{
			s:    "(1+(4+5+2)-3)+(6+8)",
			want: 23,
		},
		{
			s:    "-(4-5-(7+(91+3-(17+4))))",
			want: 81,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := calculate(tc.s)
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
