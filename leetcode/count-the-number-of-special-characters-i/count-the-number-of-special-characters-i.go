package main

// source: https://leetcode.com/problems/count-the-number-of-special-characters-i/
func numberOfSpecialChars(word string) int {
	cnt := make(map[byte]int)
	for i := range word {
		cnt[word[i]]++
	}

	ans := 0

	for i := range 26 {
		if cnt[byte(i+'a')] > 0 && cnt[byte(i+'A')] > 0 {
			ans++
		}
	}

	return ans
}
