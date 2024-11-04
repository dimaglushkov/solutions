package main

// source: https://leetcode.com/problems/string-compression-iii/
func compressedString(word string) string {
	var ans []byte

	var i, c int

	for i < len(word) {
		c = 1

		for i+c < len(word) && c < 9 && word[i+c] == word[i] {
			c++
		}
		ans = append(ans, '0'+byte(c), word[i])
		i += c + 1
	}

	return string(ans)
}
