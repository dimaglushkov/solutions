package main

// source: https://leetcode.com/problems/roman-to-integer/

// Keep summing up chars values unless the next chars value is bigger than the current one
// In this case just subtract current val
func romanToInt(str string) int {
	alphabet := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s := []rune(str)
	res := alphabet[s[len(s)-1]]
	for i := range s[:len(s)-1] {
		if alphabet[s[i+1]] > alphabet[s[i]] {
			res -= alphabet[s[i]]
		} else {
			res += alphabet[s[i]]
		}
	}
	return res
}
