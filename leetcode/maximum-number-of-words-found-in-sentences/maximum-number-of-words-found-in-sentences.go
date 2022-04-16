package main

import "strings"

// source: https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/

func mostWordsFound1(sentences []string) int {
	var tmp int
	res := strings.Count(sentences[0], " ")

	for i := 1; i < len(sentences); i++ {
		if tmp = strings.Count(sentences[i], " "); tmp > res {
			res = tmp
		}
	}

	return res + 1
}

func mostWordsFound(sentences []string) int {
	var tmp int
	res := len(strings.Split(sentences[0], " "))

	for i := 1; i < len(sentences); i++ {
		if tmp = len(strings.Split(sentences[i], " ")); tmp > res {
			res = tmp
		}
	}

	return res
}
