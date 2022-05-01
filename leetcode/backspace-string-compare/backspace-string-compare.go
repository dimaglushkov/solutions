package main

import "fmt"

// source: https://leetcode.com/problems/backspace-string-compare/

func backspaceCompare(s string, t string) bool {
	push := func(st *[]rune, val rune) {
		*st = append(*st, val)
	}

	pop := func(st *[]rune) rune {
		if len(*st) == 0 {
			return -1
		}
		res := (*st)[len(*st)-1]
		*st = (*st)[:len(*st)-1]
		return res
	}

	var st1, st2 = make([]rune, 0, len(s)), make([]rune, 0, len(t))
	for _, r := range s {
		if r == '#' {
			pop(&st1)
		} else {
			push(&st1, r)
		}
	}

	for _, r := range t {
		if r == '#' {
			pop(&st2)
		} else {
			push(&st2, r)
		}
	}

	return string(st1) == string(st2)

}

func main() {
	// Example 4
	s4 := "xywrrmp"
	t4 := "xywrrmu#p"
	fmt.Println("Expected: true	Output: ", backspaceCompare(s4, t4))

	// Example 1
	var s1 string = "ab#c"
	var t1 string = "ad#c"
	fmt.Println("Expected: true	Output: ", backspaceCompare(s1, t1))

	// Example 2
	var s2 string = "ab##"
	var t2 string = "c#d#"
	fmt.Println("Expected: true	Output: ", backspaceCompare(s2, t2))

	// Example 3
	var s3 string = "a#c"
	var t3 string = "b"
	fmt.Println("Expected: false	Output: ", backspaceCompare(s3, t3))

}
