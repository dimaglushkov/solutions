package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/simplify-path/

/*
// Stack-based implementation since this problem has "stack" in tags list
type StringStack struct {
	Values []string
	Len    int
}

func (s *StringStack) Push(val string) {
	s.Values = append(s.Values, val)
	s.Len++
}

func (s *StringStack) pop() (val string, err error) {
	if s.Len == 0 {
		return val, fmt.Errorf("can't pop from an empty stack")
	}
	s.Len--
	val = s.Values[s.Len]
	s.Values = s.Values[:s.Len]
	return val, err
}

func simplifyPath(path string) string {
	var stack StringStack
	for _, v := range strings.Split(path, "/") {
		if v == "." || v == "" {
			continue
		}
		if v == ".." {
			_, _ = stack.pop()
			continue
		}
		stack.Push(v)
	}
	return "/" + strings.Join(stack.Values, "/")
}*/

func simplifyPath(path string) string {
	var locations []string
	for _, v := range strings.Split(path, "/") {
		if v == "." || v == "" {
			continue
		}
		if v == ".." {
			if l := len(locations); l != 0 {
				locations = locations[:l-1]
			}
			continue
		}
		locations = append(locations, v)
	}
	return "/" + strings.Join(locations, "/")
}

func main() {
	// Example 1
	var path1 string = "/home/"
	fmt.Println("Expected: \"/home\"	Output: ", simplifyPath(path1))

	// Example 2
	var path2 string = "/../"
	fmt.Println("Expected: \"/\"	Output: ", simplifyPath(path2))

	// Example 4
	var path4 string = "/home/../home/../home/"
	fmt.Println("Expected: \"/home\"	Output: ", simplifyPath(path4))

	// Example 5
	var path5 string = "/home/root/Downloads/../../../../../../home/"
	fmt.Println("Expected: \"/home\"	Output: ", simplifyPath(path5))

	// Example 6
	var path6 string = "/etc/././././"
	fmt.Println("Expected: \"/etc\"	Output: ", simplifyPath(path6))

	// Example 7
	var path7 string = "/home/.././etc/././../home/./."
	fmt.Println("Expected: \"/home\"	Output: ", simplifyPath(path7))

	// Example 8
	var path8 string = "/home//foo/"
	fmt.Println("Expected: \"/home/foo\"	Output: ", simplifyPath(path8))

}
