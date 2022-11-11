package main

import "fmt"

// source: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

func removeDuplicates(s string) string {
	res := make([]rune, 0, len(s))

	for _, letter := range s {
		if len(res) > 0 && res[len(res)-1] == letter {
			res = res[:len(res)-1]
			continue
		}

		res = append(res, letter)
	}

	return string(res)
}

func main() {
	// Example 1
	var s1 string = "abbaca"
	fmt.Println("Expected: ca	Output: ", removeDuplicates(s1))

	// Example 2
	var s2 string = "azxxzy"
	fmt.Println("Expected: ay	Output: ", removeDuplicates(s2))

}
