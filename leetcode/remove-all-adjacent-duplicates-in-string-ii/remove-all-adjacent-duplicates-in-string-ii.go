package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

type Stack struct {
	values []rune
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val rune) {
	s.values = append(s.values, val)

}

func (s *Stack) Pop() (val rune, ok bool) {
	if len(s.values) == 0 {
		return val, false
	}
	val = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val, true
}

func removeDuplicates(s string, k int) string {
	var values []rune

	push := func(val rune) {
		values = append(values, val)

		if len(values) < k {
			return
		}

		for i := 1; i < k; i++ {
			if values[len(values)-1-i] != val {
				return
			}
		}
		values = values[:len(values)-k]
	}

	for _, r := range s {
		push(r)
	}

	return string(values)
}

// TLE
func removeDuplicates_(s string, k int) string {
	var cnt int
	var runes []rune
	var updated = true

	for updated {
		runes = []rune(s)
		updated = false
		cnt = 1
		for i := 1; i < len(runes); i++ {
			if runes[i] == runes[i-1] {
				cnt++
			} else {
				cnt = 1
			}
			if cnt == k {
				updated = true
				cnt = 0
				for j := i - k + 1; j <= i; j++ {
					runes[j] = '.'
				}
			}
		}
		s = strings.ReplaceAll(string(runes), ".", "")
	}
	return string(runes)
}

func main() {
	// Example 2
	var s2 string = "deeedbbcccbdaa"
	var k2 int = 3
	fmt.Println("Expected: \"aa\"	Output: ", removeDuplicates(s2, k2))

	// Example 3
	var s3 string = "pbbcggttciiippooaais"
	var k3 int = 2
	fmt.Println("Expected: \"ps\"	Output: ", removeDuplicates(s3, k3))

	// Example 1
	var s1 string = "abcd"
	var k1 int = 2
	fmt.Println("Expected: \"abcd\"	Output: ", removeDuplicates(s1, k1))

}
