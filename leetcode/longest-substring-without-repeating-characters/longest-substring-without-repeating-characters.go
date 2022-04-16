package main

// source: https://leetcode.com/problems/longest-substring-without-repeating-characters/

// at first I've tried to get the correct answer by finding the biggest distance between the same chars
// and then finding the minimum of those. After several tries I've understood that it's to hard
// to deal with severeal cases with such approach
// Therefore I've implemented what I believe is called 'sliding window' (at least this name is
// very informative in case of the problem)
func lengthOfLongestSubstring(s string) int {
	var start int
	var res = 0
	positions := make(map[rune]int)

	for pos, char := range s {
		if prevPos, ppExists := positions[char]; ppExists {
			if pos-start > res {
				res = pos - start
			}
			if start < prevPos+1 {
				start = prevPos + 1
			}
		}
		positions[char] = pos
	}

	if len(s)-start > res {
		res = len(s) - start
	}

	return res
}

func main() {
	print(lengthOfLongestSubstring("abba"))
	print(lengthOfLongestSubstring("aab"))
	print(lengthOfLongestSubstring("abb"))
	print(lengthOfLongestSubstring("bwf"))
	print(lengthOfLongestSubstring("abcb"))
	print(lengthOfLongestSubstring("abcabcbb"))
	print(lengthOfLongestSubstring("pwwkew"))
}
