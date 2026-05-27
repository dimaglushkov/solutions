package main

// source: https://leetcode.com/problems/count-the-number-of-special-characters-ii/
func numberOfSpecialChars(word string) int {
	lastLower := map[rune]int{}
	firstUpper := map[rune]int{}

	for i, c := range word {
		if c >= 'a' && c <= 'z' {
			lastLower[c] = i
		} else {
			lower := c + 32

			if _, exists := firstUpper[lower]; !exists {
				firstUpper[lower] = i
			}
		}
	}

	ans := 0

	for c, i := range lastLower {
		if upperIdx, exists := firstUpper[c]; exists && i < upperIdx {
			ans++
		}
	}

	return ans
}
