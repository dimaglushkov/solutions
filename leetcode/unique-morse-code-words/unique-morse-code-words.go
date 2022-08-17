package main

import "fmt"

// source: https://leetcode.com/problems/unique-morse-code-words/

func morsefy(s string) string {
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var res string

	for _, r := range s {
		res += codes[r-'a']
	}

	return res
}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]bool)
	for _, w := range words {
		m[morsefy(w)] = true
	}

	return len(m)
}

func main() {
	// Example 1
	var words1 = []string{"gin", "zen", "gig", "msg"}
	fmt.Println("Expected: 2	Output: ", uniqueMorseRepresentations(words1))

	// Example 2
	var words2 = []string{"a"}
	fmt.Println("Expected: 1	Output: ", uniqueMorseRepresentations(words2))

}
