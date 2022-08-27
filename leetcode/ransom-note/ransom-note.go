package main

import "fmt"

// source: https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int, 26)
	for _, r := range magazine {
		letters[r]++
	}

	for _, r := range ransomNote {
		if letters[r] == 0 {
			return false
		}
		letters[r]--
	}

	return true
}

func main() {
	// Example 1
	var ransomNote1 string = "a"
	var magazine1 string = "b"
	fmt.Println("Expected: false	Output: ", canConstruct(ransomNote1, magazine1))

	// Example 2
	var ransomNote2 string = "aa"
	var magazine2 string = "ab"
	fmt.Println("Expected: false	Output: ", canConstruct(ransomNote2, magazine2))

	// Example 3
	var ransomNote3 string = "aa"
	var magazine3 string = "aab"
	fmt.Println("Expected: true	Output: ", canConstruct(ransomNote3, magazine3))

}
