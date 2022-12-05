package main

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	n := len(words)
	for i := 0; i < n-1; i++ {
		if words[i][len(words[i])-1] != words[i+1][0] {
			return false
		}
	}
	return words[n-1][len(words[n-1])-1] == words[0][0]
}

func main() {
	isCircularSentence("leetcode exercises sound delightful")
}
