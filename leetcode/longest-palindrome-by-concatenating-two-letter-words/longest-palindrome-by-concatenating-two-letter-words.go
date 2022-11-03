package main

import "fmt"

// source: https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/

func longestPalindrome(words []string) int {
	res := 0
	ws := make(map[string]int, len(words))

	for _, w := range words {
		ww := string(w[1]) + string(w[0])
		if ws[ww] > 0 {
			res += 4
			ws[ww]--
		} else {
			ws[w]++
		}
	}

	for w, c := range ws {
		if w[0] == w[1] && c > 0 {
			res += 2
			break
		}
	}

	return res
}

func main() {
	var words0 = []string{"bb", "bb"}
	fmt.Println("Expected: 4	Output: ", longestPalindrome(words0))

	var words = []string{"qo", "fo", "fq", "qf", "fo", "ff", "qq", "qf", "of", "of", "oo", "of", "of", "qf", "qf", "of"}
	fmt.Println("Expected: 14	Output: ", longestPalindrome(words))

	// Example 1
	var words1 = []string{"lc", "cl", "gg"}
	fmt.Println("Expected: 6	Output: ", longestPalindrome(words1))

	// Example 2
	var words2 = []string{"ab", "ty", "yt", "lc", "cl", "ab"}
	fmt.Println("Expected: 8	Output: ", longestPalindrome(words2))

	// Example 3
	var words3 = []string{"cc", "ll", "xx"}
	fmt.Println("Expected: 2	Output: ", longestPalindrome(words3))

}
