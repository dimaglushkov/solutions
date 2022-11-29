package main

import (
	"fmt"

	"github.com/dimaglushkov/solutions/ADS/stack"
)

// source: https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	st := stack.NewGenStack[int32]()
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			st.Push(c)
		} else if sc, ok := st.Pop(); ok || pairs[sc] != c {
			return false
		}
	}
	return st.Len() == 0
}

func main() {
	// Example 1
	var s1 string = "()"
	fmt.Println("Expected: true	Output: ", isValid(s1))

	// Example 2
	var s2 string = "()[]{}"
	fmt.Println("Expected: true	Output: ", isValid(s2))

	// Example 3
	var s3 string = "(]"
	fmt.Println("Expected: false	Output: ", isValid(s3))

	// Example 4
	var s4 string = "((()))"
	fmt.Println("Expected: true	Output: ", isValid(s4))

	// Example 5
	var s5 string = "())"
	fmt.Println("Expected: false	Output: ", isValid(s5))

	// Example 6
	var s6 string = "([{})]"
	fmt.Println("Expected: false	Output: ", isValid(s6))

}
