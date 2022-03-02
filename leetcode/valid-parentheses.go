package main

import "fmt"

// source: https://leetcode.com/problems/valid-parentheses/

type stack struct {
	values []rune
	Len    int
}

func (s *stack) Push(val rune) {
	s.values = append(s.values, val)
	s.Len++
}

func (s *stack) Pop() (val rune, err error) {
	if s.Len == 0 {
		return val, fmt.Errorf("can't pop from an empty stack")
	}
	s.Len--
	val = s.values[s.Len]
	s.values = s.values[:s.Len]
	return val, err
}

func isValid(s string) bool {
	stack := stack{}
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack.Push(c)
		} else if sc, err := stack.Pop(); err != nil || pairs[sc] != c {
			return false
		}
	}
	return stack.Len == 0
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
