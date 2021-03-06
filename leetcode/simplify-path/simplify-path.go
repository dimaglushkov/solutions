package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/simplify-path/

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
