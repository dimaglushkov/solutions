package main

import "fmt"

// source: https://leetcode.com/problems/reverse-vowels-of-a-string/

func reverseVowels(s string) string {
	rs := []rune(s)
	vws := []rune("aeiouAEIOU")
	vowels := make(map[rune]bool, len(vws))
	for _, r := range vws {
		vowels[r] = true
	}

	l, r := 0, len(s)-1
	for l < r {
		for l < len(s) && !vowels[rs[l]] {
			l++
		}
		for r > 0 && !vowels[rs[r]] {
			r--
		}
		if l < r {
			rs[l], rs[r] = rs[r], rs[l]
		}
		l++
		r--

	}
	return string(rs)

}

func main() {
	// Example 1
	var s1 string = "hello"
	fmt.Println("Expected: holle	Output: ", reverseVowels(s1))

	// Example 2
	var s2 string = "leetcode"
	fmt.Println("Expected: leotcede	Output: ", reverseVowels(s2))

}
