package main

import "fmt"

// source: https://leetcode.com/problems/longest-common-prefix/

// Pretty straightforward solution, really nothing to comment
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs[1:] {
		for i := range prefix {
			if i >= len(str) || prefix[i] != str[i] {
				prefix = prefix[:i]
				break
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"doggo", "dog", "doggy"}
	fmt.Println(longestCommonPrefix(strs))
}
