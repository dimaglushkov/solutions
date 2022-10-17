package main

import "fmt"

// source: https://leetcode.com/problems/check-if-the-sentence-is-pangram/

func checkIfPangram(sentence string) bool {
	l := make(map[rune]bool)
	for _, r := range sentence {
		l[r] = true
		if len(l) == 26 {
			return true
		}
	}
	return len(l) == 26
}

func main() {
	// Example 1
	var sentence1 string = "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println("Expected: true	Output: ", checkIfPangram(sentence1))

	// Example 2
	var sentence2 string = "leetcode"
	fmt.Println("Expected: false	Output: ", checkIfPangram(sentence2))

}
