package main

import (
	"fmt"
	strings "strings"
)

type st []string

func (q *st) push(x string) {
	*q = append(*q, x)
}
func (q *st) pop() string {
	n := len(*q)
	x := (*q)[n-1]
	*q = (*q)[:n-1]
	return x
}
func (q *st) top() string {
	return (*q)[len(*q)-1]
}

func reverseString(x string) string {
	arr := []rune(x)
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

	return string(arr)
}

func reverseParentheses(s string) string {
	var stack = st{}

	for _, c := range s {
		switch c {
		case '(':
			stack.push("(")

		case ')':
			var rev strings.Builder
			for len(stack) > 0 && stack.top() != "(" {
				str := stack.pop()
				rev.WriteString(reverseString(str))
			}
			stack.pop()
			stack.push(rev.String())

		default:
			stack.push(string(c))
		}
	}

	return strings.Join(stack, "")
}

//func reverseParentheses(s string) string {
//	d := 0
//	m := make(map[int][]int)
//
//	for i, c := range s {
//		switch c {
//		case '(':
//			d++
//		case ')':
//			d--
//		default:
//			if d%2 != 0 {
//				m[d] = append(m[d], i)
//			}
//		}
//	}
//
//	reversed := []byte(s)
//
//	for _, x := range m {
//		l, r := 0, len(x)-1
//		for l < r {
//			reversed[x[l]], reversed[x[r]] = reversed[x[r]], reversed[x[l]]
//			l++
//			r--
//		}
//	}
//
//	ans := make([]byte, 0, len(reversed))
//	for _, c := range reversed {
//		if c != '(' && c != ')' {
//			ans = append(ans, c)
//		}
//	}
//
//	return string(ans)
//}

func main() {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "a(bcdefghijkl(mno)p)q",
			want: "apmnolkjihgfedcbq",
		},
		{
			s:    "(abcd)",
			want: "dcba",
		},
		{
			s:    "(u(love)i)",
			want: "iloveu",
		},
		{
			s:    "(ed(et(oc))el)",
			want: "leetcode",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := reverseParentheses(tc.s)
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
